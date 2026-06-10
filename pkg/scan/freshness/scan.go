//ff:func feature=scan type=rule control=iteration dimension=1
//ff:what freshness_days 초과 글 검출 — []content.Entry + 글별 신호 맵 → 우선순위 매긴 refresh 큐 후보 (CLI·백엔드 공유 순수 로직)
//ff:why 입력이 []content.Entry라 CLI(저장소 직접 파싱)와 백엔드(posts jsonb_agg)가 같은 판정을 공유한다 — payload에 now-파생값 금지(멱등·diff 판정의 전제). Phase014: hits 맵을 신호 구조체 맵으로 교체 — stale 판정 로직은 불변, 신호 통로만 바뀐다
package freshness

import (
	"time"

	"github.com/park-jun-woo/abloq/pkg/content"
	"github.com/park-jun-woo/abloq/pkg/queueio"
	"github.com/park-jun-woo/abloq/pkg/visibility/priority"
)

// Scan detects entries whose lastmod exceeded the freshness window and
// returns them as refresh queue candidates. The signals map carries each
// article's visibility signals (empty on the cold-start CLI path); the
// payload records the generation rationale (lastmod + threshold) without
// any now-derived value, keeping the queue file serialization deterministic.
func Scan(entries []content.Entry, signals map[string]priority.Signals, days int, now time.Time, scorer priority.Scorer) []queueio.Item {
	items := make([]queueio.Item, 0)
	for _, e := range entries {
		if !isStale(e.Lastmod, days, now) {
			continue
		}
		items = append(items, candidate(e, signals, days, scorer))
	}
	return items
}
