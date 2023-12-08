# Table: `longport_history_executions`

Get History Executions

https://open.longportapp.com/en/docs/trade/execution/history_executions

## Examples

- `symbol` is optional.

```sql
select * from longport_history_executions;
```

Output:

```
+--------+----------+----------+---------------+----------+-------+------+
| symbol | order_id | trade_id | trade_done_at | quantity | price | _ctx |
+--------+----------+----------+---------------+----------+-------+------+
+--------+----------+----------+---------------+----------+-------+------+
```
