package longport

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableLongPortAccountBalance(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "longport_account_balance",
		Description: "Get Account Balance.",
		List: &plugin.ListConfig{
			Hydrate: listLongPortAccountBalance,
		},
		Columns: []*plugin.Column{
			{Name: "currency", Type: proto.ColumnType_STRING, Transform: transform.FromField("Currency"), Description: "Currency"},
			{Name: "total_cash", Type: proto.ColumnType_STRING, Transform: transform.FromField("TotalCash"), Description: "Total Cash"},
			{Name: "max_finance_amount", Type: proto.ColumnType_STRING, Transform: transform.FromField("MaxFinanceAmount"), Description: "Max Finance Amount"},
			{Name: "remaining_finance_amount", Type: proto.ColumnType_STRING, Transform: transform.FromField("RemainingFinanceAmount"), Description: "Remaining Finance Amount"},
			{Name: "risk_level", Type: proto.ColumnType_INT, Transform: transform.FromField("RiskLevel"), Description: "Risk Level (0 - safe, 1 - medium, 2 - early warning, 3 - danger)"},
			{Name: "margin_call", Type: proto.ColumnType_STRING, Transform: transform.FromField("MarginCall"), Description: "Margin Call"},
			{Name: "net_assets", Type: proto.ColumnType_STRING, Transform: transform.FromField("NetAssets"), Description: "Net Assets"},
			{Name: "init_margin", Type: proto.ColumnType_STRING, Transform: transform.FromField("InitMargin"), Description: "Initial Margin"},
			{Name: "maintenance_margin", Type: proto.ColumnType_STRING, Transform: transform.FromField("MaintenanceMargin"), Description: "Maintenance Margin"},
			{Name: "cash_infos", Type: proto.ColumnType_JSON, Transform: transform.FromField("CashInfos"), Description: "Cash Details"},
			{Name: "timestamp", Type: proto.ColumnType_INT, Transform: transform.FromField("Trade.Timestamp"), Description: "Time of trading."},
		},
	}
}

func listLongPortAccountBalance(ctx context.Context, d *plugin.QueryData, p *plugin.HydrateData) (interface{}, error) {
	tradeContext, err := getTradeContext(ctx, d)

	if err != nil {
		return nil, err
	}

	accountBalances, err := tradeContext.AccountBalance(ctx)
	if err != nil {
		return nil, err
	}

	for _, item := range accountBalances {
		d.StreamListItem(ctx, item)
	}

	return nil, nil
}
