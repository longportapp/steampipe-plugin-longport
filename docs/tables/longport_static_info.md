# Table: longport_static_info - Query Basic Information of Securities using SQL

The Static Information table is used to obtain the static information of security.

https://open.longportapp.com/en/docs/quote/pull/static

## Examples

### Query Basic Information by symbols

```sql+postgres
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
   );
```

```sql+sqlite
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
   );
```
