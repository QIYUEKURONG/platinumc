package protocol

import (
	"bytes"
	"encoding/binary"
)

// CommandFin struct descibe the ping error
type CommandFin struct {
	Head      Header
	ErrorCode uint8
}

//NewObject can create a new object
func (f *CommandFin) NewObject() *CommandFin {
	message := new(CommandFin)
	message.Head.ProtocolVersion = ProtocolVersion
	message.Head.CommandID = 002
	message.Head.BodyLength = f.GetBodyLength()
	message.ErrorCode = 0
	return message
}

//GetBodyLength get body length
func (f *CommandFin) GetBodyLength() uint16 {
	return (uint16)(1)
}

// EncodeBody can encode client message to binary
func (f *CommandFin) EncodeBody(message *CommandFin) ([]byte, error) {
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
	err = binary.Write(buff, binary.BigEndian, message.ErrorCode)
	if err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}

// DecodeBody can decode binary code to struct
func (f *CommandFin) DecodeBody(buf []byte) (CommandFin, error) {
	buff := bytes.NewBuffer(buf)
	var data CommandFin
	var err error
	err = binary.Read(buff, binary.BigEndian, &data.Head.ProtocolVersion)
	if err != nil {
		return data, err
	}
	err = binary.Read(buff, binary.BigEndian, &data.Head.CommandID)
	if err != nil {
		return data, err
	}
	err = binary.Read(buff, binary.BigEndian, &data.Head.BodyLength)
	if err != nil {
		return data, err
	}
	err = binary.Read(buff, binary.BigEndian, &data.ErrorCode)
	if err != nil {
		return data, err
	}
	return data, nil
}
