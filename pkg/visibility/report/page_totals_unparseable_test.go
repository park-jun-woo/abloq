//ff:func feature=visibility type=parser control=sequence topic=report
//ff:what PageTotals가 파싱 불가 page URL을 버리는지 검증
package report

import (
	"testing"

	"github.com/park-jun-woo/abloq/pkg/visibility/cflog"
)

func TestPageTotalsUnparseable(t *testing.T) {
	urls := map[string]cflog.Article{"/tech/post-a/": {Lang: "ko", Section: "tech", Slug: "post-a"}}
	m := PageTotals([]PageSum{{Page: "http://[::1", Impressions: 5}}, urls)
	if len(m) != 0 {
		t.Errorf("unparseable page URLs must be dropped: %+v", m)
	}
}
