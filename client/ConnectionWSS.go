package client

import (
	"flag"
	"fmt"
	"strings"

	"github.com/gorilla/websocket"
)

/*
type URL struct {
        Scheme     string
        Opaque     string    // encoded opaque data
        User       *Userinfo // username and password information
        Host       string    // host or host:port
        Path       string    // path (relative paths may omit leading slash)
        RawPath    string    // encoded path hint (see EscapedPath method); added in Go 1.5
        ForceQuery bool      // append a query ('?') even if RawQuery is empty; added in Go 1.7
        RawQuery   string    // encoded query values, without '?'
        Fragment   string    // fragment for references, without '#'
}*/

// ConnectionWSS record websocket's ip and port ...
type ConnectionWSS struct {
	// Ip record peer ip
	IP string
	// Port record peer port
	Port uint16
	//LocalIp record local ip
	LocalIP string
	//Localport record local port
	LocalPort     uint16
	ServerAddress string
	Task          *Task
}

var addr = flag.String("addr", "localhost:12345", "http service address")

func addrToURL(addr string) string {
	// "192.168.200.26:59843" => "192-168-200-26:59843"
	addr = strings.Replace(addr, ".", "-", -1)

	// 192-168-200-26:59843" => "192-168-200-26.nhost.00cdn.com:59843"
	addr = strings.Replace(addr, ":", ".nhost.00cdn.com:", 1)

	// "192-168-200-26.nhost.00cdn.com:59843" => "wss://192-168-200-26.nhost.00cdn.com:59843"
	return fmt.Sprintf("wss://%s", addr)
}

// NewConnectionWSS function to create a new connection wss
func NewConnectionWSS(t *Task) (*websocket.Conn, error) {
	//u := url.URL{Scheme: "wss", Host: t.ServerAddress, Path: t.FileIndex}
	//fmt.Println(u)
	var dialer *websocket.Dialer
	conn, _, err := dialer.Dial(addrToURL(t.ServerAddress), nil)
	if err != nil {
		return nil, fmt.Errorf("To create a new websocket error :%v", err)
	}
	return conn, nil
}

// SendMessage can use websocket send message
func SendMessage(conn *websocket.Conn, data []byte) error {
	err := conn.WriteMessage(0, data)
	if err != nil {
		return fmt.Errorf("RecvMessage function RecvMessage error: %v", err)
	}
	return nil
}

// RecvMessage can use websocket recv message
func RecvMessage(conn *websocket.Conn) ([]byte, error) {
	_, buff, err := conn.ReadMessage()
	if err != nil {
		return nil, fmt.Errorf("RecvMessage function RecvMessage error: %v", err)
	}
	return buff, err
}

//./platinumc -x xianone -t tcp -a 192.168.200.26:59606 -f txvideo.ippzone.com/zyvd/4d/5b/c094-6b2f-11e9-9e1d-00163e020689 -b 0 -vc -o 1.mp4
