//ff:func feature=gate type=parser control=sequence topic=evidence
//ff:what HTML에서 <title>과 og:title 추출 — 메타 일치 검사의 비교 대상 목록
package gate

import "regexp"

var (
	reHTMLTitle   = regexp.MustCompile(`(?is)<title[^>]*>(.*?)</title>`)
	reOGTitleTag  = regexp.MustCompile(`(?is)<meta[^>]+property=["']og:title["'][^>]*>`)
	reMetaContent = regexp.MustCompile(`(?is)content=["']([^"']*)["']`)
)

// htmlTitles extracts the page's <title> text and og:title content (either
// attribute order) for the citation meta-match check.
func htmlTitles(html string) []string {
	var out []string
	if m := reHTMLTitle.FindStringSubmatch(html); m != nil {
		out = append(out, m[1])
	}
	tag := reOGTitleTag.FindString(html)
	if c := reMetaContent.FindStringSubmatch(tag); c != nil {
		out = append(out, c[1])
	}
	return out
}
