//ff:func feature=blogyaml type=parser control=iteration dimension=1 topic=diagnostics
//ff:what yamlErrorDiag가 라인 번호 추출, unknown-key/yaml-syntax 분류, "yaml: " 접두사 제거를 수행하는지 검증
package blogyaml

import "testing"

func TestYamlErrorDiag(t *testing.T) {
	cases := []struct {
		name, msg, wantRule, wantMsg string
		wantLine                     int
	}{
		{
			name: "unknown key with line", msg: "yaml: unmarshal errors:\n  line 3: field bogus not found in type blogyaml.Blog",
			wantRule: "unknown-key", wantLine: 3,
			wantMsg: "unmarshal errors:\n  line 3: field bogus not found in type blogyaml.Blog",
		},
		{
			name: "syntax error with line", msg: "yaml: line 7: did not find expected key",
			wantRule: "yaml-syntax", wantLine: 7,
			wantMsg: "line 7: did not find expected key",
		},
		{
			name: "no line falls back to 1", msg: "yaml: mapping values are not allowed in this context",
			wantRule: "yaml-syntax", wantLine: 1,
			wantMsg: "mapping values are not allowed in this context",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) { checkYamlErrorDiag(t, tc.msg, tc.wantRule, tc.wantMsg, tc.wantLine) })
	}
}
