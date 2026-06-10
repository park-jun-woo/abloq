//ff:func feature=blogyaml type=parser control=iteration dimension=1 topic=diagnostics
//ff:what yamlErrorDiags 케이스 하나를 실행해 진단 수와 각 진단의 파일·라인 하한을 검증
package blogyaml

import "testing"

func checkYamlErrorDiags(t *testing.T, err error, wantDiags int) {
	t.Helper()
	diags := yamlErrorDiags("blog.yaml", err)
	if len(diags) != wantDiags {
		t.Fatalf("want %d diagnostics, got %d: %v", wantDiags, len(diags), diags)
	}
	for _, d := range diags {
		if d.File != "blog.yaml" {
			t.Errorf("want file blog.yaml, got %q", d.File)
		}
		if d.Line < 1 {
			t.Errorf("want line >= 1, got %d", d.Line)
		}
	}
}
