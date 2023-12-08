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
			"longport_depth":         tableDepth(ctx),
			"longport_brokers":       tableBrokers(ctx),
			"longport_participants":  tableParticipants(ctx),
			"longport_trades":        tableTrades(ctx),
			"longport_intraday":      tableIntraday(ctx),
			"longport_candlesticks":  tableCandlesticks(ctx),
			// Trade
			"longport_history_executions": longport_history_executions(ctx),
			"longport_today_executions":   longport_today_executions(ctx),
			"longport_history_orders":     longport_history_orders(ctx),
			"longport_today_orders":       longport_today_orders(ctx),
		},
	}
	return p
}
