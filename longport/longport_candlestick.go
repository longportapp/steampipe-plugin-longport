package longport

import (
	"context"
	"errors"

	"github.com/longportapp/openapi-go/quote"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableLongPortCandlestick(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "longport_candlestick",
		Description: "Get Security Trades.",
		List: &plugin.ListConfig{
			Hydrate:    listLongPortCandlestick,
			KeyColumns: plugin.OptionalColumns([]string{"symbol", "period", "adjust_type"}),
		},
		Columns: []*plugin.Column{
			{Name: "symbol", Type: proto.ColumnType_STRING, Transform: transform.FromField("Symbol"), Description: "Symbol of security."},
			{Name: "close", Type: proto.ColumnType_STRING, Transform: transform.FromField("Candlestick.Close"), Description: ""},
			{Name: "open", Type: proto.ColumnType_STRING, Transform: transform.FromField("Candlestick.Open"), Description: ""},
			{Name: "low", Type: proto.ColumnType_STRING, Transform: transform.FromField("Candlestick.Low"), Description: ""},
			{Name: "high", Type: proto.ColumnType_STRING, Transform: transform.FromField("Candlestick.High"), Description: ""},
			{Name: "volume", Type: proto.ColumnType_INT, Transform: transform.FromField("Candlestick.Volume"), Description: ""},
			{Name: "turnover", Type: proto.ColumnType_STRING, Transform: transform.FromField("Candlestick.Turnover"), Description: ""},
			{Name: "timestamp", Type: proto.ColumnType_INT, Transform: transform.FromField("Candlestick.Timestamp"), Description: "Timestamp of the candlestick."},
			{Name: "period", Type: proto.ColumnType_INT, Transform: transform.FromField("Period"), Description: ""},
			{Name: "adjust_type", Type: proto.ColumnType_INT, Transform: transform.FromField("AdjustType"), Description: ""},
		},
	}
}

type Candlestick struct {
	Symbol      string
	Candlestick *quote.Candlestick
	Period      int
	AdjustType  int
}

func listLongPortCandlestick(ctx context.Context, d *plugin.QueryData, p *plugin.HydrateData) (interface{}, error) {
	quoteContext, err := getQuoteContext(ctx, d)

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
