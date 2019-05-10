package protocol

import (
	"bytes"
	"encoding/binary"
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

// EncodeBody can encode client message to binary
func (p *PieceResponse) EncodeBody(message *PieceResponse) ([]byte, error) {
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
	err = binary.Write(buff, binary.BigEndian, message.PiceeIndex)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buff, binary.BigEndian, message.PieceHash)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buff, binary.BigEndian, message.PieceSize)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buff, binary.BigEndian, uint32(len(message.PieceData)))
	if err != nil {
		return nil, err
	}
	_, err = buff.Write(([]byte)(message.PieceData))
	if err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}

// DecodeBody can decode binary code to struct
func (p *PieceResponse) DecodeBody(buf []byte) (PieceResponse, error) {
	var piece PieceResponse
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
	err = binary.Read(buff, binary.BigEndian, &piece.PiceeIndex)
	if err != nil {
		return piece, err
	}
	err = binary.Read(buff, binary.BigEndian, &piece.PieceHash)
	if err != nil {
		return piece, err
	}
	err = binary.Read(buff, binary.BigEndian, &piece.PieceSize)
	if err != nil {
		return piece, err
	}
	var datalen uint32
	err = binary.Read(buff, binary.BigEndian, &datalen)
	if err != nil {
		return piece, err
	}
	valuedata := make([]byte, datalen)
	err = binary.Read(buff, binary.BigEndian, &valuedata)
	if err != nil {
		return piece, err
	}
	piece.PieceData = string(valuedata)
	return piece, nil
}
