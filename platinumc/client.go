package platinumc

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"os"
)

// AssignmentStruct can set value to every object
func AssignmentStruct(message *BlockRequest, task *Task) {

	message.ClientID = "001"
	message.ClientType = 4
	message.FileIndex = task.FileIndex
	message.FileOffset = (uint64)(task.BlockIndex)
	message.Head.ProtocolVersion = 1
	message.Head.CommandID = 002
	num := (int)(len(message.ClientID) + len(message.FileIndex))
	message.Head.BodyLength = (uint16)(1 + 4 + num + 4 + 8)
}

// SerializateBinary can that make object to binary
func SerializateBinary(message *BlockRequest) ([]byte, error) {
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

/* type BlockResponse struct {
	Head             Header
	FileIndex        string
	FileOffset       uint64
	FileSize         uint64
	FilelastModified uint64
} */

// UnserializateBinary can taht data of server from binary to struct
func UnserializateBinary(buf []byte) (BlockResponse, error) {
	//var servermess BlockResponse
	var data BlockResponse

	// read head
	buff := bytes.NewBuffer(buf)
	err := binary.Read(buff, binary.BigEndian, &data.Head.ProtocolVersion)
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
	//read body
	var messlength uint32
	err = binary.Read(buff, binary.BigEndian, &messlength)
	if err != nil {
		return data, err
	}
	value := make([]byte, messlength)
	err = binary.Read(buff, binary.BigEndian, &value)
	if err != nil {
		return data, err
	}
	data.FileIndex = string(value)
	err = binary.Read(buff, binary.BigEndian, &data.FileOffset)
	if err != nil {
		return data, err
	}
	err = binary.Read(buff, binary.BigEndian, &data.FileSize)
	if err != nil {
		return data, err
	}
	err = binary.Read(buff, binary.BigEndian, &data.FilelastModified)
	return data, nil

}

// Run the client to download file specefied in task.
func Run(task *Task) {
	//1：首先去绑定服务器的ip和port
	conn, err := net.Dial("tcp", task.ServerAddress)
	if err != nil {
		fmt.Println("Error Connection", err)
		os.Exit(-1)
	}
	defer conn.Close()

	// start to get date frmo server
	for {
		//转换成二进制 然后发送给服务器
		var message BlockRequest

		AssignmentStruct(&message, task)
		buf, err := SerializateBinary(&message)
		if err != nil {
			fmt.Println("sorry! Object convert binary error")
			os.Exit(-1)
			fmt.Println(buf)

			buff := make([]byte, 0)
			// 循环去接收服务器发送来的数据
			for {
				_, err := conn.Read(buff)
				if err != nil {
					fmt.Println("client read error")
					os.Exit(-1)
				}
				UnserializateBinary(buff)

			}

		}
	}
}
