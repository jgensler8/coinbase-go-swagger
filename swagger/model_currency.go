/*
 * Coinbase Pro API
 *
 * The Coinbase Pro API. See https://docs.pro.coinbase.com/ for more details on rate limiting, sandbox mode, and more.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Currency struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	MinSize string `json:"min_size,omitempty"`
	Status string `json:"status,omitempty"`
	StatusMessage string `json:"status_message,omitempty"`
	MaxPrecision string `json:"max_precision,omitempty"`
	ConveritbleTo []string `json:"converitble_to,omitempty"`
	Details *interface{} `json:"details,omitempty"`
}
