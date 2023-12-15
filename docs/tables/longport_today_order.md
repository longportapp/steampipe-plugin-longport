# Table: `longport_today_order`

Get Today Orders

https://open.longportapp.com/en/docs/trade/order/today_orders

## Examples

```sql
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
   longport_today_order;
```

Output:

```
+--------+--------------------+----------------+------------+----------+-------------------+-------+----------------+--------------+------+------------+-----------+---------------+-----+
| symbol | order_id           | status         | stock_name | quantity | executed_quantity | price | executed_price | submitted_at | side | order_type | last_done | trigger_price | msg |
+--------+--------------------+----------------+------------+----------+-------------------+-------+----------------+--------------+------+------------+-----------+---------------+-----+
| AMC.US | 917716883833516032 | CanceledStatus | AMC院线    | 1        | 0                 | 6.1   | 6.1            |              | Buy  | LO         | <null>    | 0             |     |
+--------+--------------------+----------------+------------+----------+-------------------+-------+----------------+--------------+------+------------+-----------+---------------+-----+
```
