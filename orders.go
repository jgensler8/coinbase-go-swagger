package coinbase

import (
	"context"
	"net/http"

	"github.com/jgensler8/coinbase-go-swagger/swagger"
)

func (c *Client) CreateOrder(ctx context.Context, order swagger.OrderRequest) ([]swagger.Order, *http.Response, error) {
	timestamp, header, err := generateSignHeader(c.secret, "POST", "/orders", nil)
	if err != nil {
		return nil, nil, err
	}
	return c.swaggerClient.OrdersApi.OrdersPost(ctx, c.key, header, timestamp, c.passphrase, order)
}
