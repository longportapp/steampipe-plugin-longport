package longport

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableDepth(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "longport_depth",
		Description: "Security Depth.",
		List: &plugin.ListConfig{
			Hydrate:    listDepth,
			KeyColumns: plugin.SingleColumn("symbol"),
		},
		Columns: depthColumns(),
	}
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

	d.StreamListItem(ctx, info)

	return nil, nil
}
