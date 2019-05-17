package client

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/QIYUEKURONG/platinumc/protocol"
)

// Task can record Client task type
type Task struct {
	TransportProtocol string
	ServerAddress     string
	FileIndex         string
	BlockIndex        uint
	SavePath          string
	StartPieceIndex   uint
	FetchPieceCount   uint
	CheckPiece        bool
	VerboseLog        bool
	Identifier        string
}

func ParseTask() (*Task, error) {
	// 1. 解析参数
	t := &Task{}
	flag.StringVar(&t.Identifier, "x", "", "Client Identifier")
	flag.StringVar(&t.TransportProtocol, "t", "", "Transport Protocol (tcp, wss)")
	flag.StringVar(&t.ServerAddress, "a", "", "ServerAddress(tcp, wss, udp or mona address)")
	flag.StringVar(&t.FileIndex, "f", "", "FileIndex(must not be empty)")
	flag.UintVar(&t.BlockIndex, "b", 0, "BlockIndex,(default = 0)")
	flag.StringVar(&t.SavePath, "o", "", "SavePath")
	flag.UintVar(&t.StartPieceIndex, "s", 0, "StartPieceIndex")
	flag.UintVar(&t.FetchPieceCount, "n", 0, "FetchPieceCount")
	flag.BoolVar(&t.VerboseLog, "v", false, "(VerboseLog)")
	flag.BoolVar(&t.CheckPiece, "c", false, "(CheckPiece)")
	flag.Parse()

	// 2. 校验参数
	if t.TransportProtocol == "" {
		return nil, fmt.Errorf("TransportProtocol is empty")
	}
	if t.ServerAddress == "" {
		return nil, fmt.Errorf("failed!ServerAddress is empty!")
	}
	if t.FileIndex == "" {
		return nil, fmt.Errorf("failed! FileIndex is empty!")
	}

	return t, nil
}

/*
// confirmIpAndPort find error
func confirmIPAndPort(ip string, port int) bool {
	count := strings.Count(ip, ".")
	if count != 3 {
		return false
	} else {
		mess := strings.Split(ip, ".")
		for _, data := range mess {
			num, err := strconv.Atoi(data)
			if err != nil {
				return false
			}
			if num < 0 || num > 255 {
				return false
			}

		}
	}
	if port < 0 || port > 65535 {
		return false
	}

	return true
}

// ParseServerIPAndPort can get ip and port in right range
func ParseServerIPAndPort(addr string) error {
	index := strings.Index(addr, ":")
	if index == -1 {
		return fmt.Errorf("addr does not contain ':': %v", addr)
	}

	// 分析格式
	ip := addr[0:index]
	port, err := strconv.Atoi(addr[index+1 : len(addr)-1])
	if err != nil {
		return fmt.Errorf("Use error:  sorry! server port convert error", ip, port)
	}
	if !confirmIPAndPort(ip, port) {
		return false
	}

	return nil
}*/

// Run the client to download file specefied in task.
func Run(task *Task) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
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
	//  建立一个文件，存储最后的文件
	file, err := os.Create("map.mp4")
	if err != nil {
		fmt.Println("create a file error")
		os.Exit(-2)
	}
	// 为了最后打印做准备
	fileName := task.SavePath
	var fileLength uint64
	//var startTime int
	//var endTime int
	//创建一个缓冲区，用于存放服务器发送来的消息
	recvBuff := make([]byte, 1024)
	dataBuff := bytes.NewBuffer([]byte{})
	var pieceIndex uint32
	for {
		lengthSize, err := conn.Read(recvBuff)
		t := recvBuff[0:lengthSize]
		dataBuff.Write(t)
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
		//fmt.Println("bodyLength valus", bodyLength)
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
		// 选择类型
		switch commandID {
		case protocol.CommandFin1:
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
			Response, err = Response.DecodeBody(dataBuff)
			if err != nil {
				fmt.Println("sorry ! Response Decode error")
				os.Exit(-1)
			}
			str := SetLocalTime()
			fmt.Printf("[%s][S2C BlockResponse] Fileindex:%s Offset:%v FileSize:%v LastModified:%v\n",
				str, Response.FileIndex, Response.FileSize, Response.FileOffset, Response.FilelastModified)
			fileLength = Response.FileSize
			err = SendPieceRequest(task, pieceIndex, conn)
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
			file.Write(([]byte)(pieceResponse.PieceData))
			log.Printf("[S2C pieceResponse] Index:%v  Hash:%v Length:%v", pieceIndex, pieceResponse.PieceHash, bodyLength-12)
			pieceIndex++
			//说明文件已经接收完毕
			if (bodyLength - 12) < 8192 {
				log.Printf("[%s] Save file block to %v  FileSize %v", fileName, fileLength)
				err = SendFin(conn)
				if err != nil {
					fmt.Printf("call SendFin failed: %v\n", err)
				}
				break
			}
			err = SendPieceRequest(task, pieceIndex, conn)
		}
	}
}
