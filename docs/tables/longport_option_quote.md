# Table: `longport_quote`

Real-time Quotes Of Securities

This API is used to obtain the real-time quotes of securities, and supports all types of securities.

https://open.longportapp.com/en/docs/quote/pull/quote

## Examples

### Get Real-time Quotes Of Securities

```sql
select
   symbol,
   last_done,
   prev_close,
   open,
   high,
   low,
   volume,
   turnover,
   timestamp,
   trade_status
from
   longport_option_quote
where
   symbol in
   (
      'BABA.US',
      'TSLA.US',
      '700.HK',
      "AAPL.US"
   )
;
```

Output:

```
+---------+-----------+------------+--------+--------+--------+-----------+----------------+------------+--------------+
| symbol  | last_done | prev_close | open   | high   | low    | volume    | turnover       | timestamp  | trade_status |
+---------+-----------+------------+--------+--------+--------+-----------+----------------+------------+--------------+
| 700.HK  | 304.8     | 307.8      | 307.4  | 309.2  | 301    | 11198516  | 3409312698.576 | 1702016452 | 0            |
| AAPL.US | 194.27    | 192.32     | 193.63 | 195    | 193.59 | 47477655  | 9229516619     | 1701982859 | 0            |
| BABA.US | 72.33     | 71.49      | 71.545 | 72.37  | 71.44  | 18966685  | 1367696030     | 1701982845 | 0            |
| TSLA.US | 242.64    | 239.37     | 241.55 | 244.08 | 236.98 | 107142262 | 25810436588    | 1701982857 | 0            |
+---------+-----------+------------+--------+--------+--------+-----------+----------------+------------+--------------+
```
