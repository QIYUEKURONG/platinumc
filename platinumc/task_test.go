package platinumc

import (
	"fmt"
	"os"
	"testing"
)

func JSONSer(message *BlockRequest) ([]byte, error) {
	/*
		message.MclientID = "001"
		message.MclientType = 4
		message.MfileIndex = "./index"
		message.MfileOffset = 0
		message.Head.MprotocolVersion = 1
		message.Head.McommandID = 002
		num := (int)(len(message.MclientID) + len(message.MfileIndex))
		message.Head.MbodyLength = (int16)(1 + 4 + num + 4 + 8)

		data, err := json.Marshal(message)
		return data, err*/
	return nil, nil
}

func TestCheckTask(t *testing.T) {
	var message BlockRequest

	data, err := JSONSer(&message)
	if err != nil {
		os.Exit(-1)
	}
	fmt.Println(data)

}
