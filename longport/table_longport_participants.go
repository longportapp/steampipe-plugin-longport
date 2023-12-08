package longport

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableParticipants(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "longport_participants",
		Description: "Get Broker IDs.",
		List: &plugin.ListConfig{
			Hydrate: listParticipants,
		},
		Columns: participantColumns(),
	}
}

func listParticipants(ctx context.Context, d *plugin.QueryData, p *plugin.HydrateData) (interface{}, error) {
	quoteContext, err := connect(ctx, d)

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
