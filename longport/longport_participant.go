package longport

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableLongPortParticipant(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "longport_participant",
		Description: "Get Broker IDs.",
		List: &plugin.ListConfig{
			Hydrate: listLongPortParticipant,
		},
		Columns: participantColumns(),
	}
}

func listLongPortParticipant(ctx context.Context, d *plugin.QueryData, p *plugin.HydrateData) (interface{}, error) {
	quoteContext, err := getQuoteContext(ctx, d)

	if err != nil {
		return nil, err
	}

	infos, err := quoteContext.Participants(ctx)
	if err != nil {
		return nil, err
	}

	for _, info := range infos {
		d.StreamListItem(ctx, info)
	}

	return nil, nil
}
