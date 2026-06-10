//ff:func feature=visibility type=generator control=iteration dimension=1 topic=report
//ff:what 큐 적재 요약 markdown 절 — kind·status·건수 표, 비면 "no queue intake in this window" 한 줄
package report

import (
	"fmt"
	"strings"
)

// queueSection renders the per-(kind, status) queue intake table of the
// window. An empty window reads as a fixed one-liner.
func queueSection(counts []QueueCount) string {
	if len(counts) == 0 {
		return "no queue intake in this window\n"
	}
	var b strings.Builder
	b.WriteString("| kind | status | count |\n|---|---|---:|\n")
	for _, c := range counts {
		fmt.Fprintf(&b, "| %s | %s | %d |\n", c.Kind, c.Status, c.Count)
	}
	return b.String()
}
