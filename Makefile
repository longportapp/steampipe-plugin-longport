install:
	go build -o ~/.steampipe/plugins/hub.steampipe.io/plugins/longportapp/longport@latest/steampipe-plugin-longport.plugin *.go
install_ghcr:
	go build -o ~/.steampipe/plugins/longportapp/longport@latest/steampipe-plugin-longport.plugin *.go
run: install_ghcr
	steampipe service stop --force || true
	STEAMPIPE_LOG_LEVEL=INFO steampipe query

	