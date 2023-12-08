# Table: `longport_trades`

Get Security Trades

This API is used to obtain the trades data of security.

https://open.longportapp.com/en/docs/quote/pull/trade

## Examples

```sql
select * from longport_trades where symbol = 'TSLA.US' limit 3;
```

The `limit` default is `20`;

Output:

```
+--------+--------+------+
| symbol | trades | _ctx |
+--------+--------+------+
+--------+--------+------+
```
