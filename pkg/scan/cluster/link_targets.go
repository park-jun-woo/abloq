//ff:func feature=scan type=parser control=iteration dimension=1 topic=cluster
//ff:what 본문 마크다운 링크 대상 중 기본 언어 글로 해석되는 키 목록 수집 — 문서 등장 순서, 중복 제거
package cluster

import (
	"regexp"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

var reLinkTarget = regexp.MustCompile(`\]\(([^)\s]+)`)

// linkTargets extracts the markdown link targets of one body and keeps the
// ones that resolve to a default-language article URL, as deduplicated
// <section>/<slug> keys in document order. Corpus existence and self-link
// filtering happen later, once every node is known.
func linkTargets(b *blogyaml.Blog, lang, body string) []string {
	seen := map[string]bool{}
	keys := make([]string, 0)
	for _, m := range reLinkTarget.FindAllStringSubmatch(body, -1) {
		key, ok := resolveTarget(b, lang, m[1])
		if !ok || seen[key] {
			continue
		}
		seen[key] = true
		keys = append(keys, key)
	}
	return keys
}
