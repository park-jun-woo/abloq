//ff:type feature=gen type=schema
//ff:what 발행 글 1편 — llms.txt 목록 항목의 입력 (언어/섹션/slug/제목/날짜/설명)
package llms

// Post is one published article collected from content/{lang}/{section}/.
// Date keeps the front matter scalar as-is (ISO-8601 sorts lexicographically).
type Post struct {
	Lang        string
	Section     string
	Slug        string
	Title       string
	Date        string
	Description string
}
