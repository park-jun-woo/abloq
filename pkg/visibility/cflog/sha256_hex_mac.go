//ff:func feature=visibility type=client control=sequence topic=crawl
//ff:what 키드 HMAC-SHA256의 16진 다이제스트 — 서명 문자열의 최종 서명값
package cflog

import "encoding/hex"

// sha256HexMAC returns the lowercase hex HMAC-SHA256 of msg under key.
func sha256HexMAC(key []byte, msg string) string {
	return hex.EncodeToString(hmacSHA256(key, msg))
}
