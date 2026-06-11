//ff:func feature=gen type=generator control=iteration dimension=1
//ff:what 특정 섹션 그룹 헤딩과 group이 일치하는 pinned 엔트리를 그 그룹 선두 줄들로 렌더 — 선언 순서 유지
package llms

import (
	"strings"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// sectionPinned renders the pinned entries leading one section group (those
// whose group equals the group's heading text), in declaration order.
func sectionPinned(b *blogyaml.Blog, heading string) string {
	var sb strings.Builder
	for _, p := range b.Geo.LlmsTxt.Pinned {
		if p.Group != heading {
			continue
		}
		sb.WriteString(pinnedLine(b, p) + "\n")
	}
	return sb.String()
}
