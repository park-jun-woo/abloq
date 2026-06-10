//ff:func feature=visibility type=generator control=iteration dimension=1 topic=report
//ff:what 미지 봇 markdown 절 — UA·히트 표, 비면 "none" 한 줄
package report

import (
	"fmt"
	"strings"
)

// unknownSection renders the unknown-bot candidate table — the operator's
// dictionary-update input. An empty list reads "none".
func unknownSection(bots []UnknownBot) string {
	if len(bots) == 0 {
		return "none\n"
	}
	var b strings.Builder
	b.WriteString("| ua | hits |\n|---|---:|\n")
	for _, u := range bots {
		fmt.Fprintf(&b, "| %s | %d |\n", u.UA, u.Hits)
	}
	return b.String()
}
