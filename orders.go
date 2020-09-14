package coinbase

import (
	"context"
	"net/http"

	"github.com/jgensler8/coinbase-go-swagger/swagger"
)

func (c *Client) CreateOrder(ctx context.Context, order swagger.OrderRequest) (swagger.Order, *http.Response, error) {
	timestamp, header, err := generateJsonSignHeader(c.secret, "POST", "/orders", order)
	if err != nil {
		return swagger.Order{}, nil, err
	}
	return c.swaggerClient.OrdersApi.OrdersPost(ctx, c.key, header, timestamp, c.passphrase, order)
}
