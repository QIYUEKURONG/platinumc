package client

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"

	"github.com/QIYUEKURONG/platinumc/protocol"
	"github.com/gorilla/websocket"
)

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
	commandtype := buff[1:2]
	id := bytes.NewBuffer(commandtype)
	var commandid uint8
	err := binary.Read(id, binary.BigEndian, &commandid)
	if err != nil {
		fmt.Errorf("binary read error :%v", err)
	}
	return commandid, nil
}

// SendBlockRequestWSS can send block of wss
func SendBlockRequestWSS(task *Task, conn *websocket.Conn) error {
	// 1. 构造
	br := protocol.NewBlockRequest(task.Identifier, protocol.ClientTypeUnknown, task.FileIndex, 0)
	// 2. 编码
	sendBuffer, err := br.Encode()
	if err != nil {
		return fmt.Errorf("encode failed: %v", err)
	}
	// 3. 发送
	err = conn.WriteMessage(websocket.TextMessage, sendBuffer)
	if err != nil {
		return fmt.Errorf("send failed: %v", err)
	}
	log.Printf("[C2S BlockRequest] FileIndex:%s ClientId:tcp-%s FileOffset:%v\n", br.FileIndex, br.ClientID, br.FileOffset)
	return nil
}

// SendPieceRequestWSS encode and send PieceRequest to server.
func SendPieceRequestWSS(t *Task, index uint32, conn *websocket.Conn) error {
	// 1: 构造
	br := protocol.NewPieceRequest(index)
	// 2：编码
	sendBuffer, err := br.Encode()
	if err != nil {
		return fmt.Errorf("encode failed: %v", err)
	}
	// 3:发送
	err = conn.WriteMessage(websocket.TextMessage, sendBuffer)
	if err != nil {
		return fmt.Errorf("send failed: %v", err)
	}
	log.Printf("[C2S PieceRequest] Index:%v Offset:%v Length:%v\n", index, br.PiecenIndex, 8192)
	return nil
}

// SendFinWSS encode and send Fin to server
func SendFinWSS(conn *websocket.Conn) error {
	// 1:构造
	br := protocol.NewFinObject()
	//2：编码
	sendBuffer, err := br.Encode()
	if err != nil {
		return fmt.Errorf("encode failed: %v", err)
	}
	//3:发送
	err = conn.WriteMessage(websocket.TextMessage, sendBuffer)
	if err != nil {
		return fmt.Errorf("send failed: %v", err)
	}

	log.Printf("[C2S Fin] Code 53")
	return nil
}

// SendBlockRequest encode and send BlockRequest to server.
func SendBlockRequest(task *Task, conn net.Conn) error {
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

	log.Printf("[C2S BlockRequest] FileIndex:%s ClientId:tcp-%s FileOffset:%v\n", br.FileIndex, br.ClientID, br.FileOffset)
	return nil
}

// SendPieceRequest encode and send PieceRequest to server.
func SendPieceRequest(t *Task, index uint32, conn net.Conn) error {
	// 1: 构造
	br := protocol.NewPieceRequest(index)
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
	log.Printf("[C2S PieceRequest] Index:%v Offset:%v Length:%v\n", index, br.PiecenIndex, 8192)

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

	log.Printf("[C2S Fin] Code 53")
	return nil
}
