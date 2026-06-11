//ff:func feature=blogyaml type=rule control=iteration dimension=1
//ff:what [llmstxt-pinned] geo.llms_txt.pinned 각 엔트리의 title·url 필수와 url 형식(절대 URL 또는 / 시작)을 검증
package blogyaml

import "fmt"

// ruleLlmsTxtPinned requires every pinned entry to carry a title and a url,
// the url being an absolute http(s) URL or a "/"-rooted path.
func ruleLlmsTxtPinned(filename string, b *Blog, idx lineIndex) []Diagnostic {
	var diags []Diagnostic
	for i, p := range b.Geo.LlmsTxt.Pinned {
		if p.Title == "" {
			diags = append(diags, Diagnostic{
				File: filename, Line: llmsTxtItemLine(idx, fmt.Sprintf("pinned[%d]", i), "pinned"),
				Rule:    "llmstxt-pinned",
				Message: fmt.Sprintf("geo.llms_txt.pinned[%d] title is required", i),
			})
		}
		if p.URL == "" {
			diags = append(diags, Diagnostic{
				File: filename, Line: llmsTxtItemLine(idx, fmt.Sprintf("pinned[%d]", i), "pinned"),
				Rule:    "llmstxt-pinned",
				Message: fmt.Sprintf("geo.llms_txt.pinned[%d] url is required", i),
			})
		}
		if p.URL != "" && !llmsPinnedURLOK(p.URL) {
			diags = append(diags, Diagnostic{
				File: filename, Line: llmsTxtItemLine(idx, fmt.Sprintf("pinned[%d].url", i), "pinned"),
				Rule:    "llmstxt-pinned",
				Message: fmt.Sprintf("geo.llms_txt.pinned[%d] url %q must be an absolute http(s) URL or start with /", i, p.URL),
			})
		}
	}
	return diags
}
