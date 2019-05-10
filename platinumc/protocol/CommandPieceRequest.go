package protocol

import (
	"bytes"
	"encoding/binary"

	"github.com/QIYUEKURONG/platinumc/platinumc"
)

// PieceRequest can tell server the piecenindex
type PieceRequest struct {
	Head        Header
	PiecenIndex uint32
}

//NewObject can create a new object
func (p *PieceRequest) NewObject(task *platinumc.Task) *PieceRequest {
	message := new(PieceRequest)
	message.Head.ProtocolVersion = ProtocolVersion
	message.Head.CommandID = 002
	message.Head.BodyLength = p.GetBodyLength()
	message.PiecenIndex = (uint32)(task.StartPieceIndex)
	return message
}

// GetBodyLength can return value of bodylength
func (p *PieceRequest) GetBodyLength() uint16 {
	return (uint16)(4)
}

// EncodeBody can encode client message to binary
func (p *PieceRequest) EncodeBody(message *PieceRequest) ([]byte, error) {
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
	err = binary.Write(buff, binary.BigEndian, message.PiecenIndex)
	if err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}

// DecodeBody can decode binary code to struct
func (p *PieceRequest) DecodeBody(buf []byte) (PieceRequest, error) {
	var piece PieceRequest
	buff := bytes.NewBuffer(buf)
	var err error
	err = binary.Read(buff, binary.BigEndian, &piece.Head.ProtocolVersion)
	if err != nil {
		return piece, err
	}
	err = binary.Read(buff, binary.BigEndian, &piece.Head.CommandID)
	if err != nil {
		return piece, err
	}
	err = binary.Read(buff, binary.BigEndian, &piece.Head.BodyLength)
	if err != nil {
		return piece, err
	}
	err = binary.Read(buff, binary.BigEndian, &piece.PiecenIndex)
	if err != nil {
		return piece, nil
	}
	return piece, nil
}
