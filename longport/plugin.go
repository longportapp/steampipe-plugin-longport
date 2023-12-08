package longport

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-longport",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		DefaultTransform: transform.FromGo().NullIfZero(),
		TableMap: map[string]*plugin.Table{
			"longport_static_info":   tableStaticInfo(ctx),
			"longport_quote":         tableQuote(ctx),
			"longport_option_quote":  tableOptionQuote(ctx),
			"longport_warrant_quote": tableWarrantQuote(ctx),
		},
	}
	return p
}
