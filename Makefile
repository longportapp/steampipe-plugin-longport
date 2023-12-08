install:
	go build -o ~/.steampipe/plugins/ghcr.io/longportapp/longport@latest/steampipe-plugin-longport.plugin *.go
run: install
	STEAMPIPE_LOG_LEVEL=INFO steampipe query