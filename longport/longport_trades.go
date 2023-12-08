package longport

import (
	"context"

	"github.com/longportapp/openapi-go/quote"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableTrades(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "longport_trades",
		Description: "Get Security Trades.",
		List: &plugin.ListConfig{
			Hydrate:    listTrades,
			KeyColumns: plugin.SingleColumn("symbol"),
		},
		Columns: []*plugin.Column{
			{Name: "symbol", Type: proto.ColumnType_STRING, Transform: transform.FromField("Symbol"), Description: "Symbol"},
			{Name: "price", Type: proto.ColumnType_STRING, Transform: transform.FromField("Trade.Price"), Description: "Price"},
			{Name: "volume", Type: proto.ColumnType_INT, Transform: transform.FromField("Trade.Volume"), Description: "Volume"},
			{Name: "trade_type", Type: proto.ColumnType_STRING, Transform: transform.FromField("Trade.TradeType"), Description: "TradeType"},
			{Name: "direction", Type: proto.ColumnType_STRING, Transform: transform.FromField("Trade.Direction"), Description: "Trade direction, 0 - neutral, 1 - down, 2 - up"},
			{Name: "trade_session", Type: proto.ColumnType_INT, Transform: transform.FromField("Trade.TradeSession"), Description: ""},
			{Name: "timestamp", Type: proto.ColumnType_INT, Transform: transform.FromField("Trade.Timestamp"), Description: "Time of trading"},
		},
	}
}

type Trade struct {
	Symbol string
	Trade  *quote.Trade
}

func listTrades(ctx context.Context, d *plugin.QueryData, p *plugin.HydrateData) (interface{}, error) {
	quoteContext, err := getQuoteContext(ctx, d)

	if err != nil {
		return nil, err
	}

	symbol, err := symbolString(ctx, d, p)
	if err != nil {
		return nil, err
	}

	limit := int32(20)
	if d.QueryContext.Limit != nil {
		limit = int32(*d.QueryContext.Limit)
	}

	trades, err := quoteContext.Trades(ctx, symbol, limit)
	if err != nil {
		return nil, err
	}

	for _, trade := range trades {
		var info = Trade{
			Symbol: symbol,
			Trade:  trade,
		}

		d.StreamListItem(ctx, info)
	}

	return nil, nil
}
