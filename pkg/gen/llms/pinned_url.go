//ff:func feature=gen type=generator control=sequence
//ff:what pinned url 환원 — "/" 시작 경로는 baseURL과 결합해 절대 URL로, 절대 URL은 그대로
package llms

import (
	"strings"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// pinnedURL resolves a pinned url: "/"-rooted paths join the site baseURL,
// absolute URLs pass through unchanged.
func pinnedURL(b *blogyaml.Blog, u string) string {
	if strings.HasPrefix(u, "/") {
		return strings.TrimRight(b.Site.BaseURL, "/") + u
	}
	return u
}
