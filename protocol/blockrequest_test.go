package protocol

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlockRequestGetBodyLength(t *testing.T) {
	assert := assert.New(t)
	br := NewBlockRequest("001", 0x30, "/index/", 0)
	assert.NotNil(br)
	br.GetBodyLength()
}
func TestBlockRequestEncode(t *testing.T) {
	assert := assert.New(t)
	br := NewBlockRequest("001", CommandBlockRequest, "/index", 0)
	assert.NotNil(br)
	buff, err := br.Encode()
	assert.Nil(err)
	fmt.Println(buff)
	assert.Equal(buff[0], byte(1), "ProtocolVersion")
	assert.Equal(buff[1], byte(0x30), "CommandBlockRequest")
	assert.Equal(buff[3], byte(26), "BodyLength")
	assert.Equal(buff[4], byte(0x30), "clientType")
	assert.Equal(buff[9], uint8('0'), "clientID")
	assert.Equal(buff[10], uint8('0'), "clientID")
	assert.Equal(buff[11], uint8('1'), "clientID")
	assert.Equal(buff[16], byte('/'), "FileIndex")
	assert.Equal(buff[17], byte('i'), "FileIndex")
	assert.Equal(buff[18], byte('n'), "FileIndex")
	assert.Equal(buff[19], byte('d'), "FileIndex")
	assert.Equal(buff[20], byte('e'), "FileIndex")
	assert.Equal(buff[21], byte('x'), "FileIndex")
	assert.Equal(buff[22], byte(0), "FileOffert")
	assert.Equal(buff[23], byte(0), "FileOffert")
	assert.Equal(buff[24], byte(0), "FileOffert")
	assert.Equal(buff[25], byte(0), "FileOffert")
	assert.Equal(buff[26], byte(0), "FileOffert")
	assert.Equal(buff[27], byte(0), "FileOffert")
	assert.Equal(buff[28], byte(0), "FileOffert")
	assert.Equal(buff[29], byte(0), "FileOffert")
}
func TestBlockRequestDecode(t *testing.T) {
	assert := assert.New(t)
	br := NewBlockRequest("001", CommandBlockRequest, "/index", 0)
	assert.NotNil(br)
	buff, err := br.Encode()
	assert.Nil(err)
	b, err := br.DecodeBody(buff)
	assert.Nil(err)
	assert.Equal(b.Head.ProtocolVersion, byte(1), "ProtocolVersion")
	assert.Equal(b.Head.CommandID, byte(CommandBlockRequest), "CommandID")
	assert.Equal(b.FileIndex, "/index", "FileIndex")
	assert.Equal(b.FileOffset, uint64(0), "FileOffset")
	assert.Equal(b.ClientID, "001", "ClientID")
}
