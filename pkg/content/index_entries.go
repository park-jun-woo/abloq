//ff:func feature=content type=parser control=iteration dimension=1
//ff:what 검증 완료된 Blog로 저장소 전수 인덱싱 — blog.yaml 재로드 없이 선언 언어 순회 (IndexRepo의 본체)
package content

import "github.com/park-jun-woo/abloq/pkg/blogyaml"

// IndexEntries indexes every published article for an already-validated
// Blog — the loop body of IndexRepo without the blog.yaml reload. Callers
// that loaded (and validated) blog.yaml themselves use this to avoid a
// second load whose failure path could never fire.
func IndexEntries(root string, b *blogyaml.Blog) []Entry {
	entries := []Entry{}
	for _, lang := range b.Languages {
		entries = append(entries, indexLang(root, b, lang)...)
	}
	return entries
}
