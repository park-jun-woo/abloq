//ff:func feature=cli type=command control=sequence
//ff:what buildOGRuns 검증 — 안별 Provider 인스턴스 주입·실효 모델 echo·프롬프트 치환(opts.Summary 주입), provider 해석 실패 즉시 에러
package main

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func TestBuildOGRuns(t *testing.T) {
	t.Setenv("GEMINI_API_KEY", "k")
	specs := []blogyaml.OGVariantSpec{
		{Name: "minimal", Model: "imagen-4", Overlay: true, Prompt: "Art for {title} by {brand}{summary}"},
		{Name: "photo", Model: ""},
	}
	opts := imageOGOpts{Title: "T", Brand: "B", Summary: "S"}
	runs, err := buildOGRuns("gemini", specs, opts)
	if err != nil || len(runs) != 2 {
		t.Fatalf("runs = %+v, err = %v", runs, err)
	}
	if runs[0].Name != "minimal" || runs[0].Model != "imagen-4" || !runs[0].Overlay || runs[0].Provider == nil {
		t.Errorf("run[0] = %+v, want injected provider with explicit model", runs[0])
	}
	if runs[0].Prompt != "Art for T by BS" {
		t.Errorf("prompt = %q, want substituted with opts.Summary", runs[0].Prompt)
	}
	if runs[1].Model == "" || runs[1].Provider == nil {
		t.Errorf("run[1] = %+v, want default model echo", runs[1])
	}

	// provider resolution failure aborts before any run is built
	if _, err := buildOGRuns("dalle", specs, opts); err == nil || !strings.Contains(err.Error(), "dalle") {
		t.Errorf("unknown provider: want error, got %v", err)
	}
}
