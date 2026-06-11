//ff:func feature=gen type=generator control=sequence
//ff:what pinned 엔트리 1줄 렌더 — "- [제목](URL)" + desc가 있으면 ": 설명", "/" 시작 url은 baseURL과 결합
package llms

import (
	"fmt"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// pinnedLine renders one pinned llms.txt list entry.
func pinnedLine(b *blogyaml.Blog, p blogyaml.LlmsPinned) string {
	line := fmt.Sprintf("- [%s](%s)", p.Title, pinnedURL(b, p.URL))
	if p.Desc != "" {
		line += ": " + p.Desc
	}
	return line
}
