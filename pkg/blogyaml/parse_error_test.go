//ff:func feature=blogyaml type=parser control=iteration dimension=1 topic=diagnostics
//ff:what Parse가 깨진 YAML(syntax)과 빈 입력을 각각 yaml-syntax 진단으로 거부하는지 검증
package blogyaml

import "testing"

func TestParseError(t *testing.T) {
	if b, _, diags := Parse("blog.yaml", nil); b != nil || len(diags) != 1 {
		t.Fatalf("nil input: want nil Blog and 1 diagnostic, got %v %v", b, diags)
	}
	cases := []struct {
		name    string
		src     string
		wantMsg string
	}{
		{"syntax error", "\tlanguages: [ko]\n", ""},
		{"empty input", "", "blog.yaml is empty"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) { checkParseError(t, tc.src, tc.wantMsg) })
	}
}
