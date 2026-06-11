//ff:func feature=gate type=generator control=sequence topic=evidence
//ff:what 공개 HashText API — 주장 텍스트의 안정 키(sha256 hex 앞 16자), Phase010 스캐너와 Phase018 claim 룰이 공유
//ff:why Phase018에서 pkg/scan/evidence의 hashText를 공용화 — 퀘스트 claim 룰이 같은 해시로 큐 payload를 대조해야 하고, gate→evidence 역방향 import는 순환이라 gate 쪽으로 옮긴다
package gate

import (
	"crypto/sha256"
	"encoding/hex"
)

// HashText derives the stable claim key from the claim text verbatim. 16 hex
// chars (64 bits) is collision-safe at corpus scale and short enough to read
// in a queue file. The evidence scanner (queue payload) and the quest claim
// rules (claim-scope, claims-resolved) share this key.
func HashText(s string) string {
	sum := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sum[:])[:16]
}
