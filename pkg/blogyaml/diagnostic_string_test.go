//ff:func feature=blogyaml type=schema control=sequence topic=diagnostics
//ff:what 진단 한 줄 포맷이 "파일:라인 [룰ID] 메시지"인지 검증
package blogyaml

import "testing"

func TestDiagnosticString(t *testing.T) {
	d := Diagnostic{File: "blog.yaml", Line: 7, Rule: "lang-bcp47", Message: "bad code"}
	want := "blog.yaml:7 [lang-bcp47] bad code"
	if got := d.String(); got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}
