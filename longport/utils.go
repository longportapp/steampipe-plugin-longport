package longport

import (
	"context"
	"errors"
	"os"
	"strings"

	"github.com/longportapp/openapi-go/config"
	"github.com/longportapp/openapi-go/quote"
	"github.com/longportapp/openapi-go/trade"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func getQuoteContext(ctx context.Context, d *plugin.QueryData) (*quote.QuoteContext, error) {
	var _ = plugin.Logger(ctx)
	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "longport-QuoteContext"
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

	if len(appKey) == 0 || len(appSecret) == 0 || len(accessToken) == 0 {
		return nil, errors.New("app_key, app_secret and access_token must be configured")
	}

	c, err := config.New(config.WithConfigKey(appKey, appSecret, accessToken))
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

func getTradeContext(ctx context.Context, d *plugin.QueryData) (*trade.TradeContext, error) {
	var _ = plugin.Logger(ctx)
	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "longport-TradeContext"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*trade.TradeContext), nil
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

	if len(appKey) == 0 || len(appSecret) == 0 || len(accessToken) == 0 {
		return nil, errors.New("app_key, app_secret and access_token must be configured")
	}

	c, err := config.New(config.WithConfigKey(appKey, appSecret, accessToken))
	if err != nil {
		return nil, errors.New("config.New() error: " + err.Error())
	}

	c.AppKey = appKey
	c.AppSecret = appSecret
	c.AccessToken = accessToken

	var tradeContext *trade.TradeContext

	// First, try to use the bearer token and OAuth 2.0
	tradeContext, err = trade.NewFromCfg(c)
	if err != nil {
		return nil, errors.New("trade.NewFromCfg() error: " + err.Error())
	}

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, tradeContext)

	return tradeContext, nil
}

func symbolList(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) ([]string, error) {
	return equalList(ctx, d, nil, "symbol")
}

func symbolString(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (string, error) {
	return equalString(ctx, d, nil, "symbol")
}

func equalList(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData, field string) ([]string, error) {
	quals := d.EqualsQuals
	vals := []string{quals[field].GetStringValue()}
	return vals, nil
}

func equalString(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData, field string) (string, error) {
	quals := d.EqualsQuals
	return quals[field].GetStringValue(), nil
}

func equalInt(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData, field string) (int, error) {
	quals := d.EqualsQuals
	return int(quals[field].GetInt64Value()), nil
}

// Get limit, default is 20
func queryLimit(d *plugin.QueryData) int32 {
	limit := int32(20)
	if d.QueryContext.Limit != nil {
		limit = int32(*d.QueryContext.Limit)
	}

	return limit
}

func orderSide(side string) trade.OrderSide {
	side = strings.ToLower(side)
	switch side {
	case "buy":
		return trade.OrderSideBuy
	default:
		return trade.OrderSideSell
	}
}
