---
organization: longportapp
category: ['SaaS']
engines: ['steampipe', 'sqlite', 'postgres', 'export']
icon_url: 'https://github.com/longportapp/steampipe-plugin-longport/assets/5518/1ca77b56-2cd1-4d85-ae02-9483da6ae9a1'
brand_color: '#00A3FF'
display_name: LongPort
name: longport
description: Steampipe plugin for Query from LongPort OpenAPI.
og_description: Query LongPort market data with SQL! Open source CLI. No DB required.
og_image: 'https://pub.lbkrs.com/files/202211/sJswdGqSX1xDqrES/lonport-seo-img.png'
---

# LongPort + Steampipe

[Steampipe](https://steampipe.io) is an open-source zero-ETL engine to instantly query cloud APIs using SQL.

[LongPort](https://open.longportapp.com) OpenAPI provides programmatic quote trading interfaces for investors with research and development capabilities and assists them to build trading or quote strategy analysis tools based on their own investment strategies.

For example:

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
+---------+---------------------------------------------+----------+----------+----------+--------------+---------------------+
| symbol  | name_en                                     | exchange | currency | lot_size | total_shares | eps                 |
+---------+---------------------------------------------+----------+----------+----------+--------------+---------------------+
| TSLA.US | Tesla, Inc.                                 | NASD     | USD      | 1        | 3178921391   | 4.0201              |
| BABA.US | Alibaba Group Holding Limited Sponsored ADR | NYSE     | USD      | 1        | 2543424136   | 4.0327              |
| 700.HK  | TENCENT                                     | SEHK     | HKD      | 100      | 9508314888   | 22.1632033382862025 |
+---------+---------------------------------------------+----------+----------+----------+--------------+---------------------+
```

## Quick Start

### Install

Download and install the latest LongPort plugin:

```bash
steampipe plugin install longportapp/longport
```

### Credentials

Please visit https://open.longportapp.com/en/docs/how-to-access-api to get your API key.

| Item        | Description                                                                                                                                                                                                                        |
| ----------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Credentials | LongPort requires [API key](https://open.longportapp.com/docs/getting-started) for all requests.                                                                                                                                   |
| Permissions | All API keys must visit [LongPort OpenAPI](https://open.longportapp.com/) account and Open API permission, please visit: [OpenAPI Intro](https://open.longportapp.com/docs/) to checkout.                                          |
| Radius      | Each connection represents a single LongPort Installation.                                                                                                                                                                         |
| Resolution  | 1. Credentials explicitly set in a steampipe config file (`~/.steampipe/config/longport.spc`)<br />2. Credentials specified in environment variables, e.g., `LONGPORT_APP_KEY`, `LONGPORT_APP_SECRET` and `LONGPORT_ACCESS_TOKEN`. |

### Configuration

Installing the latest longport plugin will create a config file (`~/.steampipe/config/longport.spc`) with a single connection named `longport`:

```hcl
connection "longport" {
  plugin = "longportapp/longport"

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
