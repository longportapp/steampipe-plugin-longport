# Table: longport_participant - Query Security Participants using SQL

The Security Participant table is used to obtain the participant data of security, which can be synchronized once a day.

https://open.longportapp.com/en/docs/quote/pull/broker-ids

## Examples

### Query Top 5 Participants

```sql+postgres
select
   *
from
   longport_participant limit 5;
```

```sql+sqlite
select
   *
from
   longport_participant limit 5;
```
