package main

import (
	"log"

	"github.com/QIYUEKURONG/platinumc/client"
)

func main() {
	// 1. 解析下载任务
	task, err := client.ParseTask()
	if err != nil {
		log.Fatalf("parse task failed: %v", err)
	}

	// 2. 执行下载任务
	client.Run(task)
}
