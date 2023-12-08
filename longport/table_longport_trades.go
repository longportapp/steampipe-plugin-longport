package longport

import (
	"context"
	"errors"

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
		Columns: tradeColumns(),
	}
}

func listTrades(ctx context.Context, d *plugin.QueryData, p *plugin.HydrateData) (interface{}, error) {
	quoteContext, err := connect(ctx, d)

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
		d.StreamListItem(ctx, trade)
	}

	return nil, nil
}

func tradeColumns(optionalCols ...string) []*plugin.Column {
	cols := []*plugin.Column{
		// Top columns
		{Name: "symbol", Type: proto.ColumnType_STRING, Transform: transform.FromField("Symbol"), Description: "Symbol"},
		{Name: "trades", Type: proto.ColumnType_JSON, Transform: transform.FromField("trades").Transform(transformTrades), Description: "Price"},
	}

	return cols
}

func transformTrades(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	items := []map[string]interface{}{}

	if d.Value == nil {
		return items, nil
	}

	trades, ok := d.Value.([]*quote.Trade)
	if !ok {
		return items, errors.New("transformBrokers failed")
	}

	for _, t := range trades {
		var item = map[string]interface{}{}
		if t != nil {
			item["price"] = t.Price
			item["volume"] = t.Volume
			item["timestamp"] = t.Timestamp
			item["trade_type"] = t.TradeType
			item["direction"] = t.Direction
			item["trade_session"] = t.TradeSession
		}
		items = append(items, item)
	}

	return items, nil
}
