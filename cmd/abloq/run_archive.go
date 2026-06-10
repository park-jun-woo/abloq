//ff:func feature=cli type=command control=iteration dimension=1
//ff:what URL 1건을 3종 kind로 전개해 pkg/archive.ProcessBatch 실행, kind별 status와 응답을 출력 — 하나라도 done이 아니면 에러
package main

import (
	"fmt"
	"io"

	"github.com/park-jun-woo/abloq/pkg/archive"
)

// runArchive submits one URL through the shared archiver batch. The CLI has
// no receipts table — results go to stdout; a non-done result exits 1 so CI
// can react.
func runArchive(w io.Writer, target string) error {
	pending := make([]archive.Pending, 0, len(archive.Kinds))
	for _, kind := range archive.Kinds {
		pending = append(pending, archive.Pending{Kind: kind, Target: target})
	}
	failed := 0
	for _, item := range archive.ProcessBatch(pending, int64(len(pending))) {
		fmt.Fprintf(w, "%s\t%s\t%s\n", item.Kind, item.Status, item.Response)
		if item.Status != archive.StatusDone {
			failed++
		}
	}
	if failed > 0 {
		return fmt.Errorf("%d of %d submissions did not land done", failed, len(archive.Kinds))
	}
	return nil
}
