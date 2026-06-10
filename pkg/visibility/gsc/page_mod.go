//ff:type feature=visibility type=schema topic=gsc
//ff:what URL Inspection 후보 1건 — 정규 URL과 lastmod (저장소 파싱 결과의 부분집합)
package gsc

// PageMod is one inspection candidate: the canonical URL and the article's
// lastmod scalar ("YYYY-MM-DD..." ISO-8601, lexicographically sortable) as
// parsed from the repository — the single source of truth for freshness.
type PageMod struct {
	URL     string
	Lastmod string
}
