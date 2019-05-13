package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/QIYUEKURONG/platinumc/platinumc"
	"github.com/QIYUEKURONG/platinumc/platinumc/client"
)

func parseCommandLineArguemtns(t *platinumc.Task) {
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
}

func checkCommandLineArguments(t *platinumc.Task) bool {
	if t.TransportProtocol == "" {
		fmt.Println("failed! TransportProtocol is empty!")
		//flag.CommandLine()build
		return false
	}
	if t.ServerAddress == "" {
		fmt.Println("failed!ServerAddress is empty!")
		return false
	}
	if t.FileIndex == "" {
		fmt.Println("failed! FileIndex is empty!")
		return false
	}
	return true
}

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
func ParseServerIPAndPort(t platinumc.Task) bool {

	index := strings.Index(t.ServerAddress, ":")
	var ip string

	if index == -1 {
		fmt.Println(" Use error:  sorry! server ip or port find error,please try again")
		return false
	} else {
		//分析格式
		ip = t.ServerAddress[0:index]
		port, err := strconv.Atoi(t.ServerAddress[index+1 : len(t.ServerAddress)-1])
		if err != nil {
			fmt.Println("Use error:  sorry! server port convert error", ip, port)
			return false
		}
		if !confirmIPAndPort(ip, port) {
			return false
		}
	}
	return true
}

func main() {
	task := platinumc.Task{}
	parseCommandLineArguemtns(&task)
	if !checkCommandLineArguments(&task) {
		flag.PrintDefaults()
		return
	}
	if !ParseServerIPAndPort(task) {
		fmt.Println("Use error : sorry! your ip and port find error ")
		return
	}
	fmt.Println(task)
	client.Run(&task)

}
