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

//NewObject can create a new object
func (p *PingReply) NewObject(timestamp uint64) *PingReply {
	message := new(PingReply)
	message.Head.ProtocolVersion = ProtocolVersion
	message.Head.CommandID = 0x06
	message.Head.BodyLength = p.GetBodyLength()
	message.TimeStamp = timestamp
	return message
}

// GetBodyLength get body length
func (p *PingReply) GetBodyLength() uint16 {
	return (uint16)(8)
}

// EncodeBody can encode client message to binary
func (p *PingReply) EncodeBody(message *PingReply) ([]byte, error) {
	buff := new(bytes.Buffer)
	var err error
	err = binary.Write(buff, binary.BigEndian, message.Head.ProtocolVersion)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buff, binary.BigEndian, message.Head.CommandID)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buff, binary.BigEndian, message.Head.BodyLength)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buff, binary.BigEndian, message.TimeStamp)
	if err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}

// DecodeBody can decode binary code to struct
func (p *PingReply) DecodeBody(buf []byte) (PingReply, error) {
	buff := bytes.NewBuffer(buf)
	var ping PingReply
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
