//ff:func feature=image type=client control=sequence
//ff:what Gemini provider 생성 — GEMINI_API_KEY(GOOGLE_API_KEY fallback) env 키 필수, GEMINI_API_BASE 오버라이드, 빈 모델은 기본 모델
package ogprovider

import "errors"

// defaultGeminiModel is used when no model is declared anywhere. Pinned to the
// current-generation image model per BUG003 §원인B; --model and blog.yaml
// image.og.model still take precedence over this default.
const defaultGeminiModel = "gemini-3-pro-image"

// NewGemini resolves credentials and base URL from the environment. A missing
// API key is a hard error — the caller surfaces it as a clear exit-1 diagnosis.
func NewGemini(model string) (*Gemini, error) {
	key := envOr("GEMINI_API_KEY", envOr("GOOGLE_API_KEY", ""))
	if key == "" {
		return nil, errors.New("gemini provider needs an API key: set GEMINI_API_KEY (or GOOGLE_API_KEY)")
	}
	if model == "" {
		model = defaultGeminiModel
	}
	base := envOr("GEMINI_API_BASE", "https://generativelanguage.googleapis.com")
	return &Gemini{Model: model, base: base, key: key}, nil
}
