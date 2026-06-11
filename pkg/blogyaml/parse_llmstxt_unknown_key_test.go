//ff:func feature=blogyaml type=parser control=sequence
//ff:what geo.llms_txt 객체 폼의 미지 키가 KnownFields와 동일하게 unknown-key 진단(룰·라인)으로 거부되는지 검증
package blogyaml

import (
	"strings"
	"testing"
)

func TestParseLlmsTxtUnknownKey(t *testing.T) {
	src := []byte("languages: [ko]\nsections: [tech]\ngeo:\n  llms_txt:\n    mode: auto\n    foo: 1\n")
	b, _, diags := Parse("blog.yaml", src)
	if b != nil {
		t.Fatalf("want nil Blog on unknown key, got %+v", b)
	}
	if len(diags) != 1 {
		t.Fatalf("want 1 diagnostic, got %d: %v", len(diags), diags)
	}
	if diags[0].Rule != "unknown-key" {
		t.Errorf("want rule unknown-key, got %q", diags[0].Rule)
	}
	if diags[0].Line != 6 {
		t.Errorf("want line 6, got %d", diags[0].Line)
	}
	if !strings.Contains(diags[0].Message, "field foo not found in type blogyaml.LlmsTxtSpec") {
		t.Errorf("want KnownFields-shaped message, got %q", diags[0].Message)
	}
}
