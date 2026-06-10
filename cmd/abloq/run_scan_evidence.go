//ff:func feature=cli type=command control=sequence
//ff:what evidence 스캔 실행 본체 — 저장소 직접 스캔(prev=nil 무상태) → claims 항목만 quests/queue/ 기록 + rot 1회 점검 보고
//ff:why CLI는 DB가 없어 연속 실패를 셀 수 없다 — rot 확정(3회)은 백엔드 상태의 몫이고, claims 항목은 같은 입력이면 endpoint export 산출과 바이트 동일(diff -r 0)해야 한다 (Phase010)
package main

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/park-jun-woo/abloq/pkg/queueio"
	"github.com/park-jun-woo/abloq/pkg/scan/evidence"
)

// runScanEvidence runs one stateless evidence pass on the blog repository at
// dir: unsourced-claim queue files land under dir/quests/queue/ (the same
// bytes the abloqd exporter pushes), and every citation URL gets a one-shot
// liveness report — rot confirmation (3 consecutive failed scans) lives in
// the backend's citation_checks state, never here.
func runScanEvidence(out io.Writer, dir string) error {
	b, err := loadValidBlog(out, dir)
	if err != nil {
		return err
	}
	res := evidence.Scan(dir, b, nil, evidence.NewChecker())
	if err := queueio.WriteDir(filepath.Join(dir, "quests", "queue"), res.Items); err != nil {
		return err
	}
	printEvidenceQueue(out, res.Items)
	failing := printRotReport(out, res.Checks)
	fmt.Fprintf(out, "evidence: %d article(s) queued, %d citation(s) checked, %d failing (rot confirms after %d backend scans)\n",
		len(res.Items), len(res.Checks), failing, 3)
	return nil
}
