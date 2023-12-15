# Table: `longport_static_info`

Static information of securities.

This table can be used to query the static information of securities.

https://open.longportapp.com/en/docs/quote/pull/static

## Examples

### Get static information of securities

```sql
select
   symbol,
   name_en,
   exchange,
   currency,
   lot_size,
   total_shares,
   eps
from
   longport_static_info
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
+---------+---------------------------------------------+----------+----------+----------+--------------+---------------------+
| symbol  | name_en                                     | exchange | currency | lot_size | total_shares | eps                 |
+---------+---------------------------------------------+----------+----------+----------+--------------+---------------------+
| TSLA.US | Tesla, Inc.                                 | NASD     | USD      | 1        | 3178921391   | 4.0201              |
| BABA.US | Alibaba Group Holding Limited Sponsored ADR | NYSE     | USD      | 1        | 2543424136   | 4.0327              |
| 700.HK  | TENCENT                                     | SEHK     | HKD      | 100      | 9508314888   | 22.1632033382862025 |
+---------+---------------------------------------------+----------+----------+----------+--------------+---------------------+
```
