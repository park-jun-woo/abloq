//ff:func feature=claudemd type=generator control=iteration dimension=1
//ff:what structure.headings를 언어별 헤딩 마크다운 표로 렌더 — 키 정렬로 결정적, 헤딩 없으면 빈 문자열
package claudemd

import (
	"sort"
	"strings"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// headingTable renders the localized heading table (## level is canonical).
func headingTable(b *blogyaml.Blog) string {
	if len(b.Structure.Headings) == 0 {
		return ""
	}
	keys := make([]string, 0, len(b.Structure.Headings))
	for k := range b.Structure.Headings {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var sb strings.Builder
	sb.WriteString("인식 헤딩 (`##` 레벨 고정):\n\n")
	sb.WriteString("| 키 | " + strings.Join(b.Languages, " | ") + " |\n")
	sb.WriteString("|---|" + strings.Repeat("---|", len(b.Languages)) + "\n")
	for _, key := range keys {
		sb.WriteString(headingRow(key, b.Languages, b.Structure.Headings[key]))
	}
	sb.WriteString("\n")
	return sb.String()
}
