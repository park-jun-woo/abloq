//ff:func feature=cli type=command control=iteration dimension=1
//ff:what cluster 스캔 실행 본체 — 저장소 직접 파싱(기본 언어 그래프) → 백엔드와 같은 pkg/scan/cluster 판정 → quests/queue/ 기록
//ff:why CLI는 무상태다 — 그래프도 후보도 저장소 단일 소스의 결정적 연산이라, 같은 입력이면 endpoint export 산출과 바이트 동일(diff -r 0)해야 한다 (Phase011)
package main

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/park-jun-woo/abloq/pkg/queueio"
	"github.com/park-jun-woo/abloq/pkg/scan/cluster"
)

// runScanCluster runs one cluster pass on the blog repository at dir and
// writes the violating articles' queue files under dir/quests/queue/ — the
// same detection, candidate ranking and serialization the abloqd exporter
// pushes, so both outputs are byte-identical for the same repository state.
func runScanCluster(out io.Writer, dir string) error {
	b, err := loadValidBlog(out, dir)
	if err != nil {
		return err
	}
	items := cluster.Scan(dir, b)
	if err := queueio.WriteDir(filepath.Join(dir, "quests", "queue"), items); err != nil {
		return err
	}
	for _, it := range items {
		fmt.Fprintf(out, "%s\tpriority=%d\n", filepath.Join("quests", "queue", queueio.Filename(it)), it.Priority)
	}
	fmt.Fprintf(out, "cluster: %d article(s) queued (min internal links %d)\n", len(items), b.Geo.MinInternalLinks)
	return nil
}
