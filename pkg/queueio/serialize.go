//ff:func feature=queueio type=generator control=iteration dimension=1
//ff:what 큐 파일 결정적 직렬화 — key 필드(게이트 계약) 선두 고정, payload 키 정렬, 타임스탬프·DB id 불포함
package queueio

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
)

// Serialize renders one queue file. The output is byte-deterministic for a
// given item (fixed field order, sorted payload keys, quoted scalars) — the
// CLI/endpoint diff equality and the idempotent no-op commit both rest on it.
// The `key:` line carries the verbatim <lang>/<section>/<slug> join key the
// honest-lastmod gate substring-matches.
func Serialize(it Item) []byte {
	var b bytes.Buffer
	fmt.Fprintf(&b, "key: %s\n", strconv.Quote(JoinKey(it.Lang, it.Section, it.Slug)))
	fmt.Fprintf(&b, "kind: %s\n", strconv.Quote(it.Kind))
	fmt.Fprintf(&b, "lang: %s\n", strconv.Quote(it.Lang))
	fmt.Fprintf(&b, "section: %s\n", strconv.Quote(it.Section))
	fmt.Fprintf(&b, "slug: %s\n", strconv.Quote(it.Slug))
	fmt.Fprintf(&b, "priority: %d\n", it.Priority)
	if len(it.Payload) == 0 {
		b.WriteString("payload: {}\n")
		return b.Bytes()
	}
	b.WriteString("payload:\n")
	keys := make([]string, 0, len(it.Payload))
	for k := range it.Payload {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Fprintf(&b, "  %s: %s\n", k, strconv.Quote(it.Payload[k]))
	}
	return b.Bytes()
}
