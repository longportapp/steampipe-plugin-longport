# Development Guides

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)

```sh
git clone https://github.com/longportapp/steampipe-plugin-longport.git
cd steampipe-plugin-longport
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make run
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
