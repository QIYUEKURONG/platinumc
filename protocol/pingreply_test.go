package protocol

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingReplyGetbodyLength(t *testing.T) {
	assert := assert.New(t)
	br := NewPingReply(0)
	assert.NotNil(br)
	length := br.GetBodyLength()
	assert.Equal(length, uint16(8), "GetBodyLength")
}

func TestPingReplyEncode(t *testing.T) {
	assert := assert.New(t)
	br := NewPingReply(0)
	assert.NotNil(br)
	buff, err := br.EncodeBody()
	assert.Nil(err)
	assert.Equal(buff[0], byte(1), "ProtocolVersions")
	assert.Equal(buff[1], byte(0x06), "commandID")
	assert.Equal(buff[2], byte(0), "GetBodyLength")
}
func TestPingReplyDecode(t *testing.T) {
	assert := assert.New(t)
	br := NewPingReply(0)
	assert.NotNil(br)
	buff, err := br.EncodeBody()
	assert.Nil(err)
	b, err := br.DecodeBody(buff)
	assert.Nil(err)
	assert.Equal(b.Head.ProtocolVersion, byte(1), "ProtocolVersion")
	assert.Equal(b.Head.CommandID, byte(0x06), "CommandID")
	assert.Equal(b.Head.BodyLength, uint16(8), "GetBodyLength")
	assert.Equal(b.TimeStamp, uint64(0), "TimeStamp")
}
