package protocol

import (
	"bytes"
	"encoding/binary"
)

// PieceRequest can tell server the piecenindex
type PieceRequest struct {
	Head        Header
	PiecenIndex uint32
}

//NewPieceRequest can create a new object
func NewPieceRequest(index uint32) *PieceRequest {
	br := &PieceRequest{}
	br.PiecenIndex = index
	br.Head.ProtocolVersion = ProtocolVersion
	br.Head.CommandID = 0x32
	br.Head.BodyLength = br.GetBodyLength()
	return br
}

// GetBodyLength can return value of bodylength
func (p *PieceRequest) GetBodyLength() uint16 {
	return (uint16)(4)
}

// Encode can encode client message to binary
func (p *PieceRequest) Encode() ([]byte, error) {
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
	err = binary.Write(buff, binary.BigEndian, p.PiecenIndex)
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
