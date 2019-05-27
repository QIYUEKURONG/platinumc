package protocol

import (
	"bytes"
	"encoding/binary"
	"log"
)

// BlockResponse receive message from server
type BlockResponse struct {
	Head             Header
	FileIndex        string
	FileOffset       uint64
	FileSize         uint64
	FilelastModified uint64
}

// NewBlockResponse can create a new object
func NewBlockResponse() *BlockResponse {
	br := &BlockResponse{}
	return br
}

// GetBodyLength get body length
func (b *BlockResponse) GetBodyLength() uint16 {
	return (uint16)(4 + len(([]rune)(b.FileIndex)) + 8 + 8 + 8)
}

// Encode can encode client message to binary
func (b *BlockResponse) Encode() ([]byte, error) {
	buf := new(bytes.Buffer) // bytes.Buffer是一个缓冲byte类型的缓冲器
	var err error
	//Head
	err = binary.Write(buf, binary.BigEndian, b.Head.ProtocolVersion)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.BigEndian, b.Head.CommandID)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.BigEndian, b.Head.BodyLength)
	if err != nil {
		return nil, err
	}
	//Body
	err = binary.Write(buf, binary.BigEndian, uint32(len(b.FileIndex)))
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.BigEndian, ([]byte)(b.FileIndex))
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.BigEndian, b.FileOffset)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.BigEndian, b.FileSize)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.BigEndian, b.FilelastModified)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), err
}

// DecodeBody can decode binary code to struct
func (b *BlockResponse) DecodeBody(buff *bytes.Buffer) (*BlockResponse, error) {

	b = NewBlockResponse()
	//buff := bytes.NewBuffer(buf)
	err := binary.Read(buff, binary.BigEndian, &b.Head.ProtocolVersion)
	if err != nil {
		return nil, err
	}
	err = binary.Read(buff, binary.BigEndian, &b.Head.CommandID)
	if err != nil {
		return nil, err
	}
	err = binary.Read(buff, binary.BigEndian, &b.Head.BodyLength)
	if err != nil {
		return nil, err
	}
	//read body
	var fileindexlen uint32
	err = binary.Read(buff, binary.BigEndian, &fileindexlen)
	if err != nil {
		return nil, err
	}
	valueindex := make([]byte, fileindexlen)
	err = binary.Read(buff, binary.BigEndian, &valueindex)
	b.FileIndex = string(valueindex)

	err = binary.Read(buff, binary.BigEndian, &b.FileOffset)
	if err != nil {
		return nil, err
	}
	err = binary.Read(buff, binary.BigEndian, &b.FileSize)
	if err != nil {
		return nil, err
	}
	err = binary.Read(buff, binary.BigEndian, &b.FilelastModified)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// PrintfBlockResponse function to printf fileindex,filesize,fileoffset and filelastmodified
func (b *BlockResponse) PrintfBlockResponse() {
	log.Printf("[S2C BlockResponse] Fileindex:%s Offset:%v FileSize:%v LastModified:%v\n",
		b.FileIndex, b.FileSize, b.FileOffset, b.FilelastModified)
}
