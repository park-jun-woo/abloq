//ff:type feature=content type=schema
//ff:what 콘텐츠 인덱스 항목 1건 — 언어/섹션/slug/제목/날짜/본문 지표/태그/정규 URL, posts 테이블 1행과 1:1
//ff:why JSON 태그가 posts 테이블 컬럼명과 일치해야 한다 — 백엔드 @call 래퍼가 이 배열을 그대로 jsonb_to_recordset 업서트에 먹인다 (Phase007)
package content

// Entry is one indexed article. Date and Lastmod keep the front matter scalar
// as-is (ISO-8601 sorts lexicographically); Lastmod falls back to Date.
// JSON keys mirror the abloqd posts table columns.
type Entry struct {
	Lang          string   `json:"lang"`
	Section       string   `json:"section"`
	Slug          string   `json:"slug"`
	Title         string   `json:"title"`
	Date          string   `json:"date"`
	Lastmod       string   `json:"lastmod"`
	WordCount     int64    `json:"word_count"`
	Tags          []string `json:"tags"`
	InternalLinks int64    `json:"internal_links"`
	SourceCount   int64    `json:"source_count"`
	URL           string   `json:"url"`
	// Summary is the article's one-line abstract (description > summary merge),
	// surfaced for in-process consumers like OG prompt wiring. json:"-" keeps it
	// OUT of the posts upsert contract: the JSON keys above mirror the posts
	// table columns 1:1 (see //ff:why), and a stray summary key would break the
	// jsonb_to_recordset ingest. Internal-only field — never serialized.
	Summary string `json:"-"`
}
