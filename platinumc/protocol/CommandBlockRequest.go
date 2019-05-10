package protocol

import (
	"bytes"
	"encoding/binary"

	"github.com/QIYUEKURONG/platinumc/platinumc"
)

// BlockRequest record client_type,client_id and file_offset
type BlockRequest struct {
	Head       Header
	ClientType uint8
	ClientID   string
	FileIndex  string
	FileOffset uint64
}

// NewObject can create a new object
func (br *BlockRequest) NewObject(task *platinumc.Task) *BlockRequest {
	message := new(BlockRequest)
	message.ClientID = "001"
	message.ClientType = ClientTypeWindows
	message.FileIndex = task.FileIndex
	message.FileOffset = (uint64)(task.BlockIndex)
	message.Head.ProtocolVersion = ProtocolVersion
	message.Head.CommandID = 002
	message.Head.BodyLength = br.GetBodyLength(*message)
	return message
}

// GetBodyLength get body length
func (br *BlockRequest) GetBodyLength(mess BlockRequest) uint16 {
	//var = unsafe.Sizeof(br.ClientID)
	return (uint16)(1 + 4 + len(([]rune)(mess.ClientID)) + 4 + len(([]rune)(mess.FileIndex)) + 8)
}

// EncodeBody can encode client message to binary
func (br *BlockRequest) EncodeBody(message *BlockRequest) ([]byte, error) {
	buf := new(bytes.Buffer)
	var err error
	// write head
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
	// write body
	err = binary.Write(buf, binary.BigEndian, message.ClientType)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.BigEndian, uint32(len(message.ClientID)))
	if err != nil {
		return nil, err
	}
	_, err = buf.Write([]byte(message.ClientID))
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.BigEndian, uint32(len(message.FileIndex)))
	if err != nil {
		return nil, err
	}
	_, err = buf.Write([]byte(message.FileIndex))
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.BigEndian, message.FileOffset)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// DecodeBody can decode binary code to struct
func (br *BlockRequest) DecodeBody(buf []byte) (BlockRequest, error) {
	var block BlockRequest
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
	err = binary.Read(buff, binary.BigEndian, &block.ClientType)
	if err != nil {
		return block, err
	}
	var clientidlen uint32
	err = binary.Read(buff, binary.BigEndian, &clientidlen)
	if err != nil {
		return block, err
	}

	valueid := make([]byte, clientidlen)
	err = binary.Read(buff, binary.BigEndian, &valueid)
	if err != nil {
		return block, err
	}
	block.ClientID = string(valueid)

	var fileindexlen uint32
	err = binary.Read(buff, binary.BigEndian, &fileindexlen)
	if err != nil {
		return block, err
	}
	valueoffset := make([]byte, fileindexlen)
	err = binary.Read(buff, binary.BigEndian, &valueoffset)
	if err != nil {
		return block, err
	}
	block.FileIndex = string(valueoffset)

	err = binary.Read(buff, binary.BigEndian, &block.FileOffset)
	if err != nil {
		return block, err
	}
	return block, nil
}
