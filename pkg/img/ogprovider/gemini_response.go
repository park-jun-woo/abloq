//ff:type feature=image type=schema
//ff:what generateContent 응답의 이미지 추출용 부분 스키마 루트 — candidates 목록
package ogprovider

// geminiResponse is the subset of a generateContent answer the image
// extractor walks.
type geminiResponse struct {
	Candidates []geminiCandidate `json:"candidates"`
}
