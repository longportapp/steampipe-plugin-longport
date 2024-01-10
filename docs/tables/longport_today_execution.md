# Table: longport_today_execution - Query Today Executions using SQL

The Today Execution table is to query all of your orders that have been executed today.

https://open.longportapp.com/en/docs/trade/execution/today_executions

## Examples

- `symbol` is optional, if present, only return the executions of the specified symbol (e.g.: `TSLA.US`), otherwise return all executions.

### Query all of my orders that have been executed today

```sql+postgres
select
   *
from
   longport_today_execution;
```

```sql+sqlite
select
   *
from
   longport_today_execution;
```
