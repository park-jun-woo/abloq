//ff:type feature=visibility type=schema topic=citation
//ff:what Anthropic Messages 인용 1건 — web_search_result_location의 URL이 인용 출처
package citation

// anthropicCitation is one citation entry on a Messages API text block.
type anthropicCitation struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}
