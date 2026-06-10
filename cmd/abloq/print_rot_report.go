//ff:func feature=cli type=output control=iteration dimension=1
//ff:what link rot 1회 점검 보고 출력 — 실패 인용만 "rot-check: URL 판정 (글 좌표)" 한 줄씩, 실패 수 반환
package main

import (
	"fmt"
	"io"

	"github.com/park-jun-woo/abloq/pkg/scan/evidence"
)

// printRotReport prints the failing citations of a one-shot check and returns
// how many failed. Healthy citations stay silent — the report is a worklist,
// not a log.
func printRotReport(out io.Writer, checks []evidence.Check) int {
	failing := 0
	for _, c := range checks {
		if c.Status == "ok" {
			continue
		}
		failing++
		fmt.Fprintf(out, "rot-check: %s %s (%s/%s/%s)\n", c.URL, c.Status, c.Lang, c.Section, c.Slug)
	}
	return failing
}
