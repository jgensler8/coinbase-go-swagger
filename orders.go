package coinbase

import (
	"context"
	"fmt"
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

func (c *Client) GetOrder(ctx context.Context, orderId string) (swagger.Order, *http.Response, error) {
	timestamp, header, err := generateSignHeader(c.secret, "GET", fmt.Sprintf("/orders/%s", orderId), nil)
	if err != nil {
		return swagger.Order{}, nil, err
	}
	return c.swaggerClient.OrdersApi.OrdersOrderIdGet(ctx, c.key, header, timestamp, c.passphrase, orderId)
}

func (c *Client) GetOrders(ctx context.Context) ([]swagger.Order, *http.Response, error) {
	timestamp, header, err := generateSignHeader(c.secret, "GET", "/orders", nil)
	if err != nil {
		return []swagger.Order{}, nil, err
	}
	return c.swaggerClient.OrdersApi.OrdersGet(ctx, c.key, header, timestamp, c.passphrase, nil)
}
