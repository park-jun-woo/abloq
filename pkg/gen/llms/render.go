//ff:func feature=gen type=generator control=iteration dimension=1
//ff:what 발행 글 목록을 llms.txt로 렌더 — 사이트 헤더 + "## 언어/섹션" 그룹 + 글 1줄씩, 정렬 고정으로 멱등
package llms

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// Render produces static/llms.txt from the collected posts.
// Posts are re-sorted internally (language -> section -> newest date) so the
// output bytes never depend on the input order.
func Render(b *blogyaml.Blog, posts []Post) []byte {
	var sb strings.Builder
	fmt.Fprintf(&sb, "# %s\n\n> %s\n", b.Site.Title, headerNote(b))
	group := ""
	for _, p := range sortPosts(posts, b.Languages, b.Sections) {
		if g := p.Lang + "/" + p.Section; g != group {
			group = g
			fmt.Fprintf(&sb, "\n## %s\n\n", group)
		}
		sb.WriteString(postLine(b, p) + "\n")
	}
	return []byte(sb.String())
}
