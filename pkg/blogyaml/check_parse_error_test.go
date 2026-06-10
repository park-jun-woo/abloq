//ff:func feature=blogyaml type=parser control=sequence topic=diagnostics
//ff:what Parse 에러 케이스 하나를 실행해 nil Blog와 yaml-syntax 진단 1건(필요시 메시지)을 검증
package blogyaml

import "testing"

func checkParseError(t *testing.T, src, wantMsg string) {
	t.Helper()
	b, _, diags := Parse("blog.yaml", []byte(src))
	if b != nil {
		t.Fatalf("want nil Blog, got %+v", b)
	}
	if len(diags) != 1 {
		t.Fatalf("want 1 diagnostic, got %d: %v", len(diags), diags)
	}
	if diags[0].Rule != "yaml-syntax" {
		t.Errorf("want rule yaml-syntax, got %q", diags[0].Rule)
	}
	if wantMsg != "" && diags[0].Message != wantMsg {
		t.Errorf("want message %q, got %q", wantMsg, diags[0].Message)
	}
}
