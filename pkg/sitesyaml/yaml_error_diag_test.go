//ff:func feature=sitesyaml type=parser control=sequence topic=diagnostics
//ff:what yamlErrorDiag가 라인 번호 추출·unknown-key 분류·"yaml: " 접두 제거를 수행하는지 검증
package sitesyaml

import "testing"

func TestYamlErrorDiag(t *testing.T) {
	d := yamlErrorDiag("sites.yaml", "yaml: line 7: could not find expected ':'")
	if d.Line != 7 || d.Rule != "yaml-syntax" {
		t.Errorf("syntax diag = %+v, want line 7 yaml-syntax", d)
	}
	if d.Message != "line 7: could not find expected ':'" {
		t.Errorf("message must drop the yaml: prefix, got %q", d.Message)
	}

	d = yamlErrorDiag("sites.yaml", "yaml: unmarshal errors:\n  line 3: field nam not found in type sitesyaml.Site")
	if d.Line != 3 || d.Rule != "unknown-key" {
		t.Errorf("unknown-key diag = %+v, want line 3 unknown-key", d)
	}

	d = yamlErrorDiag("sites.yaml", "yaml: some error without a line")
	if d.Line != 1 {
		t.Errorf("no line info must fall back to 1, got %d", d.Line)
	}
}
