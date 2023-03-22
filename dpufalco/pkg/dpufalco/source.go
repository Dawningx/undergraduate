package dpufalco


import (
	"net"
	"strconv"
	"fmt"

	"github.com/falcosecurity/plugin-sdk-go/pkg/sdk/plugins/source"
)

const (
	dpu_ip = "127.0.0.1"
	dpu_port = 5000
)

func (k *Plugin) Open(param string) (source.Instance, error) {
	evtChan := make(chan source.PushEvent)

	go func() {
		address := dpu_ip + ":" + strconv.Itoa(dpu_port)
		addr, err := net.ResolveUDPAddr("udp", address)
		if err != nil {
			evtChan <- source.PushEvent{Err: err}
		}

		conn, err := net.ListenUDP("udp", addr)
		if err != nil {
			evtChan <- source.PushEvent{Err: err}
		}

		fmt.Println("UDP Socket Start Successfully!")
		defer conn.Close()
		for {
			data := make([]byte, 1024)
			_, _, err2 := conn.ReadFromUDP(data)
			fmt.Println(string(data))
		
			if err2 != nil {
				evtChan <- source.PushEvent{Err: err2}
			} else {
				evtChan <- source.PushEvent{Data: data}
			}
		}
	}()

		
	return source.NewPushInstance(evtChan, source.WithInstanceClose(func() {}))
}
