package client

import (
	"fmt"
	"net/url"
	"strconv"
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

// NewConnectionWSS function to create a new connection wss
func NewConnectionWSS(t *Task) (*websocket.Conn, error) {
	index := strings.Index(t.ServerAddress, ":")
	ip := t.ServerAddress[0:index]
	value := t.ServerAddress[index+1 : len(t.ServerAddress)-1]
	port, _ := (strconv.Atoi(value))
	ports := (uint16)(port)
	br := &ConnectionWSS{
		IP:            ip,
		Port:          ports,
		ServerAddress: t.ServerAddress,
		Task:          t,
	}
	u := url.URL{Scheme: "ws", Host: br.ServerAddress, Path: br.Task.SavePath}
	var dialer *websocket.Dialer
	conn, _, err := dialer.Dial(u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("To create a new websocket error :%v", err)
	}
	return conn, nil
}

//./platinumc -x xianone -t tcp -a 192.168.200.26:59606 -f txvideo.ippzone.com/zyvd/4d/5b/c094-6b2f-11e9-9e1d-00163e020689 -b 0 -vc -o 1.mp4
