//ff:func feature=blogyaml type=parser control=sequence
//ff:what pinned 엔트리의 미지 키가 unknown-key 진단(룰·라인)으로 거부되는지 검증 — 중첩 strict 디코드
package blogyaml

import (
	"strings"
	"testing"
)

func TestParseLlmsTxtPinnedUnknownKey(t *testing.T) {
	src := []byte("languages: [ko]\nsections: [tech]\ngeo:\n  llms_txt:\n    pinned:\n      - title: T\n        url: /x\n        link: y\n")
	b, _, diags := Parse("blog.yaml", src)
	if b != nil {
		t.Fatalf("want nil Blog on unknown pinned key, got %+v", b)
	}
	if len(diags) != 1 {
		t.Fatalf("want 1 diagnostic, got %d: %v", len(diags), diags)
	}
	if diags[0].Rule != "unknown-key" {
		t.Errorf("want rule unknown-key, got %q", diags[0].Rule)
	}
	if diags[0].Line != 8 {
		t.Errorf("want line 8, got %d", diags[0].Line)
	}
	if !strings.Contains(diags[0].Message, "field link not found in type blogyaml.LlmsPinned") {
		t.Errorf("want KnownFields-shaped message, got %q", diags[0].Message)
	}
}
