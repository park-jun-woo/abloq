//ff:func feature=scan type=generator control=sequence topic=evidence
//ff:what 주장 텍스트의 안정 키 — sha256 hex 앞 16자, 같은 문장이면 라인이 흔들려도 같은 키
package evidence

import (
	"crypto/sha256"
	"encoding/hex"
)

// hashText derives the stable claim key from the claim text verbatim. 16 hex
// chars (64 bits) is collision-safe at corpus scale and short enough to read
// in a queue file.
func hashText(s string) string {
	sum := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sum[:])[:16]
}
