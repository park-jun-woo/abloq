//ff:type feature=visibility type=schema topic=citation
//ff:what OpenAI Responses API 응답의 인용 추출용 부분 스키마 루트 — output 메시지 목록
package citation

// oaiResponse is the subset of an OpenAI Responses API answer the citation
// extractor walks.
type oaiResponse struct {
	Output []oaiOutput `json:"output"`
}
