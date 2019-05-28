package protocol

import (
	"bytes"
	"encoding/binary"
)

// PingReply record ping
type PingReply struct {
	Head      Header
	TimeStamp uint64
}

//NewPingReply can create a new object
func NewPingReply(timestamp uint64) *PingReply {
	br := &PingReply{}
	br.Head.ProtocolVersion = ProtocolVersion
	br.Head.CommandID = 0x06
	br.Head.BodyLength = br.GetBodyLength()
	br.TimeStamp = timestamp
	return br
}

// GetBodyLength get body length
func (p *PingReply) GetBodyLength() uint16 {
	return (uint16)(8)
}

// EncodeBody can encode client message to binary
func (p *PingReply) EncodeBody() ([]byte, error) {
	buff := new(bytes.Buffer)
	var err error
	err = binary.Write(buff, binary.BigEndian, p.Head.ProtocolVersion)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buff, binary.BigEndian, p.Head.CommandID)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buff, binary.BigEndian, p.Head.BodyLength)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buff, binary.BigEndian, p.TimeStamp)
	if err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}

// DecodeBody can decode binary code to struct
func (p *PingReply) DecodeBody(buf []byte) (*PingReply, error) {
	buff := bytes.NewBuffer(buf)
	ping := NewPingReply(0)
	var err error
	err = binary.Read(buff, binary.BigEndian, &ping.Head.ProtocolVersion)
	if err != nil {
		return ping, err
	}
	err = binary.Read(buff, binary.BigEndian, &ping.Head.CommandID)
	if err != nil {
		return ping, err
	}
	err = binary.Read(buff, binary.BigEndian, &ping.Head.BodyLength)
	if err != nil {
		return ping, err
	}
	err = binary.Read(buff, binary.BigEndian, &ping.TimeStamp)
	if err != nil {
		return ping, err
	}
	return ping, nil
}
