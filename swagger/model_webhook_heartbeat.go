/*
 * Coinbase Pro API
 *
 * The Coinbase Pro API. See https://docs.pro.coinbase.com/ for more details on rate limiting, sandbox mode, and more.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type WebhookHeartbeat struct {
	Type_ string `json:"type,omitempty"`
	Sequence int32 `json:"sequence,omitempty"`
	LastTradeId int32 `json:"last_trade_id,omitempty"`
	ProductId string `json:"product_id,omitempty"`
	Time string `json:"time,omitempty"`
}