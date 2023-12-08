# Table: `longport_today_executions`

Get Today Executions

https://open.longportapp.com/en/docs/trade/execution/today_executions

## Examples

- `symbol` is optional.

```sql
select * from longport_today_executions;
```

Output:

```
+--------+----------+----------+---------------+----------+-------+------+
| symbol | order_id | trade_id | trade_done_at | quantity | price | _ctx |
+--------+----------+----------+---------------+----------+-------+------+
+--------+----------+----------+---------------+----------+-------+------+
```
