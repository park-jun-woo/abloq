//ff:func feature=visibility type=client control=iteration dimension=2 topic=citation
//ff:what OpenAI 콘텐츠 파트들에서 url_citation 어노테이션의 URL만 수집 — 다른 어노테이션 타입은 무시
package citation

// urlCitations walks the content parts of one Responses API output message
// and collects the web-search url_citation URLs.
func urlCitations(contents []oaiContent) []string {
	var urls []string
	for _, content := range contents {
		for _, a := range content.Annotations {
			if a.Type == "url_citation" && a.URL != "" {
				urls = append(urls, a.URL)
			}
		}
	}
	return urls
}
