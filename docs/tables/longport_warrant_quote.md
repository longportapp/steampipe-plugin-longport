# Table: longport_warrant_quote - Query Real-time Quotes Of Warrant Securities using SQL

The Real-time Quotes Of Warrant table is used to obtain the real-time quotes of HK warrants, including the warrant-specific data.

https://open.longportapp.com/en/docs/quote/pull/warrant-quote

## Examples

### Query Real-time Quotes Of Warrant Securities

```sql+postgres
select
   *
from
   longport_warrant_quote
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
   longport_warrant_quote
where
   symbol in
   (
      'BABA.US',
      'TSLA.US',
      '700.HK',
      "AAPL.US"
   );
```
