//ff:func feature=claudemd type=generator control=sequence
//ff:what CLAUDE.md 글 구조 계약 섹션 — structure.order, 헤딩 표, front matter 스키마, 게이트 임계값
package claudemd

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// structureSection renders the article structure contract the gates enforce.
func structureSection(b *blogyaml.Blog) string {
	var sb strings.Builder
	sb.WriteString("## 글 구조 계약 (구조 게이트가 강제)\n\n")
	if len(b.Structure.Order) > 0 {
		fmt.Fprintf(&sb, "정규 섹션 순서: %s\n\n", strings.Join(b.Structure.Order, " → "))
	}
	sb.WriteString(headingTable(b))
	sb.WriteString("front matter 필수: `title`(비공백 문자열) · `date` · `lastmod`(date 이후) · `tags`(1개 이상)\n\n")
	fmt.Fprintf(&sb, "게이트 임계값: min_sources=%d · min_internal_links=%d · freshness_days=%d · min_meaningful_diff=%d\n\n",
		b.Geo.MinSources, b.Geo.MinInternalLinks, b.Geo.FreshnessDays, b.Geo.MinMeaningfulDiff)
	return sb.String()
}
