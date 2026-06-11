//ff:func feature=cli type=output control=iteration dimension=1
//ff:what 매칭 결과 출력 — anchored n/m 요약과 미출현 claim 목록(id: text) — REVIEW 단계 제시용
package main

import (
	"fmt"
	"io"

	"github.com/park-jun-woo/abloq/pkg/insight"
)

// printInsightMatch writes the anchored summary and the missing-claim list.
func printInsightMatch(out io.Writer, res insight.Result, total int) {
	fmt.Fprintf(out, "anchored claims: %d/%d\n", len(res.Found), total)
	if len(res.Missing) == 0 {
		return
	}
	fmt.Fprintln(out, "missing claims (no anchor found in body — present to REVIEW):")
	for _, c := range res.Missing {
		fmt.Fprintf(out, "  %s: %s\n", c.ID, c.Text)
	}
}
