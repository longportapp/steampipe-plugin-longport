# Table: longport_today_order - Query Today Orders using SQL

The Today Order table is to query all of your orders that have been submitted today (including the orders that have been executed, canceled, or rejected).

https://open.longportapp.com/en/docs/trade/order/today_orders

## Examples

### Query all of my orders that have been submitted today

```sql+postgres
select
   *
from
   longport_today_order;
```

```sql+sqlite
select
   *
from
   longport_today_order;
```
