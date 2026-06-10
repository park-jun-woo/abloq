//ff:func feature=gen type=generator control=sequence
//ff:what 발행 글의 정규 URL 조립 — baseURL/언어/섹션/slug/ (섹션 permalink와 일치)
package llms

import "strings"

// postURL builds the canonical article URL matching the hugo permalink scheme.
func postURL(baseURL string, p Post) string {
	return strings.TrimRight(baseURL, "/") + "/" + p.Lang + "/" + p.Section + "/" + p.Slug + "/"
}
