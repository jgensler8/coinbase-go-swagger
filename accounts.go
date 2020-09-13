package coinbase

import (
	"context"
	"net/http"

	"github.com/jgensler8/coinbase-go-swagger/swagger"
)

func (c *Client) AccountsGet(ctx context.Context) ([]swagger.Account, *http.Response, error) {
	timestamp, header, err := generateSignHeader(c.secret, "GET", "/accounts", nil)
	if err != nil {
		return nil, nil, err
	}
	return c.swaggerClient.AccountsApi.AccountsGet(ctx, c.key, header, timestamp, c.passphrase)
}
