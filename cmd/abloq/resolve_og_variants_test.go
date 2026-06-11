//ff:func feature=cli type=command control=sequence
//ff:what resolveOGVariants 검증 — 기본 경로의 플래그 오버라이드, --variant 시 안 선언 우선, --all-variants 빈 선언 에러
package main

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func TestResolveOGVariants(t *testing.T) {
	model := "variant-model"
	og := blogyaml.ImageOG{
		Provider: "gemini", Model: "site-model", Overlay: true, Prompt: "p",
		Variants: []blogyaml.OGVariant{{Name: "minimal", Model: &model}},
	}

	// default path: explicit flags override blog.yaml site-wide values
	opts := imageOGOpts{Model: "flag-model", ModelSet: true, Overlay: false, OverlaySet: true}
	specs, err := resolveOGVariants(og, opts)
	if err != nil || len(specs) != 1 {
		t.Fatalf("default path: %v %v", specs, err)
	}
	if specs[0].Name != "default" || specs[0].Model != "flag-model" || specs[0].Overlay {
		t.Errorf("default spec = %+v, want flag overrides", specs[0])
	}

	// default path without flags: blog.yaml values stand
	specs, _ = resolveOGVariants(og, imageOGOpts{})
	if specs[0].Model != "site-model" || !specs[0].Overlay {
		t.Errorf("default spec = %+v, want blog.yaml values", specs[0])
	}

	// --variant: variant declaration wins over --model flag
	opts = imageOGOpts{VariantList: "minimal", Model: "flag-model", ModelSet: true}
	specs, err = resolveOGVariants(og, opts)
	if err != nil || specs[0].Model != "variant-model" {
		t.Errorf("--variant: %+v %v, want variant-model to win", specs, err)
	}

	// --all-variants: every declaration, merged over the defaults
	specs, err = resolveOGVariants(og, imageOGOpts{AllVariants: true})
	if err != nil || len(specs) != 1 || specs[0].Name != "minimal" || specs[0].Model != "variant-model" {
		t.Errorf("--all-variants: %+v %v, want the declared variant", specs, err)
	}

	// --all-variants with no declarations
	if _, err := resolveOGVariants(blogyaml.ImageOG{}, imageOGOpts{AllVariants: true}); err == nil ||
		!strings.Contains(err.Error(), "no variants") {
		t.Errorf("--all-variants empty: want error, got %v", err)
	}
}
