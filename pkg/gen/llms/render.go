//ff:func feature=gen type=generator control=sequence
//ff:what 발행 글 목록을 llms.txt로 렌더 — 사이트 헤더 → header 블록 → pinned → 섹션 그룹(라벨·단일 언어 접두 제거) 고정 순서, 멱등
//ff:why 렌더 순서를 사양으로 고정(Phase021) — 같은 입력이면 바이트 동일해야 check의 드리프트 게이트가 선다
package llms

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// Render produces static/llms.txt from the collected posts. The fixed order:
// "# Title" / "> author — baseURL" / header block / pinned entries /
// section groups in declaration order. Posts are re-sorted internally
// (language -> section -> newest date) so the output bytes never depend on
// the input order.
func Render(b *blogyaml.Blog, posts []Post) []byte {
	sorted := sortPosts(posts, b.Languages, b.Sections)
	multi := len(scopeLangs(b)) > 1
	heads := sectionHeadings(b, sorted, multi)
	var sb strings.Builder
	fmt.Fprintf(&sb, "# %s\n\n> %s\n", b.Site.Title, headerNote(b))
	sb.WriteString(headerBlock(b))
	sb.WriteString(pinnedTop(b))
	sb.WriteString(pinnedGroups(b, heads))
	sb.WriteString(postGroups(b, sorted, multi))
	return []byte(sb.String())
}
