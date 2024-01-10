# Table: longport_history_execution - Query Historical Executions using SQL

The History Execution table is to query all of your historical orders that have been executed.

https://open.longportapp.com/en/docs/trade/execution/history_executions

## Examples

### Query History Executions by symbol

- `symbol` is optional, if present, only return the executions of the specified symbol, otherwise return all executions.
- `start_at` is optional, to limit the start time of the query, the format is UNIX timestamp in seconds, e.g.: `1650410999`.
- `end_at` is optional, to limit the end time of the query, the format is UNIX timestamp in seconds, e.g.: `1650410999`.

```sql+postgres
select
   *
from
   longport_history_execution;
```

```sql+sqlite
select
   *
from
   longport_history_execution;
```
