//ff:func feature=blogyaml type=rule control=iteration dimension=1
//ff:what [taxonomy-unique] geo.taxonomy 태그 중복 금지 검증 — 키 부재(빈 목록)는 적법(클러스터 스캐너가 검사를 스킵)
package blogyaml

import "fmt"

// ruleTaxonomyUnique rejects duplicate tags in the optional geo.taxonomy
// vocabulary. An absent or empty taxonomy is legal — the cluster scanner
// skips its tag-taxonomy check in that case.
func ruleTaxonomyUnique(filename string, b *Blog, idx lineIndex) []Diagnostic {
	seen := map[string]bool{}
	var diags []Diagnostic
	for _, tag := range b.Geo.Taxonomy {
		if seen[tag] {
			diags = append(diags, Diagnostic{
				File: filename, Line: lineOf(idx, "geo.taxonomy"), Rule: "taxonomy-unique",
				Message: fmt.Sprintf("geo.taxonomy has duplicate tag %q", tag),
			})
			continue
		}
		seen[tag] = true
	}
	return diags
}
