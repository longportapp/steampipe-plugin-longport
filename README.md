# [LongPort](https://open.longportapp.com) Plugin for Steampipe

Use SQL to query securities, quote from [LongPort](https://open.longportapp.com).

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install ghcr.io/longportapp/longport
```

Run a query:

```sql
select
  *
from
  longport_static_info
where
  symbol in ("700.HK", "AAPL.US", "TSLA.US", "NFLX.US")
```

## Tables

- [longport_static_info](./docs/tables/longport_static_info.md)
- [longport_quote](./docs/tables/longport_quote.md)
- [longport_option_quote](./docs/tables/longport_option_quote.md)
- [longport_warrant_quote](./docs/tables/longport_warrant_quote.md)

## License

MIT License
