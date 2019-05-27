package protocol

import (
	"bytes"
	"encoding/binary"
	"log"
)

// PieceResponse struct record all message of send filee
type PieceResponse struct {
	Head       Header
	PiceeIndex uint32
	PieceHash  uint32
	PieceSize  uint32
	PieceData  string
}

// GetBodyLength can return value of bodylength
func (p *PieceResponse) GetBodyLength() uint16 {
	return (uint16)(4 + 4 + 4 + len(p.PieceData))
}

// NewPieceResponse return  a object
func NewPieceResponse() *PieceResponse {
	br := &PieceResponse{}
	return br
}

// EncodeBody can encode client message to binary
func (p *PieceResponse) EncodeBody() ([]byte, error) {
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
	err = binary.Write(buff, binary.BigEndian, p.PiceeIndex)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buff, binary.BigEndian, p.PieceHash)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buff, binary.BigEndian, p.PieceSize)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buff, binary.BigEndian, uint32(len(p.PieceData)))
	if err != nil {
		return nil, err
	}
	_, err = buff.Write(([]byte)(p.PieceData))
	if err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}

// DecodeBody can decode binary code to struct
func (p *PieceResponse) DecodeBody(buff *bytes.Buffer) (*PieceResponse, error) {
	//buff := bytes.NewBuffer(buf)
	var err error
	err = binary.Read(buff, binary.BigEndian, &p.Head.ProtocolVersion)
	if err != nil {
		return nil, err
	}
	err = binary.Read(buff, binary.BigEndian, &p.Head.CommandID)
	if err != nil {
		return nil, err
	}
	err = binary.Read(buff, binary.BigEndian, &p.Head.BodyLength)
	if err != nil {
		return nil, err
	}
	err = binary.Read(buff, binary.BigEndian, &p.PiceeIndex)
	if err != nil {
		return nil, err
	}
	err = binary.Read(buff, binary.BigEndian, &p.PieceHash)
	if err != nil {
		return nil, err
	}
	err = binary.Read(buff, binary.BigEndian, &p.PieceSize)
	if err != nil {
		return nil, err
	}

	valuedata := make([]byte, p.Head.BodyLength-12)
	err = binary.Read(buff, binary.BigEndian, &valuedata)
	if err != nil {
		return nil, err
	}
	p.PieceData = string(valuedata)
	return p, nil
}

// PrintfPieceResponse function printf PieceResponse
func (p *PieceResponse) PrintfPieceResponse(pieceIndex uint32, bodyLength uint16) {
	log.Printf("[S2C pieceResponse] Index:%v  Hash:%v Length:%v", pieceIndex, p.PieceHash, bodyLength-12)
}
