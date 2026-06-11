//ff:func feature=quest type=parser control=iteration dimension=1 topic=queue
//ff:what Seed 시점 고정된 payload claims JSON → 주장 해시 집합 — claim-scope·claims-resolved의 변경 허가 목록
package evidence

import (
	"encoding/json"
	"fmt"

	scanevidence "github.com/park-jun-woo/abloq/pkg/scan/evidence"
)

// queuedClaims decodes the frozen queue payload's claims entry (the Phase010
// scanner's ClaimRef JSON) into the hash set the claim rules consult. An
// absent entry (rot-only item) yields an empty set.
func queuedClaims(payload map[string]string) (map[string]bool, error) {
	raw, ok := payload["claims"]
	if !ok {
		return map[string]bool{}, nil
	}
	var refs []scanevidence.ClaimRef
	if err := json.Unmarshal([]byte(raw), &refs); err != nil {
		return nil, fmt.Errorf("queue payload claims: %w", err)
	}
	hashes := make(map[string]bool, len(refs))
	for _, r := range refs {
		hashes[r.Hash] = true
	}
	return hashes, nil
}
