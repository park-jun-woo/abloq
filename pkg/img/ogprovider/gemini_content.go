//ff:type feature=image type=schema
//ff:what generateContent 응답 candidate의 content — parts 목록 보유
package ogprovider

// geminiContent holds the candidate's parts.
type geminiContent struct {
	Parts []geminiPart `json:"parts"`
}
