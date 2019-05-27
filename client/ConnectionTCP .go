package client

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

// ConnectionTCP can record local ip and port and peer ip and port
type ConnectionTCP struct {
	// Ip record peer ip
	IP string
	// Port record peer port
	Port uint16
	//LocalIp record local ip
	LocalIP string
	//Localport record local port
	LocalPort uint16
	// ReceiveBuffer record receive buff
	ReceiveBuffer []byte
	ServerAddress string
	Task          *Task
}

// NewConnectionTCP function can set value for ConnectionTCP struct object
func NewConnectionTCP(t *Task) (net.Conn, error) {
	index := strings.Index(t.ServerAddress, ":")
	ip := t.ServerAddress[0:index]
	value := t.ServerAddress[index+1 : len(t.ServerAddress)-1]
	port, _ := (strconv.Atoi(value))
	ports := (uint16)(port)
	br := &ConnectionTCP{
		IP:            ip,
		Port:          ports,
		ServerAddress: t.ServerAddress,
		Task:          t,
	}
	if ParseServerIPAndPort(br.ServerAddress) != nil {
		return nil, fmt.Errorf("server ip or port find error")
	}
	conn, err := net.Dial("tcp", br.ServerAddress)
	if err != nil {
		return nil, fmt.Errorf("connect failed: %v", err)
	}
	return conn, nil
}

//confirmIpAndPort find error
func confirmIPAndPort(ip string, port int) bool {
	count := strings.Count(ip, ".")
	if count != 3 {
		return false
	}
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
	if (port < 0) || (port > 65535) {
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
		return fmt.Errorf("Use error:  sorry! server port convert error: %s %v", ip, port)
	}
	if !confirmIPAndPort(ip, port) {
		return fmt.Errorf("call confirmIPAndPort error：%v", err)
	}
	return nil
}
