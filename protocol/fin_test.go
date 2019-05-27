package protocol

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBodyLengthFin(t *testing.T) {
	assert := assert.New(t)
	br := NewFinObject()
	assert.NotNil(br)
	length := br.GetBodyLength()
	assert.Equal(length, uint16(1), "GetBodyLength")
}
func TestCommandFinEncode(t *testing.T) {
	assert := assert.New(t)
	br := NewFinObject()
	assert.NotNil(br)
	buff, err := br.Encode()
	assert.Nil(err)
	assert.Equal(buff[0], byte(1), "ProtocolVersion")
	assert.Equal(buff[1], byte(0x07), "CommandFin1")
	assert.Equal(buff[4], byte(0), "ErrorCode")
}
func TestCommandFinDecode(t *testing.T) {
	assert := assert.New(t)
	br := NewFinObject()
	assert.NotNil(br)
	message, err := br.Encode()
	assert.Nil(err)
	buff := bytes.NewBuffer(message)
	b, err := br.DecodeBody(buff)
	assert.Nil(err)
	assert.Equal(b.Head.ProtocolVersion, byte(1), "ProtocolVersion")
	assert.Equal(b.Head.CommandID, byte(0x7), "CommandID")
	assert.Equal(b.ErrorCode, byte(0), "ErrorCode")
}
