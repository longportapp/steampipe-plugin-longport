# Table: `longport_participants`

Get Broker IDs

This API is used to obtain participant IDs data (which can be synchronized once a day).

https://open.longportapp.com/en/docs/quote/pull/broker-ids

## Examples

```sql
select * from longport_participants limit 3;
```

Output:

```
+-----------------------+---------------------+---------------------------------------------+---------------------+--------------------------------+
| broker_ids            | participant_name_cn | participant_name_en                         | participant_name_hk | _ctx                           |
+-----------------------+---------------------+---------------------------------------------+---------------------+--------------------------------+
| [727]                 | 信期国际证券        | CF International Securities Company Limited | 信期國際證券        | {"connection_name":"longport"} |
| [4900,4908,4909]      | 日盛嘉富国际        | JS Cresvale Intl.                           | 日盛嘉富國際        | {"connection_name":"longport"} |
| [4194,4195,4203,4204] | IMC Asia Pacific    | IMC Asia Pacific                            | IMC Asia Pacific    | {"connection_name":"longport"} |
+-----------------------+---------------------+---------------------------------------------+---------------------+--------------------------------+
```
