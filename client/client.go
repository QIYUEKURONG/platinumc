package client

import (
	"flag"
	"fmt"
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

// ParseTask function can parse task and check task
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
		return nil, fmt.Errorf("failed!ServerAddress is empty")
	}
	if t.FileIndex == "" {
		return nil, fmt.Errorf("failed! FileIndex is empty")
	}
	return t, nil
}

// Run the client to download file specefied in task.
func Run(task *Task) {
	switch task.TransportProtocol {
	case "tcp":
		StartTCP(task)
	case "wss":
		StartWSS(task)
	}
}
