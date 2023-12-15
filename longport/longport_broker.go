package longport

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableLongPortBroker(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "longport_broker",
		Description: "Get Security Brokers.",
		List: &plugin.ListConfig{
			Hydrate:    listLongPortBroker,
			KeyColumns: plugin.SingleColumn("symbol"),
		},
		Columns: brokerColumns(),
	}
}

func listLongPortBroker(ctx context.Context, d *plugin.QueryData, p *plugin.HydrateData) (interface{}, error) {
	quoteContext, err := getQuoteContext(ctx, d)

	if err != nil {
		return nil, err
	}

	symbol, err := symbolString(ctx, d, p)
	if err != nil {
		return nil, err
	}

	info, err := quoteContext.Brokers(ctx, symbol)
	if err != nil {
		return nil, err
	}

	d.StreamListItem(ctx, info)

	return nil, nil
}
