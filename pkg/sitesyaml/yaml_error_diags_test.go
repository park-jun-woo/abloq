//ff:func feature=sitesyaml type=parser control=sequence topic=diagnostics
//ff:what yamlErrorDiags가 단일 에러는 1건, yaml.TypeError는 메시지별 다건 진단으로 펼치는지 검증
package sitesyaml

import (
	"errors"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestYamlErrorDiags(t *testing.T) {
	diags := yamlErrorDiags("sites.yaml", errors.New("yaml: line 2: bad indent"))
	if len(diags) != 1 || diags[0].Line != 2 {
		t.Errorf("single error = %v, want one diag at line 2", diags)
	}

	te := &yaml.TypeError{Errors: []string{
		"line 3: field nam not found in type sitesyaml.Site",
		"line 5: field repo not found in type sitesyaml.Site",
	}}
	diags = yamlErrorDiags("sites.yaml", te)
	if len(diags) != 2 || diags[0].Line != 3 || diags[1].Line != 5 {
		t.Errorf("type error = %v, want two diags at lines 3 and 5", diags)
	}
	if diags[0].Rule != "unknown-key" || diags[1].Rule != "unknown-key" {
		t.Errorf("type error diags = %v, want unknown-key rule", diags)
	}
}
