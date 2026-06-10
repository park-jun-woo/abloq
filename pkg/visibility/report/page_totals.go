//ff:func feature=visibility type=parser control=iteration dimension=1 topic=report
//ff:what GSC page 합계 → URL맵으로 글 귀속한 집계 맵 — page URL의 경로를 저장소 역매핑에 조회, 미귀속 page는 버린다
package report

import (
	"net/url"

	"github.com/park-jun-woo/abloq/pkg/queueio"
	"github.com/park-jun-woo/abloq/pkg/visibility/cflog"
)

// PageTotals attributes the per-page GSC sums to articles through the
// repository URL map (cflog.BuildURLMap — the same single source the crawl
// ingest uses). Pages whose path maps to no article (home, lists, 404s)
// are dropped; unparseable page URLs likewise.
func PageTotals(sums []PageSum, urls map[string]cflog.Article) map[string]PageTally {
	m := make(map[string]PageTally, len(sums))
	for _, s := range sums {
		u, err := url.Parse(s.Page)
		if err != nil {
			continue
		}
		a, ok := urls[u.Path]
		if !ok {
			continue
		}
		key := queueio.JoinKey(a.Lang, a.Section, a.Slug)
		t := m[key]
		t.Impressions += s.Impressions
		t.Clicks += s.Clicks
		m[key] = t
	}
	return m
}
