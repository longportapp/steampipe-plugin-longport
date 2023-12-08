install:
	go build -o ~/.steampipe/plugins/hub.steampipe.io/plugins/turbot/longport@latest/steampipe-plugin-longport.plugin *.go
run: install
	STEAMPIPE_LOG_LEVEL=INFO steampipe query