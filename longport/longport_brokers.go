package longport

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableBrokers(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "longport_brokers",
		Description: "Get Security Brokers.",
		List: &plugin.ListConfig{
			Hydrate:    listBrokers,
			KeyColumns: plugin.SingleColumn("symbol"),
		},
		Columns: brokerColumns(),
	}
}

func listBrokers(ctx context.Context, d *plugin.QueryData, p *plugin.HydrateData) (interface{}, error) {
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
