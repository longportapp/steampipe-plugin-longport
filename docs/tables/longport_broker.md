# Table: `longport_broker`

Security Brokers

This API is used to obtain the real-time broker queue data of security.

https://open.longportapp.com/en/docs/quote/pull/depth

## Examples

```sql
select
   *
from
   longport_broker
where
   symbol = 'TSLA.US';
```
