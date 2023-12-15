package longport

import (
	"context"

	"github.com/longportapp/openapi-go/trade"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableLongPortHistoryExecution(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "longport_history_execution",
		Description: "History Executions.",
		List: &plugin.ListConfig{
			Hydrate:    listLongPortHistoryExecution,
			KeyColumns: plugin.OptionalColumns([]string{"symbol"}),
		},
		Columns: []*plugin.Column{
			{Name: "symbol", Type: proto.ColumnType_STRING, Transform: transform.FromField("Symbol"), Description: "Symbol of security."},
			{Name: "order_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("OrderId"), Description: "Order ID."},
			{Name: "trade_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("TradeId"), Description: "Trade ID."},
			{Name: "trade_done_at", Type: proto.ColumnType_STRING, Transform: transform.FromField("TradeDoneAt"), Description: "Trade done time, formatted as a timestamp (second)."},
			{Name: "quantity", Type: proto.ColumnType_INT, Transform: transform.FromField("Quantity"), Description: "Executed quantity."},
			{Name: "price", Type: proto.ColumnType_INT, Transform: transform.FromField("Price"), Description: "Executed price."},
		},
	}
}

func listLongPortHistoryExecution(ctx context.Context, d *plugin.QueryData, p *plugin.HydrateData) (interface{}, error) {
	context, err := getTradeContext(ctx, d)

	if err != nil {
		return nil, err
	}

	symbol, err := symbolString(ctx, d, p)
	if err != nil {
		return nil, err
	}

	items, err := context.HistoryExecutions(ctx, &trade.GetHistoryExecutions{
		Symbol: symbol,
	})
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		d.StreamListItem(ctx, item)
	}

	return nil, nil
}
