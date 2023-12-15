# Table: `longport_candlestick`

Get Security Candlesticks

This API is used to obtain the candlestick data of security.

https://open.longportapp.com/en/docs/quote/pull/candlestick

## Examples

### conditions

- `symbol` - symbol of security
- `period` - period of candlesticks, available values: `1`, `5`, `15`, `30`, `60`, `1000`, `2000`, `3000`, `4000`, see also [Period](https://open.longportapp.com/en/docs/quote/objects#period---candlestick-period)
- `adjust_type` - adjust type, available values: `0` - NO_ADJUST, `1` - FORWARD_ADJUST, see also [Adjust](https://open.longportapp.com/en/docs/quote/objects#adjusttype---candlestick-adjustment-type)

```sql
select
   *
from
   longport_candlestick
where
   symbol = 'TSLA.US'
   and period = 1
   and adjust_type = 0;
```

Output:

```
+---------+---------+---------+---------+---------+--------+--------------+------------+--------+-------------+--------------------------------+
| symbol  | close   | open    | low     | high    | volume | turnover     | timestamp  | period | adjust_type | _ctx                           |
+---------+---------+---------+---------+---------+--------+--------------+------------+--------+-------------+--------------------------------+
| TSLA.US | 241.949 | 242.05  | 241.94  | 242.08  | 143180 | 34652271.029 | 1701981660 | 1      | 0           | {"connection_name":"longport"} |
| TSLA.US | 242.17  | 241.931 | 241.931 | 242.229 | 231430 | 56035140.518 | 1701981780 | 1      | 0           | {"connection_name":"longport"} |
| TSLA.US | 242.64  | 242.57  | 242.57  | 242.9   | 77466  | 18796382.227 | 1701982800 | 1      | 0           | {"connection_name":"longport"} |
| TSLA.US | 242.669 | 242.54  | 242.54  | 242.67  | 95793  | 23237579.131 | 1701982440 | 1      | 0           | {"connection_name":"longport"} |
| TSLA.US | 241.94  | 242.16  | 241.91  | 242.21  | 123891 | 29981210.387 | 1701981840 | 1      | 0           | {"connection_name":"longport"} |
| TSLA.US | 242.574 | 242.56  | 242.52  | 242.59  | 53368  | 12944596.656 | 1701982500 | 1      | 0           | {"connection_name":"longport"} |
| TSLA.US | 241.934 | 241.94  | 241.86  | 241.98  | 78982  | 19107468.436 | 1701981720 | 1      | 0           | {"connection_name":"longport"} |
| TSLA.US | 242.21  | 241.97  | 241.96  | 242.26  | 223669 | 54163222.366 | 1701982200 | 1      | 0           | {"connection_name":"longport"} |
| TSLA.US | 242.47  | 242.58  | 242.44  | 242.64  | 187960 | 45585502.147 | 1701982560 | 1      | 0           | {"connection_name":"longport"} |
| TSLA.US | 241.958 | 241.93  | 241.69  | 241.995 | 174768 | 42266348.185 | 1701981900 | 1      | 0           | {"connection_name":"longport"} |
| TSLA.US | 242.57  | 242.49  | 242.45  | 242.63  | 181228 | 43957964.694 | 1701982680 | 1      | 0           | {"connection_name":"longport"} |
| TSLA.US | 242.56  | 242.19  | 242.11  | 242.59  | 236332 | 57289288.593 | 1701982260 | 1      | 0           | {"connection_name":"longport"} |
| TSLA.US | 242.68  | 242.601 | 242.59  | 242.68  | 68545  | 16630164.525 | 1701982740 | 1      | 0           | {"connection_name":"longport"} |
| TSLA.US | 242.472 | 242.46  | 242.42  | 242.49  | 123370 | 29915687.906 | 1701982620 | 1      | 0           | {"connection_name":"longport"} |
| TSLA.US | 242.621 | 242.54  | 242.32  | 242.64  | 230291 | 55841781.298 | 1701982320 | 1      | 0           | {"connection_name":"longport"} |
| TSLA.US | 242.55  | 242.636 | 242.54  | 242.65  | 140142 | 33994913.929 | 1701982380 | 1      | 0           | {"connection_name":"longport"} |
| TSLA.US | 241.901 | 241.95  | 241.83  | 242     | 80318  | 19429353.755 | 1701982020 | 1      | 0           | {"connection_name":"longport"} |
| TSLA.US | 241.855 | 241.9   | 241.84  | 241.95  | 72884  | 17629046.889 | 1701982080 | 1      | 0           | {"connection_name":"longport"} |
| TSLA.US | 241.98  | 241.93  | 241.83  | 241.99  | 105473 | 25514924.039 | 1701981960 | 1      | 0           | {"connection_name":"longport"} |
| TSLA.US | 241.97  | 241.869 | 241.84  | 241.97  | 74305  | 17974739.526 | 1701982140 | 1      | 0           | {"connection_name":"longport"} |
+---------+---------+---------+---------+---------+--------+--------------+------------+--------+-------------+--------------------------------+
```
