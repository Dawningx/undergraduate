package dpufalco


import (
	"strings"
	"net"
	"strconv"

	"github.com/falcosecurity/plugin-sdk-go/pkg/sdk"
	"github.com/falcosecurity/plugin-sdk-go/pkg/sdk/plugins/source"
)

const (
	dpu_ip = "192.168.1.101"
	dpu_port = 5000
)

func (k *Plugin) Open(param string) (source.Instance, error) {
	evtChan := make(chan source.PushEvent)
	res := &source.PushEvent{}

	address := dpu_ip + ":" + strconv.Itoa(dpu_port)
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return nil, err
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return nil, err
	}

	data := make([]byte, 1024)
	n, sendaddr, err := conn.ReadFromUDP(data)
	if err != nil {
		return nil, err
	}

	res.Data = data
	evtChan <- *res
		
	return source.NewPushInstance(evtChan, source.WithInstanceClose(func() {  conn.Close()  }))
}
