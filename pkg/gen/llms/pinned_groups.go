//ff:func feature=gen type=generator control=iteration dimension=1
//ff:what 섹션 그룹과 합류하지 않는 pinned group을 자체 헤딩 그룹으로 렌더 — 선언 순서, 연속 동일 group은 한 헤딩 아래로
package llms

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// pinnedGroups renders pinned entries whose group names no post-derived
// section heading: each group gets its own heading in declaration order,
// before the section groups (e.g. a "Core Content" lead group).
func pinnedGroups(b *blogyaml.Blog, sectionHeads map[string]bool) string {
	var sb strings.Builder
	group := ""
	for _, p := range b.Geo.LlmsTxt.Pinned {
		if p.Group == "" || sectionHeads[p.Group] {
			continue
		}
		if p.Group != group {
			group = p.Group
			fmt.Fprintf(&sb, "\n## %s\n\n", group)
		}
		sb.WriteString(pinnedLine(b, p) + "\n")
	}
	return sb.String()
}
