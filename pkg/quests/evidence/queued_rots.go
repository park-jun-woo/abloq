//ff:func feature=quest type=parser control=sequence topic=queue
//ff:what Seed 시점 고정된 payload rot_urls JSON → URL 목록 — rot-resolved의 잔존 검사 대상
package evidence

import (
	"encoding/json"
	"fmt"
)

// queuedRots decodes the frozen queue payload's rot_urls entry. An absent
// entry (claims-only item) yields nil — nothing to re-check.
func queuedRots(payload map[string]string) ([]string, error) {
	raw, ok := payload["rot_urls"]
	if !ok {
		return nil, nil
	}
	var urls []string
	if err := json.Unmarshal([]byte(raw), &urls); err != nil {
		return nil, fmt.Errorf("queue payload rot_urls: %w", err)
	}
	return urls, nil
}
