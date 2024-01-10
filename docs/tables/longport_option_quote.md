# Table: longport_quote - Query Real-time Option Quotes Of Securities using SQL

The Real-time Option Quote table is used to obtain the real-time option quote data of security.

https://open.longportapp.com/en/docs/quote/pull/option-quote

## Examples

### Query Real-time Quotes by symbols

```sql+postgres
select
   *
from
   longport_option_quote
where
   symbol in
   (
      'BABA.US',
      'TSLA.US',
      '700.HK',
      "AAPL.US"
   );
```

```sql+sqlite
select
   *
from
   longport_option_quote
where
   symbol in
   (
      'BABA.US',
      'TSLA.US',
      '700.HK',
      "AAPL.US"
   );
```
