package client

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/QIYUEKURONG/platinumc/platinumc"
	"github.com/QIYUEKURONG/platinumc/platinumc/protocol"
)

// SetLocalTime can return a string of local time
func SetLocalTime() string {

	year, month, day := time.Now().Date()
	str := fmt.Sprintf("%d-%d-%d", year, month, day)
	hour, minute, second := time.Now().Clock()
	str += fmt.Sprintf("  %d-%d-%d", hour, minute, second)
	nanoSecond := time.Now().Nanosecond()
	str += fmt.Sprintf("-%d", nanoSecond)
	return str
}

// GetBodyLength function get body length
func GetBodyLength(buff []byte) (uint16, error) {
	//解析长度
	bodylengthbuff := buff[2:4]
	body := bytes.NewBuffer(bodylengthbuff)
	var bodylength uint16
	var err error
	err = binary.Read(body, binary.BigEndian, &bodylength)
	return bodylength, err
}

// GetCommandID function get  object type
func GetCommandID(buff []byte) (uint8, error) {
	commandid := buff[1:2]
	id := bytes.NewBuffer(commandid)
	var commandtype uint8
	err := binary.Read(id, binary.BigEndian, &commandtype)
	return commandtype, err
}

// SendBlockRequest encode and send BlockRequest to server.
func SendBlockRequest(task *platinumc.Task, conn net.Conn) error {
	// 1. 构造
	br := protocol.NewBlockRequest(task.Identifier, protocol.ClientTypeUnknown, task.FileIndex, 0)
	// 2. 编码
	sendBuffer, err := br.Encode()
	if err != nil {
		return fmt.Errorf("encode failed: %v", err)
	}
	// 3. 发送
	_, err = conn.Write(sendBuffer)
	if err != nil {
		return fmt.Errorf("send failed: %v", err)
	}
	str := SetLocalTime()
	fmt.Printf("[%s] [C2S BlockRequest] FileIndex:%s ClientId:tcp-%s FileOffset:%v\n", str, br.FileIndex, br.ClientID, br.FileOffset)
	return nil
}

// SendPieceRequest encode and send PieceRequest to server.
func SendPieceRequest(t *platinumc.Task, index uint32, conn net.Conn) error {
	// 1: 构造
	br := protocol.NewPieceRequest(t, index)
	// 2：编码
	sendBuffer, err := br.Encode()
	if err != nil {
		return fmt.Errorf("encode failed: %v", err)
	}
	// 3:发送
	_, err = conn.Write(sendBuffer)
	if err != nil {
		return fmt.Errorf("send failed: %v", err)
	}
	str := SetLocalTime()
	fmt.Printf("[%s] [C2S PieceRequest] Index:%v Offset:%v Length:%v\n", str, index, br.PiecenIndex, 8192)

	return nil
}

// SendFin encode and send Fin to server
func SendFin(conn net.Conn) error {
	// 1:构造
	br := protocol.NewFinObject()
	//2：编码
	sendBuffer, err := br.Encode()
	if err != nil {
		return fmt.Errorf("encode failed: %v", err)
	}
	//3:发送
	_, err = conn.Write(sendBuffer)
	if err != nil {
		return fmt.Errorf("send failed: %v", err)
	}
	str := SetLocalTime()
	fmt.Printf("[%s] [C2S Fin] Code 53", str)
	return nil
}

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

	// 去了最后打印做准备
	fileName := task.SavePath
	var fileLength uint64
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
		//	fmt.Println("commandtype ", commandID)
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
			fmt.Printf("[%s] [S2C BlockResponse] Fileindex:%s Offset:%v FileSize:%v LastModified:%v\n",
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
			str := SetLocalTime()
			fmt.Printf("[%s] [S2C pieceResponse] Index:%v  Hash:%v Length:%v\n", str, pieceIndex, pieceResponse.PieceHash, bodyLength-12)
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
