//ff:type feature=visibility type=schema topic=citation
//ff:what OpenAI Responses 어노테이션 1건 — url_citation 타입이면 URL이 인용 출처
package citation

// oaiAnnotation is one annotation on a Responses API text part; the
// extractor keeps only type "url_citation".
type oaiAnnotation struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}
