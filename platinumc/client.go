package platinumc

import (
	"fmt"
	"net"
	"os"
)

// Run the client to download file specefied in task.
func Run(task *Task) {
	//1：首先去绑定服务器的ip和port
	conn, err := net.Dial("tcp", task.ServerAddress)
	if err != nil {
		fmt.Println("Error Connection", err)
		os.Exit(-1)
	}
	defer conn.Close()

	file, err := os.OpenFile("filename", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("sorry file create error")
	}
	defer file.Close()

	// start to get date frmo server
	for {
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
	}
}
