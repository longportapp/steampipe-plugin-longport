package longport

import (
	"context"
	"errors"

	"github.com/longportapp/openapi-go/quote"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

// https://github.com/longportapp/openapi-go/blob/main/quote/types.go#L313
func quoteColumns(optionalCols ...string) []*plugin.Column {
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
	}

	// https://github.com/longportapp/openapi-go/blob/main/quote/types.go#L329
	for _, col := range optionalCols {
		if col == "pre_market_quote" {
			cols = append(cols,
				[]*plugin.Column{
					{Name: "pre_market_quote", Type: proto.ColumnType_JSON, Transform: transform.FromField("PreMarketQuote").Transform(transformPrePostQuote), Description: "Pre-market quote"},
				}...,
			)
		}

		if col == "post_market_quote" {
			cols = append(cols,
				[]*plugin.Column{
					{Name: "post_market_quote", Type: proto.ColumnType_JSON, Transform: transform.FromField("PostMarketQuote").Transform(transformPrePostQuote), Description: "Pre-market quote"},
				}...,
			)
		}

		if col == "warrant_extend" {
			cols = append(cols,
				[]*plugin.Column{
					{Name: "warrant_extend", Type: proto.ColumnType_JSON, Transform: transform.FromField("WarrantExtend").Transform(transformWarrantExtend), Description: "Warrant information"},
				}...,
			)
		}
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

func transformWarrantExtend(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	items := map[string]interface{}{}

	if d.Value == nil {
		return items, nil
	}

	t, ok := d.Value.(*quote.WarrantExtend)
	if !ok {
		return items, errors.New("convert quote.PrePostQuote failed")
	}
	if t != nil {
		items["implied_volatility"] = t.ImpliedVolatility
		items["expiry_date"] = t.ExpiryDate
		items["last_trade_date"] = t.LastTradeDate
		items["outstanding_ratio"] = t.OutstandingRatio
		items["outstanding_qty"] = t.OutstandingQty
		items["conversion_ratio"] = t.ConversionRatio
		items["category"] = t.Category
		items["strike_price"] = t.StrikePrice
		items["upper_strike_price"] = t.UpperStrikePrice
		items["lower_strike_price"] = t.LowerStrikePrice
		items["underlying_symbol"] = t.UpperStrikePrice
	}

	return items, nil
}

// https://github.com/longportapp/openapi-go/blob/main/quote/types.go#L340C6-L340C19
func depthColumns(optionalCols ...string) []*plugin.Column {
	cols := []*plugin.Column{
		// Top columns
		{Name: "symbol", Type: proto.ColumnType_STRING, Transform: transform.FromField("Symbol"), Description: "Security code"},
		{Name: "ask", Type: proto.ColumnType_JSON, Transform: transform.FromField("Ask").Transform((transformDepth)), Description: "Ask depth"},
		{Name: "bid", Type: proto.ColumnType_JSON, Transform: transform.FromField("Bid").Transform((transformDepth)), Description: "Bid depth"},
	}

	return cols
}

// https://github.com/longportapp/openapi-go/blob/main/quote/types.go#L133
func transformDepth(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	items := []map[string]interface{}{}

	if d.Value == nil {
		return items, nil
	}

	depths, ok := d.Value.([]*quote.Depth)
	if !ok {
		return items, errors.New("transformDepth failed")
	}

	for _, t := range depths {
		var item = map[string]interface{}{}
		if t != nil {
			item["position"] = t.Position
			item["price"] = t.Price
			item["volume"] = t.Volume
			item["order_num"] = t.OrderNum
		}
		items = append(items, item)
	}

	return items, nil
}
