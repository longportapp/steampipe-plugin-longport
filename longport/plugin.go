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
		},
		DefaultTransform: transform.FromGo().NullIfZero(),
		TableMap: map[string]*plugin.Table{
			"longport_broker":            tableLongPortBroker(ctx),
			"longport_candlestick":       tableLongPortCandlestick(ctx),
			"longport_depth":             tableLongPortDepth(ctx),
			"longport_history_execution": tableLongPortHistoryExecution(ctx),
			"longport_history_order":     tableLongPortHistoryOrder(ctx),
			"longport_intraday":          tableLongPortIntraday(ctx),
			"longport_option_quote":      tableLongPortOptionQuote(ctx),
			"longport_participant":       tableLongPortParticipant(ctx),
			"longport_quote":             tableLongPortQuote(ctx),
			"longport_static_info":       tableLongPortStaticInfo(ctx),
			"longport_today_execution":   tableLongPortTodayExecution(ctx),
			"longport_today_order":       tableLongPortTodayOrder(ctx),
			"longport_trade":             tableLongPortTrade(ctx),
			"longport_warrant_quote":     tableLongPortWarrantQuote(ctx),
		},
	}
	return p
}
