//ff:func feature=gen type=generator control=sequence
//ff:what llms.txt 목록 항목 1줄 렌더 — "- [제목](URL)" + 설명이 있으면 ": 설명"(max_summary rune 절단 적용)
package llms

import (
	"fmt"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// postLine renders one llms.txt list entry, the description capped at
// geo.llms_txt.max_summary runes.
func postLine(b *blogyaml.Blog, p Post) string {
	line := fmt.Sprintf("- [%s](%s)", p.Title, postURL(b, p))
	if p.Description != "" {
		line += ": " + truncateSummary(p.Description, b.Geo.LlmsTxt.MaxSummary)
	}
	return line
}
