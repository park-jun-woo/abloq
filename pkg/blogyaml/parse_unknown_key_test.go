//ff:func feature=blogyaml type=parser control=sequence
//ff:what strict 파싱 검증 — 스키마에 없는 키(후속 Phase 키 포함)가 unknown-key 진단으로 거부되는지
package blogyaml

import "testing"

func TestParseUnknownKey(t *testing.T) {
	src := []byte("languages: [ko]\nsections: [tech]\nmin_meaningful_diff: 40\n")
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
	if diags[0].Line != 3 {
		t.Errorf("want line 3, got %d", diags[0].Line)
	}
}
