//ff:type feature=scan type=schema topic=evidence
//ff:what 인용 발생 1건 — citation_checks 키 좌표(lang/section/slug)와 URL, 점검 대상 수집의 단위
package evidence

// cite is one citation occurrence to probe: the citation_checks key. The same
// URL cited twice in one article collapses into one cite (the batch upsert
// must not touch a key twice), but the same URL across articles stays
// distinct.
type cite struct {
	Lang    string
	Section string
	Slug    string
	URL     string
}
