//ff:type feature=image type=schema
//ff:what generateContent 이미지 part의 inlineData — MIME 타입과 base64 페이로드
package ogprovider

// geminiInlineData is the base64 image payload of an image part.
type geminiInlineData struct {
	MimeType string `json:"mimeType"`
	Data     string `json:"data"`
}
