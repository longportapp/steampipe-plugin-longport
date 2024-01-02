# Table: longport_history_order - Query Historical Orders using SQL

The History Order table is to query all of your historical orders, it includes the all orders that have been executed, canceled, or rejected.

https://open.longportapp.com/en/docs/trade/order/history_orders

## Examples

- `symbol` is optional, if present, only return the orders of the specified symbol, otherwise return all orders.
- `side` is optional, available values: `Buy`, `Sell`.
- `market` is optional, available values: `HK`, `US`.
- `start_at` is optional, to limit the start time of the query, the format is UNIX timestamp in seconds, e.g.: `1650410999`.
- `end_at` is optional, to limit the end time of the query, the format is UNIX timestamp in seconds, e.g.: `1650410999`.

### Query All History Orders

```sql+postgres
select
   symbol,
   order_id,
   status,
   stock_name,
   quantity,
   executed_quantity,
   price,
   executed_price,
   submitted_at,
   side,
   order_type,
   last_done,
   trigger_price,
   msg
from
   longport_history_order;
```

```sql+sqlite
select
   symbol,
   order_id,
   status,
   stock_name,
   quantity,
   executed_quantity,
   price,
   executed_price,
   submitted_at,
   side,
   order_type,
   last_done,
   trigger_price,
   msg
from
   longport_history_order;
```

### Query History Orders by symbol

If you present the `symbol`, only return the orders of the specified symbol.

```sql+postgres
select
   *
from
   longport_history_order
where
   symbol = 'TSLA.US';
```

```sql+sqlite
select
   *
from
   longport_history_order
where
   symbol = 'TSLA.US';
```
