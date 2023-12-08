package longport

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableOptionQuote(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "longport_option_quote",
		Description: "Real-time Quotes Of Option Securities.",
		List: &plugin.ListConfig{
			Hydrate:    listOptionQuote,
			KeyColumns: plugin.SingleColumn("symbol"),
		},
		Columns: quoteColumns("pre_market_quote", "post_market_quote"),
	}
}

func listOptionQuote(ctx context.Context, d *plugin.QueryData, p *plugin.HydrateData) (interface{}, error) {
	quoteContext, err := getQuoteContext(ctx, d)

	if err != nil {
		return nil, err
	}

	symbols, err := symbolList(ctx, d, p)
	if err != nil {
		return nil, err
	}

	infos, err := quoteContext.OptionQuote(ctx, symbols)
	if err != nil {
		return nil, err
	}

	for _, info := range infos {
		d.StreamListItem(ctx, info)
	}

	return nil, nil
}
