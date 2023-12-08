package longport

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type longportConfig struct {
	AppKey      *string `cty:"app_key"`
	AppSecret   *string `cty:"app_secret"`
	AccessToken *string `cty:"access_token"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"app_key": {
		Type: schema.TypeString,
	},
	"app_secret": {
		Type: schema.TypeString,
	},
	"access_token": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &longportConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) longportConfig {
	if connection == nil || connection.Config == nil {
		return longportConfig{}
	}
	config, _ := connection.Config.(longportConfig)
	return config
}
