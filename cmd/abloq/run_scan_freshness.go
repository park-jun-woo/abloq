//ff:func feature=cli type=command control=iteration dimension=1
//ff:what freshness 스캔 실행 본체 — 저장소 직접 파싱 → 백엔드와 같은 pkg/scan/freshness 판정 + 합성 스코어러 → quests/queue/ 기록
//ff:why CLI는 로컬에 측정 데이터가 없으므로 신호 맵이 비어 합성 스코어러가 콜드스타트로 자연 폴백한다 — 같은 입력이면 endpoint export 산출과 바이트 동일(diff -r 0) (Phase009, Phase014 합성 주입)
package main

import (
	"fmt"
	"io"
	"path/filepath"
	"time"

	"github.com/park-jun-woo/abloq/pkg/content"
	"github.com/park-jun-woo/abloq/pkg/queueio"
	"github.com/park-jun-woo/abloq/pkg/scan/freshness"
	"github.com/park-jun-woo/abloq/pkg/visibility/priority"
)

// runScanFreshness detects stale articles in the blog repository at dir and
// writes their queue files under dir/quests/queue/ — the same detection and
// serialization the abloqd exporter pushes, so both outputs are
// byte-identical for the same repository state.
func runScanFreshness(out io.Writer, dir string) error {
	b, err := loadValidBlog(out, dir)
	if err != nil {
		return err
	}
	entries, err := content.IndexRepo(dir)
	if err != nil {
		return err
	}
	scorer := priority.Composite{W: priority.WeightsOf(b.Geo.PriorityWeights)}
	items := freshness.Scan(entries, map[string]priority.Signals{}, b.Languages, b.Geo.FreshnessDays, time.Now().UTC(), scorer)
	queueDir := filepath.Join(dir, "quests", "queue")
	if err := queueio.WriteDir(queueDir, items); err != nil {
		return err
	}
	for _, it := range items {
		fmt.Fprintf(out, "%s\tpriority=%d\n", filepath.Join("quests", "queue", queueio.Filename(it)), it.Priority)
	}
	fmt.Fprintf(out, "freshness: %d stale article(s) queued (threshold %d days)\n", len(items), b.Geo.FreshnessDays)
	return nil
}
