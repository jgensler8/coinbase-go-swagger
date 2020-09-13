/*
 * Coinbase Pro API
 *
 * The Coinbase Pro API. See https://docs.pro.coinbase.com/ for more details on rate limiting, sandbox mode, and more.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type LimitOrderRequest struct {
	Size string `json:"size,omitempty"`
	Price string `json:"price,omitempty"`
	TimeInForce string `json:"time_in_force,omitempty"`
	CancelAfter string `json:"cancel_after,omitempty"`
	PostOnly bool `json:"post_only,omitempty"`
}
