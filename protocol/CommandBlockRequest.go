package protocol

import (
	"bytes"
	"encoding/binary"
)

// BlockRequest record client_type,client_id and file_offset
type BlockRequest struct {
	Head       Header
	ClientType uint8
	ClientID   string
	FileIndex  string
	FileOffset uint64
}

// NewBlockRequest allocate and initialize a new instance of BlockRequest.
func NewBlockRequest(clientID string, clientType uint8, fileIndex string, fileOffset uint64) *BlockRequest {
	br := &BlockRequest{
		ClientID:   clientID,
		ClientType: clientType,
		FileIndex:  fileIndex,
		FileOffset: fileOffset,
	}
	br.Head.ProtocolVersion = ProtocolVersion
	br.Head.CommandID = CommandBlockRequest
	//br.Head.BodyLength = br.GetBodyLength()

	return br
}

// GetBodyLength get body length
func (br *BlockRequest) GetBodyLength() uint16 {
	return (uint16)(1 + 4 + len(([]rune)(br.ClientID)) + 4 + len(([]rune)(br.FileIndex)) + 8)
}

// Encode can encode client message to binary
func (br *BlockRequest) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	var err error
	// write head
	err = binary.Write(buf, binary.BigEndian, br.Head.ProtocolVersion)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.BigEndian, br.Head.CommandID)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.BigEndian, br.GetBodyLength())
	if err != nil {
		return nil, err
	}
	// write body
	err = binary.Write(buf, binary.BigEndian, br.ClientType)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.BigEndian, uint32(len(br.ClientID)))
	if err != nil {
		return nil, err
	}
	_, err = buf.Write([]byte(br.ClientID))
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.BigEndian, uint32(len(br.FileIndex)))
	if err != nil {
		return nil, err
	}
	_, err = buf.Write([]byte(br.FileIndex))
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.BigEndian, br.FileOffset)
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
