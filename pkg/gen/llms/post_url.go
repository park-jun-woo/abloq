//ff:func feature=gen type=generator control=sequence
//ff:what 발행 글의 정규 URL 조립 — baseURL/언어/섹션/slug/ (섹션 permalink와 일치, 루트 서빙 기본 언어는 언어 세그먼트 생략)
package llms

import (
	"strings"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// postURL builds the canonical article URL matching the hugo permalink scheme.
// The language segment follows Blog.URLLang — empty for a root-served default
// language (site.default_lang_in_subdir: false).
func postURL(b *blogyaml.Blog, p Post) string {
	url := strings.TrimRight(b.Site.BaseURL, "/")
	if seg := b.URLLang(p.Lang); seg != "" {
		url += "/" + seg
	}
	return url + "/" + p.Section + "/" + p.Slug + "/"
}
