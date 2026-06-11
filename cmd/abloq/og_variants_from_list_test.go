//ff:func feature=cli type=command control=sequence
//ff:what ogVariantsFromList 검증 — 콤마 목록을 요청 순서대로 병합 스펙으로 해석, 공백 트림, 미선언 이름 에러
package main

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func TestOGVariantsFromList(t *testing.T) {
	mA, mB := "model-a", "model-b"
	og := blogyaml.ImageOG{
		Model: "site-model", Overlay: true, Prompt: "p",
		Variants: []blogyaml.OGVariant{{Name: "a", Model: &mA}, {Name: "b", Model: &mB}},
	}

	// requested order preserved, whitespace trimmed, fields merged over defaults
	specs, err := ogVariantsFromList(og, " b , a ")
	if err != nil || len(specs) != 2 {
		t.Fatalf("specs = %+v, err = %v", specs, err)
	}
	if specs[0].Name != "b" || specs[0].Model != "model-b" || specs[1].Name != "a" || specs[1].Model != "model-a" {
		t.Errorf("specs = %+v, want b then a with variant models", specs)
	}
	if !specs[0].Overlay || specs[0].Prompt != "p" {
		t.Errorf("spec b = %+v, want inherited overlay/prompt", specs[0])
	}

	// undeclared name fails with the offending name in the message
	if _, err := ogVariantsFromList(og, "a,nope"); err == nil ||
		!strings.Contains(err.Error(), `"nope"`) {
		t.Errorf(`list "a,nope": want error naming "nope", got %v`, err)
	}
}
