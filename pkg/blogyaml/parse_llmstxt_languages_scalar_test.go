//ff:func feature=blogyaml type=parser control=sequence
//ff:what geo.llms_txt.languages 2차 union 스칼라 파싱 검증 — all/base 스칼라가 1원소 리스트로 정규화되는지
package blogyaml

import (
	"reflect"
	"testing"
)

func TestParseLlmsTxtLanguagesScalar(t *testing.T) {
	src := []byte("languages: [ko]\nsections: [tech]\ngeo:\n  llms_txt:\n    languages: all\n")
	b, _, diags := Parse("blog.yaml", src)
	if len(diags) != 0 {
		t.Fatalf("want 0 diagnostics, got %v", diags)
	}
	if !reflect.DeepEqual(b.Geo.LlmsTxt.Languages, []string{"all"}) {
		t.Errorf("want languages [all], got %v", b.Geo.LlmsTxt.Languages)
	}
	if b.Geo.LlmsTxt.Mode != "auto" {
		t.Errorf("object form without mode must keep default auto, got %q", b.Geo.LlmsTxt.Mode)
	}
	base, _, diags := Parse("blog.yaml", []byte("languages: [ko]\nsections: [tech]\ngeo:\n  llms_txt:\n    languages: base\n"))
	if len(diags) != 0 {
		t.Fatalf("want 0 diagnostics for scalar base, got %v", diags)
	}
	if !reflect.DeepEqual(base.Geo.LlmsTxt.Languages, []string{"base"}) {
		t.Errorf("want languages [base], got %v", base.Geo.LlmsTxt.Languages)
	}
}
