//ff:func feature=gen type=generator control=sequence
//ff:what jsonld.json 렌더가 멱등이고 타입 목록·저자 Person 엔티티·발행자를 고정 키 순서로 내는지 검증
package jsonld

import (
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func TestRender(t *testing.T) {
	b := &blogyaml.Blog{
		Site: blogyaml.Site{BaseURL: "https://x.com", Title: "X Blog", Author: "A"},
		Geo:  blogyaml.Geo{JSONLD: []string{"Article", "Person"}},
	}
	out := Render(b)
	want := "{\n" +
		"  \"types\": [\n    \"Article\",\n    \"Person\"\n  ],\n" +
		"  \"author\": {\n    \"@type\": \"Person\",\n    \"name\": \"A\",\n    \"url\": \"https://x.com\"\n  },\n" +
		"  \"publisher\": \"X Blog\"\n" +
		"}\n"
	if string(out) != want {
		t.Errorf("Render = %q, want %q", out, want)
	}
	if again := Render(b); string(again) != string(out) {
		t.Errorf("Render is not idempotent: %q vs %q", out, again)
	}
}
