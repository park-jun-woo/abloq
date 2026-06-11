//ff:func feature=sitesyaml type=schema control=sequence topic=diagnostics
//ff:what Diagnostic.String이 "파일:라인 [룰ID] 메시지" 형식으로 포맷하는지 검증
package sitesyaml

import "testing"

func TestDiagnosticString(t *testing.T) {
	d := Diagnostic{File: "sites.yaml", Line: 7, Rule: "name-format", Message: "sites[0].name is required"}
	want := "sites.yaml:7 [name-format] sites[0].name is required"
	if got := d.String(); got != want {
		t.Errorf("String() = %q, want %q", got, want)
	}
}
