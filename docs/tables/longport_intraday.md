# Table: longport_intraday - Query Security Intraday using SQL

The Intraday table is used to obtain the intraday data of security.

https://open.longportapp.com/en/docs/quote/pull/intraday

## Examples

```sql+postgres
select
   *
from
   longport_intraday
where
   symbol = 'TSLA.US' limit 10;
```

```sql+sqlite
select
   *
from
   longport_intraday
where
   symbol = 'TSLA.US' limit 10;
```
