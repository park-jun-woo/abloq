//ff:func feature=blogyaml type=parser control=sequence
//ff:what geo.llms_txt 문자열 단축형 파싱 검증 — manual이 mode로 들어가고 나머지 필드는 기본값 유지(auto 하위호환 포함)
package blogyaml

import (
	"reflect"
	"testing"
)

func TestParseLlmsTxtScalar(t *testing.T) {
	src := []byte("languages: [ko]\nsections: [tech]\ngeo:\n  llms_txt: manual\n")
	b, _, diags := Parse("blog.yaml", src)
	if len(diags) != 0 {
		t.Fatalf("want 0 diagnostics, got %v", diags)
	}
	if b.Geo.LlmsTxt.Mode != "manual" {
		t.Errorf("want mode manual, got %q", b.Geo.LlmsTxt.Mode)
	}
	if b.Geo.LlmsTxtMode() != "manual" {
		t.Errorf("want LlmsTxtMode manual, got %q", b.Geo.LlmsTxtMode())
	}
	if !reflect.DeepEqual(b.Geo.LlmsTxt.Languages, []string{"base"}) {
		t.Errorf("shorthand must keep default languages [base], got %v", b.Geo.LlmsTxt.Languages)
	}
	auto, _, diags := Parse("blog.yaml", []byte("languages: [ko]\nsections: [tech]\ngeo:\n  llms_txt: auto\n"))
	if len(diags) != 0 {
		t.Fatalf("legacy 'llms_txt: auto' must still parse, got %v", diags)
	}
	if auto.Geo.LlmsTxt.Mode != "auto" {
		t.Errorf("want legacy auto mode, got %q", auto.Geo.LlmsTxt.Mode)
	}
}
