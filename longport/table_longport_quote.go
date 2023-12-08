package longport

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableQuote(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "longport_quote",
		Description: "Real-time Quotes Of Securities.",
		List: &plugin.ListConfig{
			Hydrate:    listQuote,
			KeyColumns: plugin.SingleColumn("symbol"),
		},
		Columns: quoteColumns("pre_market_quote", "post_market_quote"),
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
