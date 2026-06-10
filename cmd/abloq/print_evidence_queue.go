//ff:func feature=cli type=output control=iteration dimension=1
//ff:what 기록된 evidence 큐 파일 목록 출력 — 파일 경로와 priority 한 줄씩 (freshness와 같은 형식)
package main

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/park-jun-woo/abloq/pkg/queueio"
)

// printEvidenceQueue lists the queue files runScanEvidence just wrote.
func printEvidenceQueue(out io.Writer, items []queueio.Item) {
	for _, it := range items {
		fmt.Fprintf(out, "%s\tpriority=%d\n", filepath.Join("quests", "queue", queueio.Filename(it)), it.Priority)
	}
}
