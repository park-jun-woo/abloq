//ff:func feature=gen type=generator control=sequence
//ff:what geo.llms_txt.header 포지셔닝 블록 렌더 — 자유 마크다운 그대로, trailing 개행만 정규화, 없으면 빈 문자열
package llms

import (
	"strings"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// headerBlock renders the free-markdown positioning block below the site
// header. Only trailing newlines are normalized; absent header -> "".
func headerBlock(b *blogyaml.Blog) string {
	h := strings.TrimRight(b.Geo.LlmsTxt.Header, "\n")
	if h == "" {
		return ""
	}
	return "\n" + h + "\n"
}
