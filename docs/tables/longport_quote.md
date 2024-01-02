# Table: longport_quote - Query Real-time Quotes Of Securities using SQL

The Real-time Quote table is used to obtain the real-time quote data of security.

https://open.longportapp.com/en/docs/quote/pull/quote

## Examples

### Query Real-time Quotes by symbols

This example shows to query real-time quotes of `BABA.US`, `TSLA.US` and `700.HK`.

```sql+postgres
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
   );
```

```sql+sqlite
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
   );
```

### Query Pre-market Quotes by symbols

```sql+postgres
select
   symbol,
   last_done,
   pre_market_quote ->> 'last_done' as pre_last_done
from
   longport_quote
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
   *
from
   longport_quote
where
   symbol in
   (
      'BABA.US',
      'TSLA.US',
      '700.HK'
   );
```
