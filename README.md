# [LongPort](https://open.longportapp.com) Plugin for Steampipe

Use SQL to query securities, quote from [LongPort](https://open.longportapp.com).

通过 [Steampipe](https://steampipe.io) 与此插件集成，可以采用 SQL 的方式查询 [LongPort](https://open.longportapp.com) 的证券行情和订单数据。

## Screenshots

![steampipe-plugin-longport](https://github.com/longportapp/steampipe-plugin-longport/assets/5518/0ca055d7-797e-493b-8b24-e4eb87710bc8)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install ghcr.io/longportapp/longport
```

Upgrade new version:

```shell
steampipe plugin uninstall ghcr.io/longportapp/longport
steampipe plugin install ghcr.io/longportapp/longport
```

Open `~/.steampipe/config/longport.spc` and setup your `AppKey`, `AppSecret` and `AccessToken`.

You can get them from [LongPort OpenAPI](https://open.longportapp.com/en/docs/how-to-access-api).

```conf
connection "longport" {
  plugin = "ghcr.io/longportapp/longport"

  # The longport app key. Required.
  # This can also be set via the `LONGPORT_APP_KEY` environment variable.
  # app_key      = "908ed33331781b219f8b850cc4669aa9"

  # The longport app secret. Required.
  # This can also be set via the `LONGPORT_APP_SECRET` environment variable.
  # app_secret   = "ddb3f3898f358257fa192cb0eb2acd615ac7c1377b7d9b4a4633fd4c6e4b155d"

  # The longport access token. Required.
  # This can also be set via the `LONGPORT_ACCESS_TOKEN` environment variable.
  # access_token = "tcwJAQJ7gk8RSijkrs1rN5Sn_L2Z8kAKZ_rIiI023-jJMdnIUO5T0RTl1HN7Q0tImFTHHWhz5KGMcUwHgpl7gwq44NvrR"
}
```

Alternatively, you can also use the standard Longport environment variables to obtain credentials only if other arguments (app_key, app_secret, and access_token) are not specified in the connection:

```bash
export LONGPORT_APP_KEY=908ed33331781b219f8b850cc4669aa9
export LONGPORT_APP_SECRET=ddb3f3898f358257fa192cb0eb2acd615ac7c1377b7d9b4a4633fd4c6e4b155d
export LONGPORT_ACCESS_TOKEN=tcwJAQJ7gk8RSijkrs1rN5Sn_L2Z8kAKZ_rIiI023-jJMdnIUO5T0RTl1HN7Q0tImFTHHWhz5KGMcUwHgpl7gwq44NvrR
```

Run a query, use `steampipe query` command to startup Steampipe Query interface.

```bash
$ steampipe query
```

Now you can write SQL now:

```sql
select
   symbol,
   name_en,
   exchange,
   currency,
   lot_size,
   total_shares,
   eps
from
   longport_static_info
where
   symbol in
   (
      'BABA.US',
      'TSLA.US',
      '700.HK'
   )
;
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

### Quote

- [longport_broker](./docs/tables/longport_broker.md)
- [longport_candlestick](./docs/tables/longport_candlestick.md)
- [longport_depth](./docs/tables/longport_depth.md)
- [longport_history_execution](./docs/tables/longport_history_execution.md)
- [longport_intraday](./docs/tables/longport_intraday.md)
- [longport_option_quote](./docs/tables/longport_option_quote.md)
- [longport_participant](./docs/tables/longport_participant.md)
- [longport_quote](./docs/tables/longport_quote.md)
- [longport_static_info](./docs/tables/longport_static_info.md)
- [longport_trade](./docs/tables/longport_trade.md)
- [longport_warrant_quote](./docs/tables/longport_warrant_quote.md)

### Trade

- [longport_account_balance](./docs/tables/longport_account_balance.md)
- [longport_history_execution](./docs/tables/longport_history_execution.md)
- [longport_history_order](./docs/tables/longport_history_order.md)
- [longport_today_executions](./docs/tables/longport_today_executions.md)
- [longport_today_order](./docs/tables/longport_today_order.md)

## License

MIT License
