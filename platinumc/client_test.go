package platinumc

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryEncode(t *testing.T) {
	assert := assert.New(t)

	msg := BlockRequest{}
	msg.Head.ProtocolVersion = 1
	msg.Head.CommandID = 1
	msg.Head.BodyLength = 30
	msg.ClientType = 2
	msg.ClientID = "id"
	msg.FileIndex = "1.mp4"
	msg.FileOffset = 234283

	buf, err := SerializateBinary(&msg)
	assert.Nil(err)

	//fmt.Println(buf)
	//fmt.Println(buf[9:11])

	length := 0
	// Head length
	length += 1 // uint8 1 byte ProtocolVersion
	length += 1 // uint8 1 byte CommandID
	length += 2 // uint16 2 bytes BodyLength
	// Body length
	length += 1                     // uint8 1 byte ClientType
	length += 4 + len(msg.ClientID) // uint32 + len(str)   4+len(str) bytes  ClientID
	length += 4 + len(msg.FileIndex)
	length += 8

	//fmt.Println(uint16(30))
	if length != len(buf) {
		t.Errorf("length should be: %v\n", length)
	}
	//fmt.Println(buf[])
	assert.Equal(buf[0], byte(1), "first byte should be 1")
	assert.Equal(buf[1], byte(1), "second byte should be 1")
	assert.Equal(buf[3], uint8(30), "third byte should be 30")
	assert.Equal(buf[4], uint8(2), "fourth string should be id")
	assert.Equal(buf[9], (byte)('i'), "i must right")
	assert.Equal(buf[10], (byte)('d'), "d must right")
	assert.Equal(buf[15], (byte)('1'), "1 must right")
	assert.Equal(buf[16], (byte)('.'), ". must right")
	assert.Equal(buf[17], (byte)('m'), "m must right")
	assert.Equal(buf[18], (byte)('p'), "p must right")
	assert.Equal(buf[19], (byte)('4'), "4 must right")
	//assert.Equal(buf[20], byte(234283), "234283 must right")
	/*
		expected := []byte{1, 1, 0, 30}
		if buf[0] != 1 {
		}*/
}

func TestBinaryDecode(t *testing.T) {
	msg := BlockRequest{}
	msg.Head.ProtocolVersion = 1
	msg.Head.CommandID = 1
	msg.Head.BodyLength = 30
	msg.ClientType = 2
	msg.ClientID = "id"
	msg.FileIndex = "1.mp4"
	msg.FileOffset = 234283

	buf, _ := SerializateBinary(&msg)
	//	assert.Nil(err)
	fmt.Println("buf value", buf)
	var str uint8
	buff := bytes.NewBuffer(buf)
	binary.Read(buff, binary.BigEndian, &str)
	fmt.Println(str)
	fmt.Println("buf value", buf)

}
