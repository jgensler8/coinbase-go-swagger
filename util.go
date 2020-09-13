package coinbase

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"strconv"
	"time"
)

const (
	ProductIDETHUSD = "ETH-USD"
)

// generateSignHeader creates the value for the CB-ACCESS-SIGN header
// secret: base64 encoded string created during API key generation
// method: GET
// requestPath: /accounts
// body: {...} or nil
//
// returns
// timestamp, sign header, err
func generateSignHeader(secret string, method string, requestPath string, body []byte) (string, string, error) {
	timestamp := strconv.FormatInt(time.Now().UTC().Unix(), 10)

	// create the prehash string by concatenating required parts
	what := timestamp + method + requestPath
	if len(body) > 0 {
		what += string(body)
	}

	// decode the base64 secret
	hmacKey, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		return "", "", err
	}

	// create a sha256 hmac with the secret
	h := hmac.New(sha256.New, hmacKey)
	h.Write([]byte(what))

	// sign the require message with the hmac
	// and finally base64 encode the result
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	return timestamp, signature, nil
}
