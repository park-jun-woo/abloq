//ff:func feature=queueio type=parser control=sequence
//ff:what payload 블록 라인 1개(`  k: "v"`)를 Item.Payload에 반영 — 콜론 분리 + 값 unquote, 형식 불일치는 에러 (Deserialize 전용)
package queueio

import (
	"fmt"
	"strconv"
	"strings"
)

// applyPayloadLine parses one indented payload entry. The key is everything
// up to the first ": " (payload keys never contain ": " — they are scanner
// identifiers), the value is a strconv.Quote'd scalar.
func applyPayloadLine(it *Item, ln string) error {
	k, raw, ok := strings.Cut(strings.TrimPrefix(ln, "  "), ": ")
	if !ok {
		return fmt.Errorf("queue file: malformed payload line %q", ln)
	}
	v, err := strconv.Unquote(raw)
	if err != nil {
		return fmt.Errorf("queue file: payload value of %q: %w", k, err)
	}
	it.Payload[k] = v
	return nil
}
