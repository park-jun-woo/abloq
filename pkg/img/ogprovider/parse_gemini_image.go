//ff:func feature=image type=parser control=iteration dimension=1
//ff:what generateContent 응답에서 첫 inlineData 이미지 part를 찾아 base64 디코드 → image.Image (텍스트 part는 건너뜀)
package ogprovider

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"image"

	"github.com/park-jun-woo/abloq/pkg/img"
)

// parseGeminiImage extracts the first inline image from a generateContent
// response. Text parts (the model may narrate) are skipped.
func parseGeminiImage(body []byte) (image.Image, error) {
	var resp geminiResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	if len(resp.Candidates) == 0 {
		return nil, errors.New("gemini api: response has no candidates")
	}
	for _, part := range resp.Candidates[0].Content.Parts {
		if part.InlineData == nil || part.InlineData.Data == "" {
			continue
		}
		raw, err := base64.StdEncoding.DecodeString(part.InlineData.Data)
		if err != nil {
			return nil, err
		}
		return img.DecodeBytes(raw)
	}
	return nil, errors.New("gemini api: response has no inline image part")
}
