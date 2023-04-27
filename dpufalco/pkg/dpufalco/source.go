package dpufalco

/*
#cgo LDFLAGS: -L. -ldpushmem

#include <stdlib.h>
#include "dpushmem.h"
*/
import "C"
import (

	"github.com/falcosecurity/plugin-sdk-go/pkg/sdk/plugins/source"
)


func (k *Plugin) Open(param string) (source.Instance, error) {
	evtChan := make(chan source.PushEvent)
	var res int
	// sum_c := 0

	go func() {
		for {
			res = int(C.set_signal())
			if res != 0 {
				continue
			}
			res = int(C.wait_for_buffer())
			if res != 0 {
				continue
			}
			for i:=0; i<(102400 / 256); i++ {
				var j int
				bp := C.read_buffer(C.int(i))
				buffer := C.GoString(bp)
				data := make([]byte, 256)
				for j=0; j<len(buffer); j++ {
					if buffer[j] != 0 {
						data[j] = buffer[j]
					} else {
						break
					}
				}
				if j > 0 {
					evtChan <- source.PushEvent{Data: data}
				}
			}
		}
	}()

		
	return source.NewPushInstance(evtChan, source.WithInstanceClose(func() {}))
}
