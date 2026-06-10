//ff:type feature=visibility type=schema topic=citation
//ff:what Anthropic Messages API 응답의 인용 추출용 부분 스키마 루트 — content 블록 목록
package citation

// anthropicResponse is the subset of an Anthropic Messages answer the
// citation extractor walks.
type anthropicResponse struct {
	Content []anthropicBlock `json:"content"`
}
