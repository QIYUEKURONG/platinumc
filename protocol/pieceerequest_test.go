package protocol

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBodyLengthPiece(t *testing.T) {
	assert := assert.New(t)
	br := NewPieceRequest(0)
	assert.NotNil(br)
	length := br.GetBodyLength()
	assert.Equal(length, uint16(4), "GetBodyLength")
}
func TestPieceeRequestEncode(t *testing.T) {
	assert := assert.New(t)
	br := NewPieceRequest(0)
	assert.NotNil(br)
	buff, err := br.Encode()
	assert.Nil(err)
	assert.Equal(buff[0], byte(1), "ProtocolVersion")
	assert.Equal(buff[1], byte(0x32), "CommandID")
	assert.Equal(buff[4], byte(0), "Index")
}
func TestPieceeRequestDecode(t *testing.T) {
	assert := assert.New(t)
	br := NewPieceRequest(0)
	assert.NotNil(br)
	buff, err := br.Encode()
	assert.Nil(err)
	b, err := br.DecodeBody(buff)
	assert.Nil(err)
	assert.Equal(b.Head.ProtocolVersion, byte(1), "ProtocolVersion")
	assert.Equal(b.Head.CommandID, byte(0x32), "CommandID")
	assert.Equal(b.PiecenIndex, uint32(0), "Index")

}
