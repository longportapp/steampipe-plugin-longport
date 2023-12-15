# Table: `longport_quote`

Get Real-time Quotes Of Securities

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
   longport_quote
where
   symbol in
   (
      'BABA.US',
      'TSLA.US',
      '700.HK'
   )
;
```

Output:

```
+---------+-----------+------------+--------+--------+--------+-----------+----------------+------------+--------------+
| symbol  | last_done | prev_close | open   | high   | low    | volume    | turnover       | timestamp  | trade_status |
+---------+-----------+------------+--------+--------+--------+-----------+----------------+------------+--------------+
| 700.HK  | 306.2     | 307.8      | 307.4  | 309.2  | 301    | 8533999   | 2593423878.256 | 1702007999 | 0            |
| TSLA.US | 242.64    | 239.37     | 241.55 | 244.08 | 236.98 | 107142262 | 25810436588    | 1701982857 | 0            |
| BABA.US | 72.33     | 71.49      | 71.545 | 72.37  | 71.44  | 18966685  | 1367696030     | 1701982845 | 0            |
+---------+-----------+------------+--------+--------+--------+-----------+----------------+------------+--------------+
```

### Get Real-time Quotes Of Securities with Pre-market

```sql
select
  symbol, last_done, pre_market_quote->>'last_done' as pre_last_done
from
  longport_quote
where
  symbol in ('BABA.US', 'TSLA.US', '700.HK');
```

Output:

```
+---------+-----------+---------------+
| symbol  | last_done | pre_last_done |
+---------+-----------+---------------+
| TSLA.US | 242.64    | 241.65        |
| 700.HK  | 307.4     | <null>        |
| BABA.US | 72.33     | 71.52         |
+---------+-----------+---------------+
```
