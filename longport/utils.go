package longport

import (
	"context"
	"errors"
	"os"

	"github.com/longportapp/openapi-go/config"
	"github.com/longportapp/openapi-go/quote"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func connect(ctx context.Context, d *plugin.QueryData) (*quote.QuoteContext, error) {
	var _ = plugin.Logger(ctx)
	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "longport/quote"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*quote.QuoteContext), nil
	}

	appKey := os.Getenv("LONGPORT_APP_KEY")
	appSecret := os.Getenv("LONGPORT_APP_SECRET")
	accessToken := os.Getenv("LONGPORT_ACCESS_TOKEN")

	// First, use the token config
	longportConfig := GetConfig(d.Connection)
	if longportConfig.AppKey != nil {
		appKey = *longportConfig.AppKey
	}
	if longportConfig.AppSecret != nil {
		appSecret = *longportConfig.AppSecret
	}
	if longportConfig.AccessToken != nil {
		accessToken = *longportConfig.AccessToken
	}

	c, err := config.New()
	if err != nil {
		return nil, errors.New("config.New() error: " + err.Error())
	}

	c.AppKey = appKey
	c.AppSecret = appSecret
	c.AccessToken = accessToken

	var quoteContext *quote.QuoteContext

	// First, try to use the bearer token and OAuth 2.0
	if accessToken != "" {
		quoteContext, err = quote.NewFromCfg(c)
		if err != nil {
			return nil, errors.New("quote.NewFromCfg() error: " + err.Error())
		}

	} else {
		// Credentials not set
		return nil, errors.New("accessToken (or appKey, appSecret etc) must be configured")
	}

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, quoteContext)

	return quoteContext, nil
}

func symbolList(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) ([]string, error) {
	quals := d.EqualsQuals
	symbols := []string{quals["symbol"].GetStringValue()}
	return symbols, nil
}
