package longport

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableStaticInfo(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "longport_static_info",
		Description: "Lookup Basic Information Of Securities.",
		List: &plugin.ListConfig{
			Hydrate:    listStaticInfo,
			KeyColumns: plugin.SingleColumn("symbol"),
		},
		Columns: securityInfoColumns(),
	}
}

func listStaticInfo(ctx context.Context, d *plugin.QueryData, p *plugin.HydrateData) (interface{}, error) {
	var logger = plugin.Logger(ctx).With("plugin", "longport_static_info")

	quoteContext, err := connect(ctx, d)
	logger.Info("quoteContext", quoteContext)

	if err != nil {
		logger.Error("connection_error", err)
		return nil, err
	}

	symbols, err := symbolList(ctx, d, p)
	if err != nil {
		logger.Error("where_error", err)
		return nil, err
	}

	logger.Info("symbols", symbols)

	infos, err := quoteContext.StaticInfo(ctx, symbols)
	if err != nil {
		logger.Error("query_error", err)
		return nil, err
	}

	for _, info := range infos {
		d.StreamListItem(ctx, info)
	}

	return infos, nil
}

// https://github.com/longportapp/openapi-go/blob/main/quote/types.go#L187
func securityInfoColumns() []*plugin.Column {
	cols := []*plugin.Column{
		// Top columns
		{Name: "symbol", Type: proto.ColumnType_STRING, Transform: transform.FromField("Symbol"), Description: "Security code"},
		{Name: "name_cn", Type: proto.ColumnType_STRING, Transform: transform.FromField("NameCn"), Description: "Security name (zh-CN)"},
		{Name: "name_en", Type: proto.ColumnType_STRING, Transform: transform.FromField("NameEn"), Description: "Security name (en)"},
		{Name: "name_hk", Type: proto.ColumnType_STRING, Transform: transform.FromField("NameHk"), Description: "Security name (zh-HK)"},
		{Name: "exchange", Type: proto.ColumnType_STRING, Transform: transform.FromField("Exchange"), Description: ""},
		{Name: "currency", Type: proto.ColumnType_STRING, Transform: transform.FromField("Currency"), Description: ""},
		{Name: "lot_size", Type: proto.ColumnType_INT, Transform: transform.FromField("LotSize"), Description: ""},
		{Name: "total_shares", Type: proto.ColumnType_INT, Transform: transform.FromField("TotalShares"), Description: ""},
		{Name: "circulating_shares", Type: proto.ColumnType_INT, Transform: transform.FromField("CirculatingShares"), Description: ""},
		{Name: "hk_shares", Type: proto.ColumnType_INT, Transform: transform.FromField("HkShares"), Description: ""},
		{Name: "eps", Type: proto.ColumnType_STRING, Transform: transform.FromField("Eps"), Description: ""},
		{Name: "eps_ttm", Type: proto.ColumnType_STRING, Transform: transform.FromField("EpsTtm"), Description: ""},
		{Name: "bps", Type: proto.ColumnType_STRING, Transform: transform.FromField("Bps"), Description: ""},
		{Name: "dividend_yield", Type: proto.ColumnType_STRING, Transform: transform.FromField("DividendYield"), Description: ""},
		{Name: "stock_derivatives", Type: proto.ColumnType_JSON, Transform: transform.FromField("StockDerivatives"), Description: ""},
	}
	return cols
}
