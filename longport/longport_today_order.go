package longport

import (
	"context"
	"strings"

	"github.com/longportapp/openapi-go"
	"github.com/longportapp/openapi-go/trade"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableLongPortTodayOrder(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "longport_today_order",
		Description: "History Orders.",
		List: &plugin.ListConfig{
			Hydrate:    listLongPortTodayOrder,
			KeyColumns: plugin.OptionalColumns([]string{"symbol", "side"}),
		},
		Columns: orderColumns(),
	}
}

func listLongPortTodayOrder(ctx context.Context, d *plugin.QueryData, p *plugin.HydrateData) (interface{}, error) {
	context, err := getTradeContext(ctx, d)

	if err != nil {
		return nil, err
	}

	symbol, err := symbolString(ctx, d, p)
	if err != nil {
		return nil, err
	}

	side, err := equalString(ctx, d, p, "side")
	if err != nil {
		return nil, err
	}

	market, err := equalString(ctx, d, p, "market")
	if err != nil {
		return nil, err
	}

	var params = trade.GetTodayOrders{}
	if symbol != "" {
		params.Symbol = symbol
	}
	if side != "" {
		params.Side = orderSide(side)
	}
	if market != "" {
		params.Market = openapi.Market(strings.ToUpper(market))
	}

	var items []*trade.Order
	items, err = context.TodayOrders(ctx, &params)
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		d.StreamListItem(ctx, item)
	}

	return nil, nil
}

func orderColumns() []*plugin.Column {
	return []*plugin.Column{
		{Name: "symbol", Type: proto.ColumnType_STRING, Transform: transform.FromField("Symbol"), Description: "Symbol of security."},
		{Name: "market", Type: proto.ColumnType_STRING, Transform: transform.FromField("Symbol").Transform((transformMarketFromSymbol)), Description: "Market of security."},
		{Name: "order_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("OrderId"), Description: "Order ID."},
		{Name: "status", Type: proto.ColumnType_STRING, Transform: transform.FromField("Status"), Description: "Order status."},
		{Name: "stock_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("StockName"), Description: "Stock Name."},
		{Name: "quantity", Type: proto.ColumnType_STRING, Transform: transform.FromField("Quantity"), Description: "Submitted Quantity."},
		{Name: "executed_quantity", Type: proto.ColumnType_STRING, Transform: transform.FromField("ExecutedQuantity"), Description: "Executed Quantity. When the order is not filled, value is 0."},
		{Name: "price", Type: proto.ColumnType_STRING, Transform: transform.FromField("Price"), Description: "Submitted Price."},
		{Name: "executed_price", Type: proto.ColumnType_STRING, Transform: transform.FromField("Price"), Description: "Executed Price."},
		{Name: "submitted_at", Type: proto.ColumnType_STRING, Transform: transform.FromField("SubmittedAt"), Description: "Submitted Time."},
		{Name: "side", Type: proto.ColumnType_STRING, Transform: transform.FromField("Side"), Description: "Order Side."},
		{Name: "order_type", Type: proto.ColumnType_STRING, Transform: transform.FromField("OrderType"), Description: "Order Type."},
		{Name: "last_done", Type: proto.ColumnType_STRING, Transform: transform.FromField("LastDone"), Description: "Last done."},
		{Name: "trigger_price", Type: proto.ColumnType_STRING, Transform: transform.FromField("TriggerPrice"), Description: "LIT / MIT Order Trigger Price."},
		{Name: "msg", Type: proto.ColumnType_STRING, Transform: transform.FromField("Msg"), Description: "Rejected message or remark, default value is empty string."},
		{Name: "tag", Type: proto.ColumnType_STRING, Transform: transform.FromField("Tag"), Description: "Order tag."},
		{Name: "time_in_force", Type: proto.ColumnType_STRING, Transform: transform.FromField("TimeInForce"), Description: "Time in force Type."},
		{Name: "expire_date", Type: proto.ColumnType_STRING, Transform: transform.FromField("ExpireDate"), Description: "Long term order expire date, format: YYYY-MM-DD, example: 2022-12-05."},
		{Name: "updated_at", Type: proto.ColumnType_STRING, Transform: transform.FromField("UpdatedAt"), Description: "Last updated time, formatted as a timestamp (second)."},
		{Name: "trigger_at", Type: proto.ColumnType_STRING, Transform: transform.FromField("TriggerAt"), Description: "Conditional order trigger time. formatted as a timestamp (second)."},
		{Name: "trailing_amount", Type: proto.ColumnType_STRING, Transform: transform.FromField("TrailingAmount"), Description: "TSMAMT / TSLPAMT order trailing amount."},
		{Name: "trailing_percent", Type: proto.ColumnType_STRING, Transform: transform.FromField("TrailingPercent"), Description: "TSMPCT / TSLPPCT order trailing percent."},
		{Name: "limit_offset", Type: proto.ColumnType_STRING, Transform: transform.FromField("LimitOffset"), Description: "TSLPAMT / TSLPPCT order limit offset amount."},
		{Name: "trigger_status", Type: proto.ColumnType_STRING, Transform: transform.FromField("TriggerStatus"), Description: "Conditional Order Trigger Status. When an order is not a conditional order or a conditional order is not triggered, the trigger status is NOT_USED."},
		{Name: "currency", Type: proto.ColumnType_STRING, Transform: transform.FromField("Currency"), Description: ""},
		{Name: "outside_rth", Type: proto.ColumnType_STRING, Transform: transform.FromField("OutsideRth"), Description: "Enable or disable outside regular trading hours. Default is UnknownOutsideRth when the order is not a US stock."},
		{Name: "remark", Type: proto.ColumnType_STRING, Transform: transform.FromField("Remark"), Description: ""},
	}
}

func transformMarketFromSymbol(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	item := map[string]interface{}{}

	if d.Value == nil {
		return item, nil
	}

	symbol, ok := d.Value.(string)
	if !ok {
		return item, nil
	}

	parts := strings.Split(symbol, ".")
	if len(parts) < 2 {
		return item, nil
	}

	return strings.ToUpper(parts[1]), nil
}
