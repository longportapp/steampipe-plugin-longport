# Table: longport_broker - Query Security Brokers using SQL

The Security Broker table is used to obtain the broker queue data of security.

https://open.longportapp.com/en/docs/quote/pull/brokers

## Table Usage Guide

The `longport_broker` table provides insights into the broker queue data of security. The `symbol` column is required to query the broker queue data of the security.

## Examples

### List all brokers by symbol

You must specify the `symbol` to query the broker queue data of the security.

```sql+postgres
select
   *
from
   longport_broker
where
   symbol = 'TSLA.US';
```

```sql+sqlite
select
   *
from
   longport_broker
where
   symbol = 'TSLA.US';
```
