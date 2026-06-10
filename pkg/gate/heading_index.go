//ff:type feature=gate type=schema
//ff:what blog.yaml structure에서 파생한 헤딩 인덱스 — 언어별 정규화 헤딩→키 역색인과 키→정규 순위
package gate

// headingIndex is the per-target derived view of blog.yaml structure:
// byLang resolves a normalized heading text to its heading key, and rank is
// the canonical position of each structure.order entry.
type headingIndex struct {
	byLang map[string]map[string]string // lang -> normText(heading) -> heading key
	rank   map[string]int               // structure.order entry -> index
}
