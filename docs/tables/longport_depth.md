# Table: `longport_depth`

Security Depth, used to obtain the depth data of security.

https://open.longportapp.com/en/docs/quote/pull/depth

## Examples

### Get Real-time Quotes Of Securities

```sql
select
  *
from
  longport_depth
where
  symbol = '700.HK';
```

Output:

```
+---------+------------------------------------------------------------+----------------------------------------------------------+--------------------------------+
| symbol  | ask                                                        | bid                                                      | _ctx                           |
+---------+------------------------------------------------------------+----------------------------------------------------------+--------------------------------+
| TSLA.US | {"order_num":0,"position":1,"price":"242.75","volume":100} | {"order_num":0,"position":1,"price":"242.65","volume":1} | {"connection_name":"longport"} |
| TSLA.US | {"order_num":0,"position":1,"price":"242.75","volume":100} | {"order_num":0,"position":1,"price":"242.65","volume":1} | {"connection_name":"longport"} |
+---------+------------------------------------------------------------+----------------------------------------------------------+--------------------------------+
```
