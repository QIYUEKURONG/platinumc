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

// GetCommandID function get  object type
func GetCommandID(buff []byte) (uint8, error) {
	commandid := buff[1:2]
	id := bytes.NewBuffer(commandid)
	var commandtype uint8
	err := binary.Read(id, binary.BigEndian, &commandtype)
	return commandtype, err
}

// SendBlockRequest encode and send BlockRequest to server.
func SendBlockRequest(task *platinumc.Task, conn net.Conn) error {
	// 1. 构造
	br := protocol.NewBlockRequest(task.Identifier, protocol.ClientTypeUnknown, task.FileIndex, 0)
	// 2. 编码
	sendBuffer, err := br.Encode()
	if err != nil {
		return fmt.Errorf("encode failed: %v", err)
	}
	// 3. 发送
	_, err = conn.Write(sendBuffer)
	if err != nil {
		return fmt.Errorf("send failed: %v", err)
	}
	fmt.Printf("[C2S BlockRequest] FileIndex:%s ClientId:tcp-%s FileOffset:%v\n", br.FileIndex, br.ClientID, br.FileOffset)
	return nil
}

// SendPieceRequest encode and send PieceRequest to server.
func SendPieceRequest(t *platinumc.Task, index uint32, conn net.Conn) error {
	// 1: 构造
	br := protocol.NewPieceRequest(t, index)
	//fmt.Println("br value ", br)
	// 2：编码
	sendBuffer, err := br.Encode()
	if err != nil {
		return fmt.Errorf("encode failed: %v", err)
	}
	//fmt.Println(sendBuffer)
	// 3:发送
	_, err = conn.Write(sendBuffer)
	if err != nil {
		return fmt.Errorf("send failed: %v", err)
	}
	fmt.Printf("[C2S PieceRequest] Index:%v Offset:%v Length:%v\n", index, br.PiecenIndex, 8192)
	return nil
}

// Run the client to download file specefied in task.
func Run(task *platinumc.Task) {
	// 1. 创建应用层连接，负责处理消息编解
	c := NewConnectionTCP(task)
	conn, err := c.Connect()
	if err != nil {
		fmt.Println("call Connect failed: %v\n", err)
		os.Exit(-1)
	}
	err = SendBlockRequest(task, conn)
	if err != nil {
		fmt.Printf("call SendBlockRequest failed: %v\n", err)
		os.Exit(1)
	}
	//创建一个缓冲区，用于存放服务器发送来的消息
	recvBuff := make([]byte, 1024)
	dataBuff := bytes.NewBuffer([]byte{})
	//databuff := make([]byte, protocol.MAXDATA)
	// start to get date frmo server
	for {
		lengthSize, err := conn.Read(recvBuff)
		fmt.Println("recvBuff and lengthSize", recvBuff, lengthSize)
		t := recvBuff[0:lengthSize]
		dataBuff.Write(t)
		fmt.Println("databuff ", dataBuff.Bytes(), dataBuff.Len())
		if lengthSize == 0 {
			break
		}
		if err != nil {
			fmt.Println("Sorry! read error")
			os.Exit(-1)
		}
		if (dataBuff.Len()) < 4 {
			continue
		}
		//解析长度
		bodyLength, err := GetBodyLength(dataBuff.Bytes())
		fmt.Println("bodyLength valus", bodyLength)
		if err != nil {
			fmt.Println("sorry ! to read body length error!")
			os.Exit(-1)
		}
		if bodyLength > (uint16)((dataBuff.Len())-4) {
			continue
		}
		//解析类型
		var commandID uint8
		commandID, err = GetCommandID(dataBuff.Bytes())
		if err != nil {
			fmt.Println("Sorry ! binary convert error")
			os.Exit(-1)
		}
		fmt.Println("commandtype ", commandID)
		// 选择类型
		switch commandID {
		case protocol.CommandFin1:
			fmt.Println("sorry! no have the file")
			fin := protocol.NewFinObject()
			fin, err := fin.DecodeBody(dataBuff)
			if err != nil {
				fmt.Println("sorry ! Fin Decode error")
				continue
			} else {
				fmt.Println("erron is:", fin.ErrorCode)
				os.Exit(-1)
			}
		case protocol.CommandBlockResponse:
			Response := protocol.NewBlockResponse()
			Response, err := Response.DecodeBody(dataBuff)
			if err != nil {
				fmt.Println("sorry ! Response Decode error")
				os.Exit(-1)
			}
			fmt.Printf("[S2C BlockResponse] Fileindex:%s Offset:%v FileSize:%v LastModified:%v\n",
				Response.FileIndex, Response.FileSize, Response.FileOffset, Response.FilelastModified)
			err = SendPieceRequest(task, 0, conn)
			if err != nil {
				fmt.Printf("call SendPieceRequest failed: %v\n", err)
			}
		case protocol.CommandPieceResponse:
			pieceResponse := protocol.NewPieceResponse()
			pieceResponse, err = pieceResponse.DecodeBody(dataBuff)
			if err != nil {
				fmt.Println("sorry ! the pieceresponse decode error")
				os.Exit(-1)
			}
			fmt.Printf("PieceData %s \n", pieceResponse.PieceData)
		}

		//fmt.Println("file ")

	}
}

//跳出for循环。说明服务器有这个数据
/* fmt.Println("the file exits", *Response)

//建一个Piece结构体 要开始接收数据了
//piecebuff := make([]byte, 5048)
piecereq := new(protocol.PieceRequest)
piecereq.NewObject(task, piecereq)
sedbuff, err3 := piecereq.EncodeBody(piecereq)
if err3 != nil {
	fmt.Println("Sorry ! the piece encode error")
	os.Exit(-1)
}
conn.Write(sedbuff)
//接收数据
piecebuff := make([]byte, 18000)
fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!")
count, _ := conn.Read(piecebuff)
fmt.Println("count", count)
fmt.Println("piece buff", piecebuff)
fmt.Println(len(piecebuff))
*/
/* pieceres := new(protocol.PieceResponse)
*pieceres, err = pieceres.DecodeBody(piecebuff)
fmt.Printf("%s", pieceres.PieceData) */

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
