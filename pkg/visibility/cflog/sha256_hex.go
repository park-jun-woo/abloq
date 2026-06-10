//ff:func feature=visibility type=client control=sequence topic=crawl
//ff:what SHA-256 16진 다이제스트 — 바이트 입력의 hex 문자열
package cflog

import (
	"crypto/sha256"
	"encoding/hex"
)

// sha256Hex returns the lowercase hex SHA-256 digest of b.
func sha256Hex(b []byte) string {
	sum := sha256.Sum256(b)
	return hex.EncodeToString(sum[:])
}
