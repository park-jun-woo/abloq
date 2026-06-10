//ff:func feature=queueio type=client control=sequence
//ff:what export 1회전 — 클론 최신화 → consumed 검출 → open 파일 기록 → 커밋·푸시 → 전이 대상 id 반환 (Phase010·011 재사용)
package queueio

import "path/filepath"

// Export runs one queue export cycle on the dedicated work clone: pull the
// agents' latest commits, detect consumed files (exported rows whose file was
// deleted), write every open item's file and push. Open rows become exported
// only after the push succeeds — on error nothing transitions and the next
// cycle retries idempotently.
func Export(cfg Config, open, exported []Row) (Result, error) {
	if err := ensureClone(cfg); err != nil {
		return Result{}, err
	}
	queueDir := filepath.Join(cfg.Workdir, "quests", "queue")
	consumed := consumedIDs(queueDir, exported)
	if err := WriteDir(queueDir, rowItems(open)); err != nil {
		return Result{}, err
	}
	if err := commitPush(cfg); err != nil {
		return Result{}, err
	}
	return Result{ExportedIDs: rowIDs(open), ConsumedIDs: consumed}, nil
}
