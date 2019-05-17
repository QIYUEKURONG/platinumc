package client

import (
	"bytes"
	"fmt"
	"os"

	"github.com/QIYUEKURONG/platinumc/platinumc"
	"github.com/QIYUEKURONG/platinumc/platinumc/protocol"
)

// Run the client to download file specefied in task.
func Run(task *platinumc.Task) {
	// 1. 创建应用层连接，负责处理消息编解
	c := NewConnectionTCP(task)
	conn, err := c.Connect()
	if err != nil {
		fmt.Println("call Connect failed: %v\n", err)
		os.Exit(-1)
	}
	err = SendBlockRequest(task, conn)
	if err != nil {
		fmt.Printf("call SendBlockRequest failed: %v\n", err)
		os.Exit(1)
	}
	//  建立一个文件，存储最后的文件
	file, err := os.Create("map.mp4")
	if err != nil {
		fmt.Println("create a file error")
		os.Exit(-2)
	}
	// 为了最后打印做准备
	fileName := task.SavePath
	var fileLength uint64
	//var startTime int
	//var endTime int
	//创建一个缓冲区，用于存放服务器发送来的消息
	recvBuff := make([]byte, 1024)
	dataBuff := bytes.NewBuffer([]byte{})
	var pieceIndex uint32
	for {
		lengthSize, err := conn.Read(recvBuff)
		//fmt.Println("recvBuff and lengthSize", recvBuff, lengthSize)
		t := recvBuff[0:lengthSize]
		dataBuff.Write(t)
		//	fmt.Println("databuff ", dataBuff.Bytes(), dataBuff.Len())
		if lengthSize == 0 {
			break
		}
		if err != nil {
			fmt.Println("Sorry! read error")
			os.Exit(-1)
		}
		if (dataBuff.Len()) < 4 {
			continue
		}
		//解析长度
		bodyLength, err := GetBodyLength(dataBuff.Bytes())
		//fmt.Println("bodyLength valus", bodyLength)
		if err != nil {
			fmt.Println("sorry ! to read body length error!")
			os.Exit(-1)
		}
		if bodyLength > (uint16)((dataBuff.Len())-4) {
			continue
		}
		//解析类型
		var commandID uint8
		commandID, err = GetCommandID(dataBuff.Bytes())
		if err != nil {
			fmt.Println("Sorry ! binary convert error")
			os.Exit(-1)
		}
		// 选择类型
		switch commandID {
		case protocol.CommandFin1:
			fin := protocol.NewFinObject()
			fin, err := fin.DecodeBody(dataBuff)
			if err != nil {
				fmt.Println("sorry ! Fin Decode error")
				continue
			} else {
				fmt.Println("erron is:", fin.ErrorCode)
				os.Exit(-1)
			}
		case protocol.CommandBlockResponse:
			Response := protocol.NewBlockResponse()
			Response, err = Response.DecodeBody(dataBuff)
			if err != nil {
				fmt.Println("sorry ! Response Decode error")
				os.Exit(-1)
			}
			str := SetLocalTime()
			fmt.Printf("[%s][S2C BlockResponse] Fileindex:%s Offset:%v FileSize:%v LastModified:%v\n",
				str, Response.FileIndex, Response.FileSize, Response.FileOffset, Response.FilelastModified)
			fileLength = Response.FileSize
			err = SendPieceRequest(task, pieceIndex, conn)
			if err != nil {
				fmt.Printf("call SendPieceRequest failed: %v\n", err)
			}
		case protocol.CommandPieceResponse:
			pieceResponse := protocol.NewPieceResponse()
			pieceResponse, err = pieceResponse.DecodeBody(dataBuff)
			if err != nil {
				fmt.Println("sorry ! the pieceresponse decode error")
				os.Exit(-1)
			}
			file.Write(([]byte)(pieceResponse.PieceData))
			str := SetLocalTime()
			fmt.Printf("[%s][S2C pieceResponse] Index:%v  Hash:%v Length:%v\n", str, pieceIndex, pieceResponse.PieceHash, bodyLength-12)
			fmt.Printf("[%s]PieceIndex:%d finish.\n", str, pieceIndex)
			pieceIndex++
			//fmt.Println(pieceResponse)
			//说明文件已经接收完毕
			if (bodyLength - 12) < 8192 {
				fmt.Printf("[%s] Download finish.\n", str)
				fmt.Printf("[%s] Save file block to %v  FileSize %v\n", str, fileName, fileLength)
				err = SendFin(conn)
				if err != nil {
					fmt.Printf("call SendFin failed: %v\n", err)
				}
				break
			}
			err = SendPieceRequest(task, pieceIndex, conn)
			//fmt.Printf("PieceData %s \n", pieceResponse.PieceData)
		}
	}
}
