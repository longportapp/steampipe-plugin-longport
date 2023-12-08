package longport

import (
	"context"
	"errors"

	"github.com/longportapp/openapi-go/quote"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableQuote(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "longport_quote",
		Description: "This API is used to obtain the real-time quotes of securities, and supports all types of securities.",
		List: &plugin.ListConfig{
			Hydrate:    listQuote,
			KeyColumns: plugin.SingleColumn("symbol"),
		},
		Columns: quoteColumns(),
	}
}

func listQuote(ctx context.Context, d *plugin.QueryData, p *plugin.HydrateData) (interface{}, error) {
	quoteContext, err := connect(ctx, d)

	if err != nil {
		return nil, err
	}

	symbols, err := symbolList(ctx, d, p)
	if err != nil {
		return nil, err
	}

	infos, err := quoteContext.Quote(ctx, symbols)
	if err != nil {
		return nil, err
	}

	for _, info := range infos {
		d.StreamListItem(ctx, info)
	}

	return nil, nil
}

// https://github.com/longportapp/openapi-go/blob/main/quote/types.go#L313
func quoteColumns() []*plugin.Column {
	cols := []*plugin.Column{
		// Top columns
		{Name: "symbol", Type: proto.ColumnType_STRING, Transform: transform.FromField("Symbol"), Description: "Security code"},
		{Name: "last_done", Type: proto.ColumnType_STRING, Transform: transform.FromField("LastDone"), Description: "Latest price"},
		{Name: "prev_close", Type: proto.ColumnType_STRING, Transform: transform.FromField("PrevClose"), Description: "Previous closing price"},
		{Name: "open", Type: proto.ColumnType_STRING, Transform: transform.FromField("Open"), Description: "Opening price"},
		{Name: "high", Type: proto.ColumnType_STRING, Transform: transform.FromField("High"), Description: "Highest price"},
		{Name: "low", Type: proto.ColumnType_STRING, Transform: transform.FromField("Low"), Description: "Lowest price"},
		{Name: "volume", Type: proto.ColumnType_INT, Transform: transform.FromField("Volume"), Description: "Volume"},
		{Name: "turnover", Type: proto.ColumnType_STRING, Transform: transform.FromField("Turnover"), Description: "Turnover"},
		{Name: "timestamp", Type: proto.ColumnType_INT, Transform: transform.FromField("Timestamp"), Description: "Time of latest price"},
		{Name: "trade_status", Type: proto.ColumnType_INT, Transform: transform.FromField("TradeStatus"), Description: "Security trading status, see TradeStatus"},
		// https://github.com/longportapp/openapi-go/blob/main/quote/types.go#L329
		{Name: "pre_market_quote", Type: proto.ColumnType_JSON, Transform: transform.FromField("PreMarketQuote").Transform(transformPrePostQuote), Description: "Pre-market quote"},
		{Name: "post_market_quote", Type: proto.ColumnType_JSON, Transform: transform.FromField("PostMarketQuote").Transform(transformPrePostQuote), Description: "After-hours quote"},
	}
	return cols
}

func transformPrePostQuote(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	items := map[string]interface{}{}

	if d.Value == nil {
		return items, nil
	}

	t, ok := d.Value.(*quote.PrePostQuote)
	if !ok {
		return items, errors.New("convert quote.PrePostQuote failed")
	}
	if t != nil {
		items["last_done"] = t.LastDone
		items["prev_close"] = t.PrevClose
		items["high"] = t.High
		items["low"] = t.Low
		items["volume"] = t.Volume
		items["turnover"] = t.Turnover
		items["timestamp"] = t.Timestamp
	}

	return items, nil
}
