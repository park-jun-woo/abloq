//ff:func feature=gen type=generator control=iteration dimension=1
//ff:what 정렬된 발행 글을 "## 그룹 헤딩" 단위로 렌더 — 그룹 전환 시 헤딩과 합류 pinned 선두 줄을 먼저 출력
package llms

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// postGroups renders the section groups in blog.yaml declaration order
// (guaranteed by the pre-sorted input). On each group change it emits the
// heading and any pinned entries merged into that group, then the post lines.
func postGroups(b *blogyaml.Blog, sorted []Post, multi bool) string {
	var sb strings.Builder
	group := ""
	for _, p := range sorted {
		if g := p.Lang + "/" + p.Section; g != group {
			group = g
			head := headingText(b, p.Lang, p.Section, multi)
			fmt.Fprintf(&sb, "\n## %s\n\n", head)
			sb.WriteString(sectionPinned(b, head))
		}
		sb.WriteString(postLine(b, p) + "\n")
	}
	return sb.String()
}
