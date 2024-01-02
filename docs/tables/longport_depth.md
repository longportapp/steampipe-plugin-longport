# Table: longport_depth - Query Security Depth using SQL

The Security Depth table is used to obtain the real-time depth data of security.

https://open.longportapp.com/en/docs/quote/pull/depth

## Examples

### Query Real-time Quotes by symbol

```sql+postgres
select
  *
from
  longport_depth
where
  symbol = '700.HK';
```

```sql+sqlite
select
  *
from
  longport_depth
where
  symbol = '700.HK';
```
