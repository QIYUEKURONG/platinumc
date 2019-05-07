package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/QIYUEKURONG/platinumc/platinumc"
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

// ConfirmIpAndPort find error
func ConfirmIpAndPort(ip string, port int) bool {

	var num int
	count := strings.Count(ip, ".")
	if count != 3 {
		return false
	} else {
		mess := strings.Split(ip, ".")
		for _, data := range mess {
			num, _ := strconv.Atoi(data)
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

func ParseServerIpAndPort(t *platinumc.Task) bool {

	index := strings.Index(t.ServerAddress, ":")
	var ip string

	if index == -1 {
		fmt.Println("sorry! server ip or port find error,please try again")
		return false
	} else {
		//分析格式
		ip = t.ServerAddress[0:index]
		port, err := strconv.Atoi(t.ServerAddress[index+1 : len(t.ServerAddress)-1])
		if err != nil {
			fmt.Println("sorry! server port convert error", ip, port)
			return false
		}
		if !ConfirmIpAndPort(ip, port) {
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
	if !ParseServerIpAndPort(&task) {
		fmt.Println("Use error : sorry! your ip and port find error ")
		return
	}

	platinumc.Run(&task)

}
