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

	Task *Task
}

// NewConnectionTCP function can set value for ConnectionTCP struct object
func NewConnectionTCP(t *Task) *ConnectionTCP {
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
	return br
}

// Connect to create a link and return a net.conn and error
func (c *ConnectionTCP) Connect() (net.Conn, error) {
	conn, err := net.Dial("tcp", c.ServerAddress)
	if err != nil {
		return nil, fmt.Errorf("connect failed: %v", err)
	}
	return conn, nil
}
