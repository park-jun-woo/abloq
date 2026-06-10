//ff:func feature=gate type=frame control=sequence
//ff:what 특수 페이지 판별 — front matter에 layout 키가 있으면 전용 레이아웃 소유 페이지로 분류
//ff:why layout 페이지(about 등)는 글 템플릿이 아니라 전용 Hugo 레이아웃이 렌더하므로 글 모양 룰(구조·스키마·근거)의 대상이 아니다 — 무결성 룰(baseline 비교·slug-consistency·hreflang)은 여전히 적용 (Phase006 도그푸드: {lang}/dabel/about.md 12편 정책 결정)
package gate

// special reports whether the article is a layout-owned special page: its
// front matter declares a `layout` key. Special pages are rendered by a
// dedicated Hugo layout, not the article template, so the article-shape rules
// (image-first, image-attribution, section-order, heading-canonical,
// front-matter-schema, min-sources, numeric-claim-sourced) skip them.
// Integrity rules (baseline comparisons, slug-consistency, hreflang-complete,
// citation-exists) still apply.
func special(a *Article) bool {
	if !a.Doc.HasFM {
		return false
	}
	m, ok := fmMap(a.Doc.FrontMatter)
	if !ok {
		return false
	}
	_, has := m["layout"]
	return has
}
