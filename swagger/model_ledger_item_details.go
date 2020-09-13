/*
 * Coinbase Pro API
 *
 * The Coinbase Pro API. See https://docs.pro.coinbase.com/ for more details on rate limiting, sandbox mode, and more.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type LedgerItemDetails struct {
	// uuid of order
	OrderId string `json:"order_id,omitempty"`
	TradeId string `json:"trade_id,omitempty"`
	ProductId string `json:"product_id,omitempty"`
}
