//ff:type feature=image type=schema
//ff:what generateContent 응답의 candidate 1개 — content.parts 경로만
package ogprovider

// geminiCandidate is one response candidate; only content.parts matters here.
type geminiCandidate struct {
	Content geminiContent `json:"content"`
}
