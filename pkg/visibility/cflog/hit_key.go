//ff:type feature=visibility type=schema topic=crawl
//ff:what 히트 누적의 맵 키 — crawl_hits 유니크 키(일자, 봇, lang, section, slug)와 동형
package cflog

// hitKey is the crawl_hits unique key.
type hitKey struct {
	Date    string
	Bot     string
	Lang    string
	Section string
	Slug    string
}
