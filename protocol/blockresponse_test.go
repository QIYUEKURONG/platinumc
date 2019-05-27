package protocol

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlockResponseGetBodyLength(t *testing.T) {
	assert := assert.New(t)
	br := NewBlockResponse()
	assert.NotNil(br)
	br.GetBodyLength()
}
func TestBlockResponseEncode(t *testing.T) {
	assert := assert.New(t)
	br := NewBlockResponse()
	assert.NotNil(br)
	br.Head.ProtocolVersion = 1
	br.Head.CommandID = CommandBlockResponse
	br.FileIndex = "/index"
	br.FileOffset = 0
	br.FileSize = 0
	br.FilelastModified = 0
	br.Head.BodyLength = br.GetBodyLength()
	buff, err := br.Encode()
	assert.Nil(err)
	assert.Equal(buff[0], (byte)(1), "ProtocolVersion")
	assert.Equal(buff[1], byte(0x31), "CommandBlockResponse")
	assert.Equal(buff[2], byte(0), "GetBodyLength")
	assert.Equal(buff[4], byte('/'), "FileIndex")
	assert.Equal(buff[5], byte('i'), "FileIndex")
	assert.Equal(buff[6], byte('n'), "FileIndex")
	assert.Equal(buff[7], byte('d'), "FileIndex")
	assert.Equal(buff[8], byte('e'), "FileIndex")
	assert.Equal(buff[9], byte('x'), "FileIndex")
	assert.Equal(buff[10], byte(0), "FileOffert")
	assert.Equal(buff[11], byte(0), "FileOffert")
	assert.Equal(buff[12], byte(0), "FileOffert")
	assert.Equal(buff[13], byte(0), "FileOffert")
	assert.Equal(buff[14], byte(0), "FileOffert")
	assert.Equal(buff[15], byte(0), "FileOffert")
	assert.Equal(buff[16], byte(0), "FileOffert")
	assert.Equal(buff[18], byte(0), "FileOffert")
}
func TestBlockResponseDecode(t *testing.T) {
	assert := assert.New(t)
	br := NewBlockResponse()
	assert.NotNil(br)
	br.Head.ProtocolVersion = 1
	br.Head.CommandID = CommandBlockResponse
	br.FileIndex = "/index"
	br.FileOffset = uint64(0)
	br.FileSize = uint64(0)
	br.FilelastModified = uint64(0)
	br.Head.BodyLength = br.GetBodyLength()
	message, err := br.Encode()
	assert.Nil(err)
	buff := bytes.NewBuffer(message)
	b, err := br.DecodeBody(buff)
	assert.Nil(err)
	assert.Equal(b.Head.ProtocolVersion, byte(1), "ProtocolVersion")
	assert.Equal(b.Head.CommandID, byte(0x31), "CommandID")
	assert.Equal(b.FileIndex, "/index", "FileIndex")
	assert.Equal(b.FileOffset, uint64(0), "FileOffset")
	assert.Equal(b.FileSize, uint64(0), "FileSize")
	assert.Equal(b.FilelastModified, uint64(0), "FilelastModifyied")
}
