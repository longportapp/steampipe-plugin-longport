package longport

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type longportConfig struct {
	AppKey      *string `hcl:"app_key"`
	AppSecret   *string `hcl:"app_secret"`
	AccessToken *string `hcl:"access_token"`
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
