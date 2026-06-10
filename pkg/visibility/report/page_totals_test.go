//ff:func feature=visibility type=parser control=sequence topic=report
//ff:what PageTotals가 page URL 경로를 URL맵으로 글에 귀속해 노출·클릭을 누적하고 미귀속 page는 버리는지 검증
package report

import (
	"testing"

	"github.com/park-jun-woo/abloq/pkg/visibility/cflog"
)

func TestPageTotals(t *testing.T) {
	urls := map[string]cflog.Article{
		"/tech/post-a/": {Lang: "ko", Section: "tech", Slug: "post-a"},
	}
	m := PageTotals([]PageSum{
		{Page: "https://fixture.example.com/tech/post-a/", Impressions: 120, Clicks: 8},
		{Page: "https://fixture.example.com/tech/post-a/", Impressions: 30, Clicks: 1},
		{Page: "https://fixture.example.com/", Impressions: 999, Clicks: 9},
	}, urls)
	got := m["ko/tech/post-a"]
	if got.Impressions != 150 || got.Clicks != 9 {
		t.Errorf("want 150/9, got %+v", got)
	}
	if len(m) != 1 {
		t.Errorf("unmapped pages must be dropped: %+v", m)
	}
}
