package coinbase

import (
	"github.com/jgensler8/coinbase-go-swagger/swagger"
)

// Client will wrap the Swagger library so headers get set easily
type Client struct {
	key        string
	secret     string
	passphrase string

	swaggerClient *swagger.APIClient
}

func NewClient(sandbox bool, key string, secret string, passphrase string) *Client {
	config := swagger.NewConfiguration()
	config.BasePath = ProURL(sandbox, URLTypeREST).String()

	return &Client{
		key:        key,
		secret:     secret,
		passphrase: passphrase,

		swaggerClient: swagger.NewAPIClient(config),
	}
}
