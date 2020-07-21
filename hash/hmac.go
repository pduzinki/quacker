package hash

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"hash"
)

// NewHmac creates and returns a new Hmac object
func NewHmac(key string) Hmac {
	h := hmac.New(sha256.New, []byte(key))
	return Hmac{
		hmac: h,
	}
}

// Hmac is a wrapper around the crypto/hmac package
type Hmac struct {
	hmac hash.Hash
}

// Hash will hash the provided string using HMAC
func (h Hmac) Hash(input string) string {
	h.hmac.Reset()
	h.hmac.Write([]byte(input))
	b := h.hmac.Sum(nil)
	return base64.URLEncoding.EncodeToString(b)
}
