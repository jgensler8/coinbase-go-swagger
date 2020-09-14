package coinbase

import (
	"context"
	"net/http"

	"github.com/jgensler8/coinbase-go-swagger/swagger"
)

func (c *Client) GetTrades(ctx context.Context, productId string) ([]swagger.Trade, *http.Response, error) {
	timestamp, header, err := generateSignHeader(c.secret, "GET", "/accounts", nil)
	if err != nil {
		return nil, nil, err
	}
	return c.swaggerClient.ProductsApi.ProductsProductIdTradesGet(ctx, c.key, header, timestamp, c.passphrase, productId)
}

func (c *Client) GetHistoricRates(ctx context.Context, productId string, parameters swagger.HistoricRateRequest) ([]swagger.Candle, *http.Response, error) {
	timestamp, header, err := generateSignHeader(c.secret, "GET", "/accounts", nil)
	if err != nil {
		return nil, nil, err
	}
	candlesRaw, res, err := c.swaggerClient.ProductsApi.ProductsProductIdCandlesGet(ctx, c.key, header, timestamp, c.passphrase, productId, parameters)
	candles := []swagger.Candle{}
	for _, arr := range candlesRaw {
		// Silently ignore missing data
		if len(arr) < 6 {
			continue
		}
		candles = append(candles, swagger.Candle{
			Time:   int32(arr[0]),
			Low:    arr[1],
			High:   arr[2],
			Open:   arr[3],
			Close:  arr[4],
			Volume: arr[5],
		})
	}
	return candles, res, err
}
