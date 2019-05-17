package client

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/QIYUEKURONG/platinumc/protocol"
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
	str := SetLocalTime()
	fmt.Printf("[%s][C2S BlockRequest] FileIndex:%s ClientId:tcp-%s FileOffset:%v\n", str, br.FileIndex, br.ClientID, br.FileOffset)
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
	str := SetLocalTime()
	fmt.Printf("[%s][C2S PieceRequest] Index:%v Offset:%v Length:%v\n", str, index, br.PiecenIndex, 8192)

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
	fmt.Printf("[%s][C2S Fin] Code 53", str)
	return nil
}

// ReturnLocalTime can return local time
func ReturnLocalTime() (int, error) {
	hour, minute, second := time.Now().Clock()
	nanoSecond := time.Now().Nanosecond()
	str := strconv.Itoa(hour) + strconv.Itoa(minute) + strconv.Itoa(second) + strconv.Itoa(nanoSecond)
	t, err := strconv.Atoi(str)
	if err != nil {
		fmt.Errorf("get time error:%v", err)
	}
	return t, nil
}
