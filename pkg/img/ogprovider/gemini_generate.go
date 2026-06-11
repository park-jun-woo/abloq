//ff:func feature=image type=client control=sequence
//ff:what Gemini 이미지 생성 1발 — generateContent에 프롬프트 POST(키는 헤더), inline base64 응답을 image.Image로 환원
package ogprovider

import (
	"context"
	"image"
)

// Generate asks the Gemini API for one image. The raw (any-size) image comes
// back as an inline base64 part; post-processing to 1200x630 WebP is the
// caller's job (pkg/img pipeline).
func (g *Gemini) Generate(ctx context.Context, prompt string) (image.Image, error) {
	endpoint := g.base + "/v1beta/models/" + g.Model + ":generateContent"
	body, err := postJSON(ctx, endpoint, g.key, map[string]any{
		"contents": []map[string]any{
			{"parts": []map[string]string{{"text": prompt}}},
		},
		"generationConfig": map[string]any{
			"responseModalities": []string{"TEXT", "IMAGE"},
		},
	})
	if err != nil {
		return nil, err
	}
	return parseGeminiImage(body)
}
