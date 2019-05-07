package main

import (
	"flag"
	"fmt"

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

func main() {
	task := platinumc.Task{}
	parseCommandLineArguemtns(&task)
	if !checkCommandLineArguments(&task) {
		flag.PrintDefaults()
	}

	// Parse command line arguemnts
	// Check task
	//	platinumc.Run(&task)

}
