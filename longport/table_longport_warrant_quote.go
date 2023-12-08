package longport

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableWarrantQuote(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "longport_warrant_quote",
		Description: "Real-time Quotes Of Warrant Securities.",
		List: &plugin.ListConfig{
			Hydrate:    listWarrantQuote,
			KeyColumns: plugin.SingleColumn("symbol"),
		},
		Columns: quoteColumns("warrant_extend"),
	}
}

func listWarrantQuote(ctx context.Context, d *plugin.QueryData, p *plugin.HydrateData) (interface{}, error) {
	quoteContext, err := connect(ctx, d)

	if err != nil {
		return nil, err
	}

	symbols, err := symbolList(ctx, d, p)
	if err != nil {
		return nil, err
	}

	infos, err := quoteContext.WarrantQuote(ctx, symbols)
	if err != nil {
		return nil, err
	}

	for _, info := range infos {
		d.StreamListItem(ctx, info)
	}

	return nil, nil
}
