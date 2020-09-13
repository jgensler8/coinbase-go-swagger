/*
 * Coinbase Pro API
 *
 * The Coinbase Pro API. See https://docs.pro.coinbase.com/ for more details on rate limiting, sandbox mode, and more.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type VolumeItem struct {
	ProductId string `json:"product_id,omitempty"`
	ExchangeVolume string `json:"exchange_volume,omitempty"`
	Volume string `json:"volume,omitempty"`
	RecordedAt string `json:"recorded_at,omitempty"`
}