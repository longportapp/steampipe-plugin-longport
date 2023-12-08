package longport

import (
	"context"
	"errors"

	"github.com/longportapp/openapi-go/quote"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableDepth(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "longport_depth",
		Description: "Security Depth.",
		List: &plugin.ListConfig{
			Hydrate:    listDepth,
			KeyColumns: plugin.SingleColumn("symbol"),
		},
		// https://github.com/longportapp/openapi-go/blob/main/quote/types.go#L340C6-L340C19
		Columns: []*plugin.Column{
			// Top columns
			{Name: "symbol", Type: proto.ColumnType_STRING, Transform: transform.FromField("Symbol"), Description: "Security code"},
			{Name: "ask", Type: proto.ColumnType_JSON, Transform: transform.FromField("Ask").Transform((transformDepth)), Description: "Ask depth"},
			{Name: "bid", Type: proto.ColumnType_JSON, Transform: transform.FromField("Bid").Transform((transformDepth)), Description: "Bid depth"},
		},
	}
}

type Depth struct {
	Symbol string
	Ask    *quote.Depth
	Bid    *quote.Depth
}

func listDepth(ctx context.Context, d *plugin.QueryData, p *plugin.HydrateData) (interface{}, error) {
	quoteContext, err := connect(ctx, d)

	if err != nil {
		return nil, err
	}

	symbol, err := symbolString(ctx, d, p)
	if err != nil {
		return nil, err
	}

	info, err := quoteContext.Depth(ctx, symbol)
	if err != nil {
		return nil, err
	}

	for i, ask := range info.Ask {
		var depth = Depth{
			Symbol: symbol,
			Ask:    ask,
		}
		if len(info.Bid) > i {
			depth.Bid = info.Bid[i]
		}

		d.StreamListItem(ctx, depth)
	}

	return nil, nil
}

// https://github.com/longportapp/openapi-go/blob/main/quote/types.go#L133
func transformDepth(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	item := map[string]interface{}{}

	if d.Value == nil {
		return item, nil
	}

	t, ok := d.Value.(*quote.Depth)
	if !ok {
		return item, errors.New("transformDepth failed")
	}

	if t != nil {
		item["position"] = t.Position
		item["price"] = t.Price
		item["volume"] = t.Volume
		item["order_num"] = t.OrderNum
	}

	return item, nil
}
