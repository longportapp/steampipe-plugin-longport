# Table: longport_candlestick - Query Security Candlesticks using SQL

The Security Candlestick table is used to obtain the candlestick data of security.

https://open.longportapp.com/en/docs/quote/pull/candlestick

## Examples

### Query candlestick data by symbol

- `symbol` - symbol of security
- `period` - period of candlesticks, available values: `1`, `5`, `15`, `30`, `60`, `1000`, `2000`, `3000`, `4000`, see also [Period](https://open.longportapp.com/en/docs/quote/objects#period---candlestick-period)
- `adjust_type` - adjust type, available values: `0` - NO_ADJUST, `1` - FORWARD_ADJUST, see also [Adjust](https://open.longportapp.com/en/docs/quote/objects#adjusttype---candlestick-adjustment-type)

```sql+postgres
select
   *
from
   longport_candlestick
where
   symbol = 'TSLA.US'
   and period = 1
   and adjust_type = 0;
```

```sql+sqlite
select
   *
from
   longport_candlestick
where
   symbol = 'TSLA.US'
   and period = 1
   and adjust_type = 0;
```
