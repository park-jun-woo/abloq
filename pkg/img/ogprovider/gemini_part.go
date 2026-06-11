//ff:type feature=image type=schema
//ff:what generateContent 응답 part 1개 — inlineData(base64 이미지)만 추출 대상, 텍스트 part는 nil
package ogprovider

// geminiPart is one content part. Text parts leave InlineData nil.
type geminiPart struct {
	InlineData *geminiInlineData `json:"inlineData"`
}
