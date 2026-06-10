//ff:func feature=cli type=command control=iteration dimension=1
//ff:what freshness 스캔 실행 본체 — 저장소 직접 파싱(콜드스타트, 측정 데이터 없음) → 백엔드와 같은 pkg/scan/freshness 판정 → quests/queue/ 기록
//ff:why CLI는 로컬에 crawl_hits가 없으므로 영구 콜드스타트다 — 같은 입력이면 endpoint export 산출과 바이트 동일(diff -r 0)해야 한다 (Phase009)
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
	items := freshness.Scan(entries, map[string]int64{}, b.Geo.FreshnessDays, time.Now().UTC(), priority.ColdStart{})
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
