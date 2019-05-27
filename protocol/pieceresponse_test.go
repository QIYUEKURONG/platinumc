package protocol

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPieceResponseGetBodyLength(t *testing.T) {
	assert := assert.New(t)
	br := NewPieceResponse()
	assert.NotNil(br)
	length := br.GetBodyLength()
	assert.Equal(length, uint16(12), "GetBodyLength")
}
func TestPieceResponseEncode(t *testing.T) {
	assert := assert.New(t)
	br := NewPieceResponse()
	assert.NotNil(br)
	br.Head.ProtocolVersion = ProtocolVersion
	br.Head.CommandID = CommandPieceResponse
	br.PieceIndex = 0
	br.PieceHash = 0
	br.PieceSize = 0
	br.PieceData = "ju"
	buff, err := br.EncodeBody()
	assert.Nil(err)
	assert.Equal(buff[0], byte(1), "ProtocolVersion")
	assert.Equal(buff[1], byte(0x33), "CommandID")
	assert.Equal(buff[4], byte(0), "PieceIndex")
	assert.Equal(buff[5], byte(0), "PieceIndex")
	assert.Equal(buff[6], byte(0), "PieceIndex")
	assert.Equal(buff[7], byte(0), "PieceIndex")
	assert.Equal(buff[8], byte(0), "PieceHash")
	assert.Equal(buff[9], byte(0), "PieceHash")
	assert.Equal(buff[10], byte(0), "PieceHash")
	assert.Equal(buff[11], byte(0), "PieceHash")
	assert.Equal(buff[12], byte(0), "PieceSize")
	assert.Equal(buff[13], byte(0), "PieceSize")
	assert.Equal(buff[14], byte(0), "PieceSize")
	assert.Equal(buff[15], byte(0), "PieceSize")
	assert.Equal(buff[20], byte('j'), "PieceData")
	assert.Equal(buff[21], byte('u'), "PieceData")
}
func TestPieceResponseDecode(t *testing.T) {
	assert := assert.New(t)
	br := NewPieceResponse()
	assert.NotNil(br)
	br.Head.ProtocolVersion = ProtocolVersion
	br.Head.CommandID = CommandPieceResponse
	br.PieceIndex = 0
	br.PieceHash = 0
	br.PieceSize = 0
	br.PieceData = "ju"
	br.Head.BodyLength = br.GetBodyLength()
	message, err := br.EncodeBody()
	buff := bytes.NewBuffer(message)
	b, err := br.DecodeBody(buff)
	assert.Nil(err)
	assert.Equal(b.Head.ProtocolVersion, byte(1), "ProtocolVersion")
	assert.Equal(b.Head.CommandID, byte(0x33), "CommandID")
	assert.Equal(b.PieceHash, uint32(0), "PieceHash")
	assert.Equal(b.PieceIndex, uint32(0), "PieceIndex")
	assert.Equal(b.PieceSize, uint32(0), "PieceSize")
	assert.Equal(b.PieceData, "ju", "PieceData")
}
