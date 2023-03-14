package dpufalco


import (
	"encoding/json"

	"github.com/falcosecurity/plugin-sdk-go/pkg/sdk"
	"github.com/falcosecurity/plugin-sdk-go/pkg/sdk/plugins"
	"github.com/alecthomas/jsonschema"
)

const pluginName = "dpufalco"

type PluginConfig struct {
	Jitter uint64 `json:"jitter" jsonschema:"title:Sample jitter,description=empty,default=0"`
}

type Plugin struct {
	plugins.BasePlugin
	config PluginConfig
}

func (k *Plugin) Info() *plugins.Info {
	return &plugins.Info{
		ID:	         999,
		Name:        pluginName,
		Description: "Set up a UDP connection and listen to it",
		Contact:     "github.com/Dawningx/undergraduate",
		Version:     "1.0.0",
		EventSource: "dpufalco",
	}
}

func (k *Plugin) Init(cfg string) error {
	// k.config.Reset()
	err := json.Unmarshal([]byte(cfg), &k.config)
	if err != nil {
		return err
	}

	return nil
}

func (p *Plugin) InitSchema() *sdk.SchemaInfo {
	reflector := jsonschema.Reflector{
		// all properties are optional by default
		RequiredFromJSONSchemaTags: true,
		// unrecognized properties don't cause a parsing failures
		AllowAdditionalProperties: true,
	}
	if schema, err := reflector.Reflect(&PluginConfig{}).MarshalJSON(); err == nil {
		return &sdk.SchemaInfo{
			Schema: string(schema),
		}
	}
	return nil
}
