//ff:type feature=visibility type=schema topic=citation
//ff:what OpenAI Responses output 항목 1건 — 타입(message/web_search_call)과 콘텐츠 파트 목록
package citation

// oaiOutput is one Responses API output item; only message items carry
// content parts.
type oaiOutput struct {
	Type    string       `json:"type"`
	Content []oaiContent `json:"content"`
}
