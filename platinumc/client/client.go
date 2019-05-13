package client

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"os"

	"github.com/QIYUEKURONG/platinumc/platinumc"
	"github.com/QIYUEKURONG/platinumc/platinumc/protocol"
)

// GetBodyLength function get body length
func GetBodyLength(buff []byte) (uint16, error) {
	//解析长度
	bodylengthbuff := buff[2:4]
	body := bytes.NewBuffer(bodylengthbuff)
	var bodylength uint16
	var err error
	err = binary.Read(body, binary.BigEndian, &bodylength)
	return bodylength, err
}

// GetCommandType function get  object type
func GetCommandType(buff []byte) (uint8, error) {
	commandid := buff[1:2]
	id := bytes.NewBuffer(commandid)
	var commandtype uint8
	err := binary.Read(id, binary.BigEndian, &commandtype)
	return commandtype, err
}

// Run the client to download file specefied in task.
func Run(task *platinumc.Task) {
	tcplink := new(ConnectionTCP)
	var err error
	err = tcplink.NewObject(task, tcplink)
	fmt.Println("tcplink ", tcplink)
	if err != nil {
		fmt.Println("tcplink object create error")
		os.Exit(-1)
	}
	var conn net.Conn
	conn, err = tcplink.Connect(tcplink)
	if err != nil {
		fmt.Println("sorry! Connect server error")
		os.Exit(-1)
	}
	// 先创建一个CommandBlockRequest的结构
	CommandBlock := new(protocol.BlockRequest)
	CommandBlock.NewObject(task, CommandBlock)
	sendbuff, err := CommandBlock.EncodeBody(CommandBlock)
	if err != nil {
		fmt.Println("Sorry!,The CommandBlock EncodeBody error")
		os.Exit(-1)
	}
	fmt.Println(sendbuff)
	fmt.Println("start to send")
	conn.Write(sendbuff)

	/* objects, _ := CommandBlock.DecodeBody(sendbuff)
	fmt.Println(objects) */

	//创建一个缓冲区，用于存放服务器发送来的数据
	blockbuff := make([]byte, 512)
	var Response *(protocol.BlockResponse)
	// start to get date frmo server
	for {
		_, err = conn.Read(blockbuff)
		if err != nil {
			fmt.Println("Sorry! read error")
			os.Exit(-1)
		}
		fmt.Println("read sucess ", blockbuff)
		if len(blockbuff) < 4 {
			continue
		}
		//解析长度
		bodylength, err1 := GetBodyLength(blockbuff)
		if err1 != nil {
			fmt.Println("sorry ! to read body length error!")
			os.Exit(-1)
		}
		if bodylength > (uint16)(len(blockbuff)-4) {
			continue
		}
		fmt.Println("bodylength ", bodylength)
		//解析类型
		var commandtype uint8
		message := blockbuff[0:(4 + bodylength)]
		fmt.Println("message ", message)
		commandtype, err = GetCommandType(message)
		if err != nil {
			fmt.Println("Sorry ! binary convert error")
			os.Exit(-1)
		}
		comtype := fmt.Sprintf("%x", commandtype)
		fmt.Println("comtype ", comtype)
		switch comtype {
		case "07":
			fmt.Println("sorry! no have the file")
			fin := new(protocol.CommandFin)
			*fin, err = fin.DecodeBody(message)
			if err != nil {
				fmt.Println("sorry ! Fin Decode error")
				continue
			} else {
				fmt.Println("erron is:", fin.ErrorCode)
				os.Exit(-1)
			}
		case "31":
			Response = new(protocol.BlockResponse)
			*Response, err = Response.DecodeBody(message)
			if err != nil {
				fmt.Println("sorry ! Response Decode error")
				continue
			} else {
				blockbuff = blockbuff[(4+bodylength)+1:]
				fmt.Println(len(blockbuff), cap(blockbuff))
			}
		}
		break
	} //跳出for循环。说明服务器有这个数据
	fmt.Println("the file exits", *Response)

	//建一个Piece结构体 要开始接收数据了
	piecebuff := make([]byte, 1024)
	piecereq := new(protocol.PieceRequest)
	piecereq.NewObject(task, piecereq)
	sedbuff, err3 := piecereq.EncodeBody(piecereq)
	if err3 != nil {
		fmt.Println("Sorry ! the piece encode error")
		os.Exit(-1)
	}
	conn.Write(sedbuff)
	//接收数据
	conn.Read(piecebuff)
	fmt.Println("piece buff", piecebuff)
	fmt.Println(len(piecebuff))
	/* pieceres := new(protocol.PieceResponse)
	*pieceres, err = pieceres.DecodeBody(piecebuff)

	fmt.Printf("%s", pieceres.PieceData) */

}

//转换成二进制 然后发送给服务器
/* 	var message protocol.BlockRequest
AssignmentStruct(&message, task)
buf, err := SerializateBinary(&message)
if err != nil {
	fmt.Println("sorry! Object convert binary error")
	os.Exit(-1)
}
fmt.Println(buf)

//buff := make([]byte, 0)
//fmt.Println(buff)
// 循环去接收服务器发送来的数据
buff := make([]byte, 0)
_, err1 := conn.Read(buff)
if err1 != nil {
	fmt.Println("client read error")
	os.Exit(-1)
}
data, err1 := UnserializateBinary(buff)
if err1 != nil {
	fmt.Println("sorry! UnserializateBinary  error")
}
//写入文件里面
file.Write(([]byte)(data.FileIndex))
//如果文件的尺寸不是0的话，就继续请求文件
if data.FileSize != 0 {
	task.FileIndex = data.FileIndex
	task.BlockIndex = (uint)(data.FileOffset)
} else {
	fmt.Println("file download sucess")
	break

}
*/
