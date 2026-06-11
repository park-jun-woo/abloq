//ff:func feature=queueio type=parser control=selection
//ff:what 큐 파일 라인 1개를 Item에 반영 — 들여쓰기로 keys 항목/payload 항목/최상위 필드를 분류, 모르는 라인은 에러 (Deserialize 전용)
package queueio

import (
	"fmt"
	"strconv"
	"strings"
)

// applyQueueLine classifies one serialized queue-file line and folds it into
// the item under construction. Indented `- "..."` lines are keys entries,
// other indented lines are payload entries, top-level lines are the fixed
// scalar fields; the key:/keys:/payload: markers carry no value themselves.
func applyQueueLine(it *Item, ln string) error {
	switch {
	case ln == "" || ln == "keys:" || ln == "payload:" || ln == "payload: {}":
		return nil
	case strings.HasPrefix(ln, "  - "):
		k, err := strconv.Unquote(strings.TrimPrefix(ln, "  - "))
		it.Keys = append(it.Keys, k)
		return err
	case strings.HasPrefix(ln, "  "):
		return applyPayloadLine(it, ln)
	case strings.HasPrefix(ln, "key: "):
		return nil // recomputed from lang/section/slug (JoinKey)
	case strings.HasPrefix(ln, "kind: "):
		return unquoteInto(&it.Kind, ln, "kind: ")
	case strings.HasPrefix(ln, "lang: "):
		return unquoteInto(&it.Lang, ln, "lang: ")
	case strings.HasPrefix(ln, "section: "):
		return unquoteInto(&it.Section, ln, "section: ")
	case strings.HasPrefix(ln, "slug: "):
		return unquoteInto(&it.Slug, ln, "slug: ")
	case strings.HasPrefix(ln, "priority: "):
		n, err := strconv.ParseInt(strings.TrimPrefix(ln, "priority: "), 10, 64)
		it.Priority = n
		return err
	}
	return fmt.Errorf("queue file: unrecognized line %q", ln)
}
