package coinbase

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/url"
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

type URLType int

const (
	URLTypeWebsite   URLType = iota
	URLTypeOAuth     URLType = iota
	URLTypeREST      URLType = iota
	URLTypeWebsocket URLType = iota
	URLTypeFIX       URLType = iota
)

func mustParseURL(u string) *url.URL {
	p, _ := url.Parse(u)
	return p
}

// ProURL returns the URL for a given service or nil if no URLType matched
// param sandbox: setting to `true` will use Coinbase's sandbox environment
// param URLType: should be one specified in this package (Website, REST, ...)
func ProURL(sandbox bool, urlType URLType) *url.URL {
	switch urlType {
	case URLTypeWebsite:
		if sandbox {
			return mustParseURL("https://public.sandbox.pro.coinbase.com")
		}
		// TODO(jeffg): figure out if this is the case
		return mustParseURL("https://pro.coinbase.com")
	case URLTypeOAuth:
		if sandbox {
			return nil
		}
		return mustParseURL("https://coinbase.com")
	case URLTypeREST:
		if sandbox {
			return mustParseURL("https://api-public.sandbox.pro.coinbase.com")
		}
		return mustParseURL("https://api.pro.coinbase.com")
	case URLTypeWebsocket:
		if sandbox {
			return mustParseURL("wss://ws-feed-public.sandbox.pro.coinbase.com")
		}
		return mustParseURL("wss://ws-feed.pro.coinbase.com")
	case URLTypeFIX:
		if sandbox {
			return mustParseURL("tcp+ssl://fix-public.sandbox.pro.coinbase.com:4198")
		}
		return mustParseURL("tcp+ssl://fix.pro.coinbase.com:4198")
	}
	return nil
}
