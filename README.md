# [LongPort](https://open.longportapp.com) Plugin for Steampipe

Use SQL to query securities, quote from [LongPort](https://open.longportapp.com).

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install ghcr.io/longportapp/longport
```

Open `~/.steampipe/config/longport.spc` and setup your `AppKey`, `AppSecret` and `AccessToken`.

You can get them from [LongPort OpenAPI](https://open.longportapp.com/en/docs/how-to-access-api).

```conf
connection "longport" {
  plugin = "ghcr.io/longportapp/longport"

  app_key      = "YOUR_APP_KEY"
  app_secret   = "YOUR_ACCESS_SECRET"
  access_token = "YOUR_ACCESS_TOKEN"
}
```

Run a query:

```sql
select
  symbol, name_en, exchange, lot_size, total_shares, eps
from
  longport_static_info
where
  symbol in ('700.HK', 'AAPL.US', 'TSLA.US', 'NFLX.US');
```

Output:

```
+---------+---------------+----------+----------+--------------+---------------------+
| symbol  | name_en       | exchange | lot_size | total_shares | eps                 |
+---------+---------------+----------+----------+--------------+---------------------+
| AAPL.US | Apple Inc.    | NASD     | 1        | 15552752000  | 6.1607              |
| NFLX.US | Netflix, Inc. | NASD     | 1        | 437679669    | 10.1011             |
| TSLA.US | Tesla, Inc.   | NASD     | 1        | 3178921391   | 4.0201              |
| 700.HK  | TENCENT       | SEHK     | 100      | 9508314888   | 22.1632033382862025 |
+---------+---------------+----------+----------+--------------+---------------------+
```

## Tables

- [longport_brokers](./docs/tables/longport_brokers.md)
- [longport_candlesticks](./docs/tables/longport_candlesticks.md)
- [longport_depth](./docs/tables/longport_depth.md)
- [longport_history_executions](./docs/tables/longport_history_executions.md)
- [longport_intraday](./docs/tables/longport_intraday.md)
- [longport_option_quote](./docs/tables/longport_option_quote.md)
- [longport_participants](./docs/tables/longport_participants.md)
- [longport_quote](./docs/tables/longport_quote.md)
- [longport_static_info](./docs/tables/longport_static_info.md)
- [longport_trades](./docs/tables/longport_trades.md)
- [longport_warrant_quote](./docs/tables/longport_warrant_quote.md)

## License

MIT License
