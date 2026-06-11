//ff:type feature=image type=client
//ff:what Gemini 이미지 생성 provider — 모델명(echo용 노출)과 env에서 해석된 베이스 URL·API 키
//ff:why HTTP는 ogprovider에만 존재(img는 네트워크 0) — 키는 env 전용으로 blog.yaml·인자·OGSpec에 절대 싣지 않는다 (BUG002)
package ogprovider

// Gemini generates OG background images through the Gemini API
// (generateContent with inline base64 image responses). Construct it with
// NewGemini — the API key never travels outside this struct. Model is the
// effective model id, exposed so the CLI can echo it.
type Gemini struct {
	Model string
	base  string
	key   string
}
