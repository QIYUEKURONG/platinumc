package client

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/QIYUEKURONG/platinumc/protocol"
)

// StartTCP to start a tcp link
func StartTCP(task *Task) {
	//log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	// 1. 创建应用层连接，负责处理消息编解
	conn, err := NewConnectionTCP(task)
	if err != nil {
		fmt.Printf("call Connect failed: %v\n", err)
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
	//创建一个缓冲区，用于存放服务器发送来的消息
	recvBuff := make([]byte, 1024)
	dataBuff := bytes.NewBuffer([]byte{})
	var pieceIndex uint32
	for {
		lengthSize, err := conn.Read(recvBuff)
		t := recvBuff[0:lengthSize]
		dataBuff.Write(t)
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
		if err != nil {
			fmt.Println("sorry ! to read body length error!")
			os.Exit(-1)
		}
		//如何获得的除去头部的长度比总体长度小的话。继续循环接受
		if bodyLength > (uint16)((dataBuff.Len())-4) {
			continue
		}
		//解析类型
		commandID, err := GetCommandID(dataBuff.Bytes())
		if err != nil {
			fmt.Printf("call GetCommandID error : %v", err)
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
			Response.PrintfBlockResponse()
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
			pieceResponse.PrintfPieceResponse(pieceIndex, bodyLength)
			pieceIndex++
			//说明文件已经接收完毕
			if (bodyLength - 12) < 8192 {
				log.Printf("Save file block to %v  FileSize %v", fileName, fileLength)
				err = SendFin(conn)
				if err != nil {
					fmt.Printf("call SendFin failed: %v\n", err)
				}
				break
			}
			err = SendPieceRequest(task, pieceIndex, conn)
		}
	}
}

// StartWSS to create a wss link
func StartWSS(task *Task) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	// 1. 创建应用层连接，负责处理消息编解
	conn, err := NewConnectionWSS(task)
	if err != nil {
		fmt.Printf("call NewConnectionWSS error :%v", err)
	}
	err = SendBlockRequestWSS(task, conn)
	if err != nil {
		fmt.Printf("call SendBlockRequest failed: %v\n", err)
		os.Exit(1)
	}
	//  建立一个文件，存储最后的文件
	f, err := os.Create("my.mp4")
	if err != nil {
		fmt.Println("create a file error")
		os.Exit(-2)
	}
	// 为了最后打印做准备
	fileName := task.SavePath
	//var fileLength uint64
	//创建一个缓冲区，用于存放服务器发送来的消息
	//recvBuff := make([]byte, 1024)
	var pieceIndex uint32
	for {
		_, recvBuff, err := conn.ReadMessage()

		if err != nil {
			log.Printf("Save file block to %v  ", fileName)
			os.Exit(-1)
		}
		//解析长度
		bodyLength, err := GetBodyLength(recvBuff)
		if err != nil {
			fmt.Println("sorry ! to read body length error!")
			os.Exit(-1)
		}
		//解析类型
		commandID, err := GetCommandID(recvBuff)
		if err != nil {
			fmt.Printf("call GetCommandID error : %v", err)
			os.Exit(-1)
		}
		switch commandID {
		case protocol.CommandFin1:
			dataBuff := bytes.NewBuffer(recvBuff)
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
			dataBuff := bytes.NewBuffer(recvBuff)
			Response := protocol.NewBlockResponse()
			Response, err = Response.DecodeBody(dataBuff)
			if err != nil {
				fmt.Println("sorry ! Response Decode error")
				os.Exit(-1)
			}
			Response.PrintfBlockResponse()
			//fileLength = Response.FileSize
			err = SendPieceRequestWSS(task, pieceIndex, conn)
			if err != nil {
				fmt.Printf("call SendPieceRequest failed: %v\n", err)
			}
		case protocol.CommandPieceResponse:
			dataBuff := bytes.NewBuffer(recvBuff)
			pieceResponse := protocol.NewPieceResponse()
			pieceResponse, err = pieceResponse.DecodeBody(dataBuff)
			if err != nil {
				fmt.Println("sorry ! the pieceresponse decode error")
				os.Exit(-1)
			}
			f.Write(([]byte)(pieceResponse.PieceData))
			pieceResponse.PrintfPieceResponse(pieceIndex, bodyLength)
			pieceIndex++
			//说明文件已经接收完毕
			err = SendPieceRequestWSS(task, pieceIndex, conn)
		}
	}
}
