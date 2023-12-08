![image](https://hub.steampipe.io/images/plugins/turbot/longport-social-graphic.png)

# [LongPort](https://open.longportapp.com) Plugin for Steampipe

Use SQL to query securities, quote from [LongPort](https://open.longportapp.com).

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install longport
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

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)

```sh
git clone https://github.com/longportapp/steampipe-plugin-longport.git
cd steampipe-plugin-longport
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/longport.spc
```

Try it!

```
steampipe query
> .inspect longport
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-longport/blob/main/LICENSE).
