package longport

import (
	"context"

	"github.com/longportapp/openapi-go/quote"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableIntraday(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "longport_intraday",
		Description: "Get Security Intraday.",
		List: &plugin.ListConfig{
			Hydrate:    listIntraday,
			KeyColumns: plugin.SingleColumn("symbol"),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "symbol", Type: proto.ColumnType_STRING, Transform: transform.FromField("Symbol"), Description: "Symbol"},
			{Name: "price", Type: proto.ColumnType_STRING, Transform: transform.FromField("Line.Price"), Description: "Close price of the minute"},
			{Name: "volume", Type: proto.ColumnType_STRING, Transform: transform.FromField("Line.Volume"), Description: "Close price of the minute"},
			{Name: "turnover", Type: proto.ColumnType_STRING, Transform: transform.FromField("Line.Turnover"), Description: ""},
			{Name: "avg_price", Type: proto.ColumnType_STRING, Transform: transform.FromField("Line.AvgPrice"), Description: ""},
			{Name: "timestamp", Type: proto.ColumnType_INT, Transform: transform.FromField("Line.Timestamp"), Description: "Start time stamp of the minute"},
		},
	}
}

type Intrady struct {
	Symbol string
	Line   *quote.IntradayLine
}

func listIntraday(ctx context.Context, d *plugin.QueryData, p *plugin.HydrateData) (interface{}, error) {
	quoteContext, err := getQuoteContext(ctx, d)

	if err != nil {
		return nil, err
	}

	symbol, err := symbolString(ctx, d, p)
	if err != nil {
		return nil, err
	}

	lines, err := quoteContext.Intraday(ctx, symbol)
	if err != nil {
		return nil, err
	}

	for _, line := range lines {
		var info = Intrady{
			Symbol: symbol,
			Line:   line,
		}

		d.StreamListItem(ctx, info)
	}

	return nil, nil
}
