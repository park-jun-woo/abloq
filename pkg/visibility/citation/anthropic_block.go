//ff:type feature=visibility type=schema topic=citation
//ff:what Anthropic Messages content 블록 1건 — 타입(text 등)과 web search 인용 목록
package citation

// anthropicBlock is one Messages API content block; text blocks carry the
// web-search citations.
type anthropicBlock struct {
	Type      string              `json:"type"`
	Citations []anthropicCitation `json:"citations"`
}
