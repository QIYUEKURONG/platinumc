package protocol

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingGetBodyLength(t *testing.T) {
	assert := assert.New(t)
	br := NewPingObject(0)
	assert.NotNil(br)
	assert.Equal(br.GetBodyLength(), uint16(8), "GetBodyLength")
}
func TestPingEncode(t *testing.T) {
	assert := assert.New(t)
	br := NewPingObject(0)
	assert.NotNil(br)
	buff, err := br.EncodeBody()
	assert.Nil(err)
	assert.Equal(buff[0], byte(1), "ProtocolVersion")
	assert.Equal(buff[1], byte(0x05), "CommandID")
	assert.Equal(buff[3], uint8(8), "GetBodyLength")
	assert.Equal(buff[11], uint8(0), "TimeStamp")
}
func TestPingDecode(t *testing.T) {
	assert := assert.New(t)
	br := NewPingObject(0)
	assert.NotNil(br)
	buff, err := br.EncodeBody()
	assert.Nil(err)
	b, err := br.DecodeBody(buff)
	assert.Nil(err)
	assert.Equal(b.Head.ProtocolVersion, byte(1), "ProtocolVersion")
	assert.Equal(b.Head.CommandID, byte(0x05), "CommandID")
	assert.Equal(b.Head.BodyLength, uint16(8), "BodyLength")
	assert.Equal(b.TimeStamp, uint64(0), "TimeStamp")
}
