package dpufalco


import (
	"strings"

	"github.com/falcosecurity/plugin-sdk-go/pkg/sdk"
)

func (k *Plugin) Fields() []sdk.FieldEntry {
	return []sdk.FieldEntry{
		{Type: "string", Name: "docker.id", Display: "Container Id", Desc: "Container Id"},
		{Type: "string", Name: "docker.name", Display: "Container Name", Desc: "Container Name"},
		{Type: "string", Name: "docker.evt.type", Display: "Evt Type", Desc: "Event Type, e.g. 'read"},
		{Type: "string", Name: "docker.evt.args", Display: "Evt Args", Desc: "Event Args"},
		{Type: "string", Name: "docker.proc.name", Display: "Proc Name", Desc: "Process Name"},
		{Type: "string", Name: "docker.proc.pname", Display: "Proc PName", Desc: "Parent Process Name"},
	}
}

func GetField(data string, field string) (bool, string) {
	data_slice := strings.Split(data, "\t")
	data_slice_len := len(data_slice)
	if data_slice_len < 6 {
		return false, ""
	}

	switch field {
	case "docker.id":
		return true, data_slice[0]
	case "docker.name":
		return true, data_slice[1]
	case "docker.evt.type":
		return true, data_slice[2]
	case "docker.evt.args":
		return true, data_slice[3]
	case "docker.proc.name":
		return true, data_slice[4]
	case "docker.proc.pname":
		return true, data_slice[5]
	}

	return false, ""
}


func (k *Plugin) Extract(req sdk.ExtractRequest, evt sdk.EventReader) error {
	// fmt.Println("Extract")
	return k.ExtractFromEvent(req, evt)
}

func (k *Plugin) ExtractFromEvent(req sdk.ExtractRequest, evt sdk.EventReader) error {
	data := make([]byte, 1024)
	reader := evt.Reader()
	_, err := reader.Read(data)
	if err != nil {
		return err
	}

	evtStr := string(data)
	// fmt.Printf("data=%s\n", data)
	
	present, value := GetField(evtStr, req.Field())
	if present {
		req.SetValue(value)
	}
	return nil
}
