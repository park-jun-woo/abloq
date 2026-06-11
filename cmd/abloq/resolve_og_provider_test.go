//ff:func feature=cli type=command control=sequence
//ff:what resolveOGProvider 검증 — gemini는 인스턴스+실효 모델 echo(기본 모델 포함), 미지 provider는 에러, 키 부재 전파
package main

import (
	"strings"
	"testing"
)

func TestResolveOGProvider(t *testing.T) {
	t.Setenv("GEMINI_API_KEY", "k")
	p, model, err := resolveOGProvider("gemini", "imagen-4")
	if err != nil || p == nil || model != "imagen-4" {
		t.Errorf("gemini explicit model: %v %q %v", p, model, err)
	}
	_, model, err = resolveOGProvider("gemini", "")
	if err != nil || model == "" {
		t.Errorf("gemini default model: want non-empty echo, got %q %v", model, err)
	}
	if _, _, err := resolveOGProvider("dalle", ""); err == nil || !strings.Contains(err.Error(), "dalle") {
		t.Errorf("unknown provider: want error, got %v", err)
	}
	t.Setenv("GEMINI_API_KEY", "")
	t.Setenv("GOOGLE_API_KEY", "")
	if _, _, err := resolveOGProvider("gemini", ""); err == nil || !strings.Contains(err.Error(), "GEMINI_API_KEY") {
		t.Errorf("missing key: want clear diagnosis, got %v", err)
	}
}
