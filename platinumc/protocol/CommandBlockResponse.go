package protocol

import (
	"bytes"
	"encoding/binary"

	"github.com/QIYUEKURONG/platinumc/platinumc"
)

// BlockResponse receive message from server
type BlockResponse struct {
	Head             Header
	FileIndex        string
	FileOffset       uint64
	FileSize         uint64
	FilelastModified uint64
}

// NewObject can create a new object
func (b *BlockResponse) NewObject(task *platinumc.Task) *BlockResponse {
	message := new(BlockResponse)
	message.Head.ProtocolVersion = ProtocolVersion
	message.Head.CommandID = 002
	message.Head.BodyLength = b.GetBodyLength(*message)
	message.FileIndex = task.FileIndex
	message.FileOffset = (uint64)(task.BlockIndex)
	message.FilelastModified = 0
	return message
}

// GetBodyLength get body length
func (b *BlockResponse) GetBodyLength(mess BlockResponse) uint16 {
	return (uint16)(4 + len(([]rune)(mess.FileIndex)) + 8 + 8 + 8)
}

// EncodeBody can encode client message to binary
func (b *BlockResponse) EncodeBody(message *BlockResponse) ([]byte, error) {
	buf := new(bytes.Buffer) // bytes.Buffer是一个缓冲byte类型的缓冲器
	var err error
	err = binary.Write(buf, binary.BigEndian, message.Head.ProtocolVersion)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.BigEndian, message.Head.CommandID)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.BigEndian, message.Head.BodyLength)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.BigEndian, ([]byte)(message.FileIndex))
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.BigEndian, message.FileOffset)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.BigEndian, message.FileSize)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.BigEndian, message.FilelastModified)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), err
}

// DecodeBody can decode binary code to struct
func (b *BlockResponse) DecodeBody(buf []byte) (BlockResponse, error) {
	var block BlockResponse

	// read head
	buff := bytes.NewBuffer(buf)
	err := binary.Read(buff, binary.BigEndian, &block.Head.ProtocolVersion)
	if err != nil {
		return block, err
	}
	err = binary.Read(buff, binary.BigEndian, &block.Head.CommandID)
	if err != nil {
		return block, err
	}
	err = binary.Read(buff, binary.BigEndian, &block.Head.BodyLength)
	if err != nil {
		return block, err
	}
	//read body
	var fileindexlen uint32
	err = binary.Read(buff, binary.BigEndian, &fileindexlen)
	if err != nil {
		return block, err
	}
	valueindex := make([]byte, fileindexlen)
	err = binary.Read(buff, binary.BigEndian, &valueindex)
	block.FileIndex = string(valueindex)

	err = binary.Read(buff, binary.BigEndian, &block.FileOffset)
	if err != nil {
		return block, err
	}
	err = binary.Read(buff, binary.BigEndian, &block.FileSize)
	if err != nil {
		return block, err
	}
	err = binary.Read(buff, binary.BigEndian, &block.FilelastModified)
	if err != nil {
		return block, err
	}
	return block, nil
}
