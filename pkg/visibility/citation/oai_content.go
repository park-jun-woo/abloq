//ff:type feature=visibility type=schema topic=citation
//ff:what OpenAI Responses 콘텐츠 파트 1건 — 타입(output_text)과 web search가 다는 어노테이션 목록
package citation

// oaiContent is one content part of a Responses API message output.
type oaiContent struct {
	Type        string          `json:"type"`
	Annotations []oaiAnnotation `json:"annotations"`
}
