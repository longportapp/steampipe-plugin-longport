# Table: `longport_intraday`

Get Security Intraday

This API is used to obtain the intraday data of security.

https://open.longportapp.com/en/docs/quote/pull/intraday

## Examples

```sql
select * from longport_intraday where symbol = 'TSLA.US' limit 10;
```

Output:

```
+---------+---------+--------+---------------+------------+------------+--------------------------------+
| symbol  | price   | volume | turnover      | avg_price  | timestamp  | _ctx                           |
+---------+---------+--------+---------------+------------+------------+--------------------------------+
| TSLA.US | 240.789 | 437476 | 105490700.876 | 241.480176 | 1701959520 | {"connection_name":"longport"} |
| TSLA.US | 240.93  | 285870 | 68863272.889  | 241.085309 | 1701959880 | {"connection_name":"longport"} |
| TSLA.US | 240.82  | 467465 | 112384055.727 | 241.117487 | 1701959640 | {"connection_name":"longport"} |
| TSLA.US | 240.25  | 427995 | 102910214.029 | 241.272485 | 1701959580 | {"connection_name":"longport"} |
| TSLA.US | 239.63  | 551273 | 132407446.588 | 240.977713 | 1701959940 | {"connection_name":"longport"} |
| TSLA.US | 242.45  | 622313 | 150466647.891 | 241.786124 | 1701959400 | {"connection_name":"longport"} |
| TSLA.US | 241.058 | 429678 | 103657018.038 | 241.115281 | 1701959760 | {"connection_name":"longport"} |
| TSLA.US | 241.237 | 427167 | 102935679.234 | 241.09706  | 1701959700 | {"connection_name":"longport"} |
| TSLA.US | 240.51  | 323273 | 77888662.655  | 241.100081 | 1701959820 | {"connection_name":"longport"} |
| TSLA.US | 241.07  | 640829 | 154708184.62  | 241.599783 | 1701959460 | {"connection_name":"longport"} |
+---------+---------+--------+---------------+------------+------------+--------------------------------+
```
