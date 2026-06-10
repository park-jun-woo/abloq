//ff:func feature=claudemd type=generator control=sequence
//ff:what CLAUDE.md 글 구조 계약 섹션 — structure.order, 헤딩 표, 본문 헤딩·front matter 규칙, 게이트 임계값
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
	sb.WriteString("본문에 H1(`#`)을 쓰지 않는다 — 제목은 front matter `title`이고 레이아웃이 렌더한다. 본문 헤딩은 `##`부터.\n\n")
	sb.WriteString("front matter 필수: `title`(비공백 문자열) · `date` · `lastmod`(date 이후) · `tags`(1개 이상)\n")
	sb.WriteString("front matter 권장: `summary`(한 줄 요약 — llms.txt 설명·meta description) · `image`(`/images/{slug}.webp` — OG 카드)\n")
	sb.WriteString("`date`/`lastmod`는 타임존 포함 RFC3339, **현재 시각 이전**으로 — hugo는 미래 날짜 글을 빌드하지 않는다.\n\n")
	fmt.Fprintf(&sb, "게이트 임계값: min_sources=%d · min_internal_links=%d · freshness_days=%d · min_meaningful_diff=%d\n\n",
		b.Geo.MinSources, b.Geo.MinInternalLinks, b.Geo.FreshnessDays, b.Geo.MinMeaningfulDiff)
	return sb.String()
}
