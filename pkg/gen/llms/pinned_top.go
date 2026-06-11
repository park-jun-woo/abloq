//ff:func feature=gen type=generator control=iteration dimension=1
//ff:what group 미지정 pinned 엔트리를 목록 최상단 무헤딩 블록으로 렌더 — 선언 순서 유지, 없으면 빈 문자열
package llms

import (
	"strings"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// pinnedTop renders the ungrouped pinned entries as a heading-less block at
// the very top of the list, in declaration order.
func pinnedTop(b *blogyaml.Blog) string {
	var sb strings.Builder
	for _, p := range b.Geo.LlmsTxt.Pinned {
		if p.Group != "" {
			continue
		}
		sb.WriteString(pinnedLine(b, p) + "\n")
	}
	if sb.Len() == 0 {
		return ""
	}
	return "\n" + sb.String()
}
