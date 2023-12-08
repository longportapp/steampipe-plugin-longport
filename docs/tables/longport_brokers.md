# Table: `longport_brokers`

Security Brokers

This API is used to obtain the real-time broker queue data of security.

https://open.longportapp.com/en/docs/quote/pull/depth

## Examples

```sql
select * from longport_brokers where symbol = 'TSLA.US';
```

Output:

```
+---------+-------------+-------------+--------------------------------+
| symbol  | ask_brokers | bid_brokers | _ctx                           |
+---------+-------------+-------------+--------------------------------+
| TSLA.US | []          | []          | {"connection_name":"longport"} |
+---------+-------------+-------------+--------------------------------+
```
