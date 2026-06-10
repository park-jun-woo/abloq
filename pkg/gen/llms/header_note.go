//ff:func feature=gen type=generator control=sequence
//ff:what llms.txt 인용구 헤더 1줄 조립 — "저자 — baseURL", 저자 미지정이면 baseURL만
package llms

import "github.com/park-jun-woo/abloq/pkg/blogyaml"

// headerNote builds the "> author — baseURL" blockquote line body.
func headerNote(b *blogyaml.Blog) string {
	if b.Site.Author == "" {
		return b.Site.BaseURL
	}
	return b.Site.Author + " — " + b.Site.BaseURL
}
