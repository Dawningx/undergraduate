package dpufalco


import (
	"encoding/json"

	"github.com/falcosecurity/plugin-sdk-go/pkg/sdk"
	"github.com/falcosecurity/plugin-sdk-go/pkg/sdk/plugins"
	"github.com/falcosecurity/plugin-sdk-go/pkg/sdk/symbols/extract"
)

const pluginName = "dpufalco"

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
		Version:     "0.1.0",
		EventSource: "dpufalco",
	}
}

func (k *Plugin) Init(cfg string) error {
	k.Config.Reset()
	err := json.Unmarshal([]byte(cfg), &k.Config)
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