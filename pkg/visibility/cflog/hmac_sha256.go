//ff:func feature=visibility type=client control=sequence topic=crawl
//ff:what HMAC-SHA256 1회 — 키와 메시지로 MAC 바이트 반환
package cflog

import (
	"crypto/hmac"
	"crypto/sha256"
)

// hmacSHA256 computes one HMAC-SHA256 round.
func hmacSHA256(key []byte, msg string) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(msg))
	return mac.Sum(nil)
}
