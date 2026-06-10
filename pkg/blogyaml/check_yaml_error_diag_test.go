//ff:func feature=blogyaml type=parser control=sequence topic=diagnostics
//ff:what yamlErrorDiag 케이스 하나를 실행해 파일·룰ID·라인·메시지를 검증
package blogyaml

import "testing"

func checkYamlErrorDiag(t *testing.T, msg, wantRule, wantMsg string, wantLine int) {
	t.Helper()
	d := yamlErrorDiag("blog.yaml", msg)
	if d.File != "blog.yaml" {
		t.Errorf("want file blog.yaml, got %q", d.File)
	}
	if d.Rule != wantRule {
		t.Errorf("want rule %q, got %q", wantRule, d.Rule)
	}
	if d.Line != wantLine {
		t.Errorf("want line %d, got %d", wantLine, d.Line)
	}
	if d.Message != wantMsg {
		t.Errorf("want message %q, got %q", wantMsg, d.Message)
	}
}
