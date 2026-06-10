//ff:func feature=blogyaml type=rule control=iteration dimension=1
//ff:what ruleLangBCP47 케이스 하나를 실행해 진단 수와 룰ID를 검증
package blogyaml

import "testing"

func checkRuleLangBCP47(t *testing.T, languages []string, wantDiags int) {
	t.Helper()
	b := &Blog{Languages: languages}
	diags := ruleLangBCP47("blog.yaml", b, lineIndex{"languages": 3, "languages[1]": 5})
	if len(diags) != wantDiags {
		t.Fatalf("want %d diagnostics, got %d: %v", wantDiags, len(diags), diags)
	}
	for _, d := range diags {
		if d.Rule != "lang-bcp47" {
			t.Errorf("want rule lang-bcp47, got %q", d.Rule)
		}
	}
}
