/*
 * Coinbase Pro API
 *
 * The Coinbase Pro API. See https://docs.pro.coinbase.com/ for more details on rate limiting, sandbox mode, and more.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type WebsocketActivate struct {
	Type_ string `json:"type,omitempty"`
	ProductId string `json:"product_id,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	UserId string `json:"user_id,omitempty"`
	ProfileId string `json:"profile_id,omitempty"`
	OrderId string `json:"order_id,omitempty"`
	StopType string `json:"stop_type,omitempty"`
	Side string `json:"side,omitempty"`
	StopPrice string `json:"stop_price,omitempty"`
	Size string `json:"size,omitempty"`
	Funds string `json:"funds,omitempty"`
	Private bool `json:"private,omitempty"`
}