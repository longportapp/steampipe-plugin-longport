package longport

import (
	"context"
	"errors"

	"github.com/longportapp/openapi-go/quote"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableCandlesticks(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "longport_candlesticks",
		Description: "Get Security Trades.",
		List: &plugin.ListConfig{
			Hydrate:    listCandlesticks,
			KeyColumns: plugin.AnyColumn([]string{"symbol", "period", "adjust_type"}),
		},
		Columns: []*plugin.Column{
			{Name: "symbol", Type: proto.ColumnType_STRING, Transform: transform.FromField("Symbol"), Description: "Symbol"},
			{Name: "close", Type: proto.ColumnType_STRING, Transform: transform.FromField("Candlestick.Close"), Description: "Close price"},
			{Name: "open", Type: proto.ColumnType_STRING, Transform: transform.FromField("Candlestick.Open"), Description: "Open Price"},
			{Name: "low", Type: proto.ColumnType_STRING, Transform: transform.FromField("Candlestick.Low"), Description: "Low Price"},
			{Name: "high", Type: proto.ColumnType_STRING, Transform: transform.FromField("Candlestick.High"), Description: "High Price"},
			{Name: "volume", Type: proto.ColumnType_INT, Transform: transform.FromField("Candlestick.Volume"), Description: "Volume"},
			{Name: "turnover", Type: proto.ColumnType_STRING, Transform: transform.FromField("Candlestick.Turnover"), Description: "Turnover"},
			{Name: "timestamp", Type: proto.ColumnType_INT, Transform: transform.FromField("Candlestick.Timestamp"), Description: "Timestamp"},
			{Name: "period", Type: proto.ColumnType_INT, Transform: transform.FromField("Period"), Description: "Period"},
			{Name: "adjust_type", Type: proto.ColumnType_INT, Transform: transform.FromField("AdjustType"), Description: "AdjustType"},
		},
	}
}

type Candlestick struct {
	Symbol      string
	Candlestick *quote.Candlestick
	Period      int
	AdjustType  int
}

func listCandlesticks(ctx context.Context, d *plugin.QueryData, p *plugin.HydrateData) (interface{}, error) {
	quoteContext, err := connect(ctx, d)

	if err != nil {
		return nil, err
	}

	symbol, err := symbolString(ctx, d, p)
	if err != nil {
		return nil, err
	}

	period, err := equalInt(ctx, d, p, "period")
	if err != nil {
		return nil, err
	}
	if period == 0 {
		period = 1
	}
	var period1 = quote.Period(period)

	adjust_type, err := equalInt(ctx, d, p, "adjust_type")
	if err != nil {
		return nil, err
	}
	if adjust_type != 0 && adjust_type != 1 {
		return nil, errors.New("adjust_type must be 0 or 1")
	}

	var adjust_type1 = quote.AdjustType(adjust_type)

	limit := queryLimit(d)

	items, err := quoteContext.Candlesticks(ctx, symbol, period1, limit, adjust_type1)
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		var info = Candlestick{
			Symbol:      symbol,
			Candlestick: item,
			Period:      period,
			AdjustType:  adjust_type,
		}

		d.StreamListItem(ctx, info)
	}

	return nil, nil
}
