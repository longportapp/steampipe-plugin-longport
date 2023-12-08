# Table: `longport_trades`

Get Security Trades

This API is used to obtain the trades data of security.

https://open.longportapp.com/en/docs/quote/pull/trade

## Examples

```sql
select * from longport_trades where symbol = 'TSLA.US' limit 10;
```

The `limit` default is `20`;

Output:

```
+---------+---------+--------+------------+-----------+---------------+------------+--------------------------------+
| symbol  | price   | volume | trade_type | direction | trade_session | timestamp  | _ctx                           |
+---------+---------+--------+------------+-----------+---------------+------------+--------------------------------+
| TSLA.US | 242.590 | 5      | I          | 0         | 1             | 1702026168 | {"connection_name":"longport"} |
| TSLA.US | 242.640 | 1      | I          | 0         | 1             | 1702026176 | {"connection_name":"longport"} |
| TSLA.US | 242.590 | 5      | I          | 0         | 1             | 1702026168 | {"connection_name":"longport"} |
| TSLA.US | 242.590 | 2      | I          | 0         | 1             | 1702026168 | {"connection_name":"longport"} |
| TSLA.US | 242.590 | 10     | I          | 0         | 1             | 1702026168 | {"connection_name":"longport"} |
| TSLA.US | 242.590 | 9      | I          | 0         | 1             | 1702026176 | {"connection_name":"longport"} |
| TSLA.US | 242.630 | 100    |            | 0         | 1             | 1702026176 | {"connection_name":"longport"} |
| TSLA.US | 242.590 | 5      | I          | 0         | 1             | 1702026168 | {"connection_name":"longport"} |
| TSLA.US | 242.590 | 10     | I          | 0         | 1             | 1702026168 | {"connection_name":"longport"} |
| TSLA.US | 242.610 | 5      | I          | 0         | 1             | 1702026176 | {"connection_name":"longport"} |
+---------+---------+--------+------------+-----------+---------------+------------+--------------------------------+
```
