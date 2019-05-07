package platinumc

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func confirmIPAndPort(ip string, port int) bool {
	count := strings.Count(ip, ".")
	if count < 3 || count > 3 {
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
func TestCheckTask(t *testing.T) {

	str := "255.0.0.1"
	port := 6600
	result := confirmIPAndPort(str, port)
	if result {
		fmt.Println("ok")

	} else {

		fmt.Println("sorry")
	}
}
