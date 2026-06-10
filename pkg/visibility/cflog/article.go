//ff:type feature=visibility type=schema topic=crawl
//ff:what URI 역매핑 1건 — 경로가 가리키는 글(lang/section/slug)과 .md 병행 서빙 여부
package cflog

// Article is the value of the URI reverse map: which article a request path
// belongs to, and whether the path is the article's parallel-served .md
// (counted into md_hits instead of hits).
type Article struct {
	Lang    string
	Section string
	Slug    string
	MD      bool
}
