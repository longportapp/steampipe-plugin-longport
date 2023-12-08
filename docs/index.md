---
organization: longportapp
category: ['SaaS']
icon_url: 'https://github.com/longportapp/steampipe-plugin-longport/assets/5518/1ca77b56-2cd1-4d85-ae02-9483da6ae9a1'
brand_color: '#00A3FF'
display_name: LongPort
name: longport
description: Steampipe plugin for Query from LongPort OpenAPI.
og_description: Query LongPort market data with SQL! Open source CLI. No DB required.
og_image: 'https://pub.lbkrs.com/files/202211/sJswdGqSX1xDqrES/lonport-seo-img.png'
---

# LongPort + Steampipe

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

[LongPort](https://open.longportapp.com) OpenAPI provides programmatic quote trading interfaces for investors with research and development capabilities and assists them to build trading or quote strategy analysis tools based on their own investment strategies.

For example:

```sql
select
  symbol, name_en, exchange, currency, lot_size, total_shares, eps
from
  longport_static_info
where
  symbol in ('BABA.US', 'TSLA.US', '700.HK');
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

## Get started

### Install

Download and install the latest Twitter plugin:

```bash
steampipe plugin install longport
```

### Credentials

Please visit https://open.longportapp.com/en/docs/how-to-access-api to get your API key.

### Configuration

Installing the latest longport plugin will create a config file (`~/.steampipe/config/longport.spc`) with a single connection named `longport`:

```hcl
connection "longport" {
  plugin = "longportapp/longport"

  app_key      = "YOUR_APP_KEY"
  app_secret   = "YOUR_ACCESS_SECRET"
  access_token = "YOUR_ACCESS_TOKEN"
}
```

## Get involved

https://github.com/longportapp/steampipe-plugin-longport
