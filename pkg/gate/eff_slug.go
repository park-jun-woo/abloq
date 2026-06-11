//ff:func feature=gate type=rule control=sequence
//ff:what 글의 유효 slug 결정 — front matter slug가 있으면 그 값, 없으면 파일 어간 (공개 API)
//ff:why Phase017 번역 퀘스트의 스코프드 slug-consistency가 원문↔번역 유효 slug를 아이템 단위로 비교해야 해서 export — 재구현(복제) 대신 단일 출처 유지
package gate

// EffSlug returns the slug Hugo uses for the article's URL. Exported for the
// translation quest's scoped slug-consistency rule (Phase017).
func EffSlug(a *Article) string {
	if s := fmLineValue(a.Doc.FrontMatter, "slug"); s != "" {
		return s
	}
	return a.Slug
}
