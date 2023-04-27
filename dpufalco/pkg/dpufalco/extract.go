package dpufalco


import (
	"strings"
	//"fmt"
	//"strconv"

	"github.com/falcosecurity/plugin-sdk-go/pkg/sdk"
)

func (k *Plugin) Fields() []sdk.FieldEntry {
	return []sdk.FieldEntry{
		{Type: "string", Name: "docker.id", Display: "Container Id", Desc: "Container Id"},
		{Type: "string", Name: "docker.fd.num", Display: "Fd Number", Desc: "Fd Number"},
		{Type: "string", Name: "docker.fd.type", Display: "Fd Type", Desc: "Fd Type"},
		{Type: "string", Name: "docker.proc.pname", Display: "Parent Process Name", Desc: "Parent Process Name"},
		{Type: "string", Name: "docker.proc.name", Display: "Process Name", Desc: "Process Name"},
		{Type: "string", Name: "docker.evt.type", Display: "Evt Type", Desc: "Event Type, e.g. read"},
		{Type: "string", Name: "docker.exe", Display: "Exe Name", Desc: "Exe Name"},
		{Type: "string", Name: "docker.cmdline", Display: "Command Line", Desc: "Command Line"},
	}
}

func GetField(data string, field string) (bool, string) {
	data_slice := strings.Split(data, "\t")
	data_slice_len := len(data_slice)
	if data_slice_len < 8 {
		fmt.Println("Wrong format!")
		// fmt.Println(strconv.Itoa(len(data)) + "[" + data + "]")
		return false, ""
	} else if data_slice_len > 8 {
		// fmt.Println("Packet coalesce!")
		return false, ""
	}

	switch field {
	case "docker.id":
		return true, data_slice[0]
	case "docker.fd.num":
		return true, data_slice[1]
	case "docker.fd.type":
		return true, data_slice[2]
	case "docker.proc.pname":
		return true, data_slice[3]
	case "docker.proc.name":
		return true, data_slice[4]
	case "docker.evt.type":
		return true, data_slice[5]
	case "docker.exe":
		return true, data_slice[6]
	case "docker.cmdline":
		return true, data_slice[7]
	}

	return false, ""
}


func (k *Plugin) Extract(req sdk.ExtractRequest, evt sdk.EventReader) error {
	// fmt.Println("Extract")
	return k.ExtractFromEvent(req, evt)
}

func (k *Plugin) ExtractFromEvent(req sdk.ExtractRequest, evt sdk.EventReader) error {
	data := make([]byte, 256)
	reader := evt.Reader()
	_, err := reader.Read(data)
	if err != nil {
		return err
	}

	evtStr := string(data)
	// fmt.Printf("%s\n", data)
	
	present, value := GetField(evtStr, req.Field())
	if present {
		req.SetValue(value)
	}
	return nil
}
