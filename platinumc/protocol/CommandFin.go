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

//NewFinObject can create a new object
func NewFinObject() *CommandFin {
	br := &CommandFin{
		ErrorCode: 0,
	}
	br.Head.ProtocolVersion = ProtocolVersion
	br.Head.CommandID = CommandFin1
	br.Head.BodyLength = br.GetBodyLength()
	return br
}

//GetBodyLength get body length
func (f *CommandFin) GetBodyLength() uint16 {
	return (uint16)(1)
}

// EncodeBody can encode client message to binary
func (f *CommandFin) EncodeBody() ([]byte, error) {
	buff := new(bytes.Buffer)
	var err error
	err = binary.Write(buff, binary.BigEndian, f.Head.ProtocolVersion)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buff, binary.BigEndian, f.Head.CommandID)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buff, binary.BigEndian, f.Head.BodyLength)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buff, binary.BigEndian, f.ErrorCode)
	if err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}

// DecodeBody can decode binary code to struct
func (f *CommandFin) DecodeBody(buff *bytes.Buffer) (*CommandFin, error) {

	data := NewFinObject()
	//buff := bytes.NewBuffer(buf)

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
