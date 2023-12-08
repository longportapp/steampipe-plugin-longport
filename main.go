package main

import (
	"github.com/longportapp/steampipe-plugin-longport/longport"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: longport.Plugin})
}
