# Table: longport_trade - Query Trades Detail of Security using SQL

The Trade table is used to obtain the real-time trade details of security.

https://open.longportapp.com/en/docs/quote/pull/trade

## Examples

### Query the latest 10 trades by symbol

```sql+postgres
select
   *
from
   longport_trades
where
   symbol = 'TSLA.US' limit 10;
```

```sql+sqlite
select
   *
from
   longport_trades
where
   symbol = 'TSLA.US' limit 10;
```
