package authorization

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func sign(nonce, timestamp, secret string) string {
	hash := hmac.New(sha256.New, []byte(secret))
	hash.Write([]byte(fmt.Sprintf("nonce=%s&timestamp=%s", nonce, timestamp)))
	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}
