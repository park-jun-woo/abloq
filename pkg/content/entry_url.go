//ff:func feature=content type=generator control=sequence
//ff:what 인덱스 항목의 정규 URL 조립 — baseURL/언어/섹션/slug/ (Blog.URLLang 규칙: 루트 서빙 기본 언어는 언어 세그먼트 생략)
package content

import (
	"strings"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// entryURL builds the canonical article URL under the blog.yaml URL contract,
// matching the hugo permalink scheme (and pkg/gen/llms postURL).
func entryURL(b *blogyaml.Blog, lang, section, slug string) string {
	url := strings.TrimRight(b.Site.BaseURL, "/")
	if seg := b.URLLang(lang); seg != "" {
		url += "/" + seg
	}
	return url + "/" + section + "/" + slug + "/"
}
