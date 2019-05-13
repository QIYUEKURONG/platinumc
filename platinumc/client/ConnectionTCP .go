package client

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/QIYUEKURONG/platinumc/platinumc"
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
}

// NewObject function can set value for ConnectionTCP struct object
func (c *ConnectionTCP) NewObject(t *platinumc.Task, object *ConnectionTCP) error {
	index := strings.Index(t.ServerAddress, ":")
	object.IP = t.ServerAddress[0:index]
	Port := t.ServerAddress[index+1 : len(t.ServerAddress)-1]
	value, err := strconv.Atoi(Port)
	if err != nil {
		fmt.Println("ConnectionTCP'S  NewObject convert string to int error")
		return err
	}
	object.Port = (uint16)(value)
	object.ServerAddress = t.ServerAddress
	return nil
}

// Connect to create a link and return a net.conn and error
func (c *ConnectionTCP) Connect(co *ConnectionTCP) (net.Conn, error) {
	fmt.Println(co.ServerAddress)
	Conn, err := net.Dial("tcp", co.ServerAddress)
	return Conn, err
}
