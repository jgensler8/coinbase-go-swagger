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
