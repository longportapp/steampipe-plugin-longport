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
		Description: "Basic Information Of Securities.",
		List: &plugin.ListConfig{
			Hydrate:    listStaticInfo,
			KeyColumns: plugin.SingleColumn("symbol"),
		},
		Columns: securityInfoColumns(),
	}
}

func listStaticInfo(ctx context.Context, d *plugin.QueryData, p *plugin.HydrateData) (interface{}, error) {
	quoteContext, err := connect(ctx, d)

	if err != nil {
		return nil, err
	}

	symbols, err := symbolList(ctx, d, p)
	if err != nil {
		return nil, err
	}

	infos, err := quoteContext.StaticInfo(ctx, symbols)
	if err != nil {
		return nil, err
	}

	for _, info := range infos {
		d.StreamListItem(ctx, info)
	}

	return nil, nil
}

// https://github.com/longportapp/openapi-go/blob/main/quote/types.go#L187
func securityInfoColumns() []*plugin.Column {
	cols := []*plugin.Column{
		// Top columns
		{Name: "symbol", Type: proto.ColumnType_STRING, Transform: transform.FromField("Symbol"), Description: "Security code"},
		{Name: "name_cn", Type: proto.ColumnType_STRING, Transform: transform.FromField("NameCn"), Description: "Security name (zh-CN)"},
		{Name: "name_en", Type: proto.ColumnType_STRING, Transform: transform.FromField("NameEn"), Description: "Security name (en)"},
		{Name: "name_hk", Type: proto.ColumnType_STRING, Transform: transform.FromField("NameHk"), Description: "Security name (zh-HK)"},
		{Name: "exchange", Type: proto.ColumnType_STRING, Transform: transform.FromField("Exchange"), Description: "Exchange which the security belongs to"},
		{Name: "currency", Type: proto.ColumnType_STRING, Transform: transform.FromField("Currency"), Description: "Trading currency"},
		{Name: "lot_size", Type: proto.ColumnType_INT, Transform: transform.FromField("LotSize"), Description: "Lot size"},
		{Name: "total_shares", Type: proto.ColumnType_INT, Transform: transform.FromField("TotalShares"), Description: "Total shares"},
		{Name: "circulating_shares", Type: proto.ColumnType_INT, Transform: transform.FromField("CirculatingShares"), Description: "Circulating shares"},
		{Name: "hk_shares", Type: proto.ColumnType_INT, Transform: transform.FromField("HkShares"), Description: "HK shares (only HK stocks)"},
		{Name: "eps", Type: proto.ColumnType_STRING, Transform: transform.FromField("Eps"), Description: "Earnings per share"},
		{Name: "eps_ttm", Type: proto.ColumnType_STRING, Transform: transform.FromField("EpsTtm"), Description: "Earnings per share (TTM)"},
		{Name: "bps", Type: proto.ColumnType_STRING, Transform: transform.FromField("Bps"), Description: "Net assets per share"},
		{Name: "dividend_yield", Type: proto.ColumnType_STRING, Transform: transform.FromField("DividendYield"), Description: "Dividend yield"},
		{Name: "stock_derivatives", Type: proto.ColumnType_JSON, Transform: transform.FromField("StockDerivatives"), Description: "Types of supported derivatives"},
		{Name: "board", Type: proto.ColumnType_STRING, Transform: transform.FromField("Board"), Description: "The board to whitch the security belongs, see Board for details"},
	}
	return cols
}
