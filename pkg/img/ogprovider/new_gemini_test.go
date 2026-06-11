//ff:func feature=image type=client control=sequence
//ff:what NewGemini의 env 키 해석(GEMINI_API_KEY 우선, GOOGLE_API_KEY fallback, 부재 시 에러)·기본 모델·베이스 오버라이드 검증
package ogprovider

import (
	"strings"
	"testing"
)

func TestNewGemini(t *testing.T) {
	t.Setenv("GEMINI_API_KEY", "")
	t.Setenv("GOOGLE_API_KEY", "")
	t.Setenv("GEMINI_API_BASE", "")
	if _, err := NewGemini(""); err == nil || !strings.Contains(err.Error(), "GEMINI_API_KEY") {
		t.Errorf("missing key: want clear diagnosis, got %v", err)
	}

	t.Setenv("GOOGLE_API_KEY", "google-key")
	g, err := NewGemini("")
	if err != nil {
		t.Fatalf("GOOGLE_API_KEY fallback: %v", err)
	}
	if g.key != "google-key" {
		t.Errorf("key = %q, want GOOGLE_API_KEY fallback", g.key)
	}
	if g.Model != defaultGeminiModel {
		t.Errorf("model = %q, want default %q", g.Model, defaultGeminiModel)
	}
	if g.base != "https://generativelanguage.googleapis.com" {
		t.Errorf("base = %q, want default", g.base)
	}

	t.Setenv("GEMINI_API_KEY", "gemini-key")
	t.Setenv("GEMINI_API_BASE", "http://127.0.0.1:1")
	g, err = NewGemini("imagen-4")
	if err != nil {
		t.Fatalf("NewGemini: %v", err)
	}
	if g.key != "gemini-key" {
		t.Errorf("key = %q, want GEMINI_API_KEY to win over GOOGLE_API_KEY", g.key)
	}
	if g.Model != "imagen-4" || g.base != "http://127.0.0.1:1" {
		t.Errorf("model/base = %q/%q, want explicit values", g.Model, g.base)
	}
}
