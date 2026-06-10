//ff:func feature=blogyaml type=rule control=iteration dimension=1
//ff:what ruleHeadingDefaultLang 케이스 하나를 실행해 진단 수와 룰ID를 검증
package blogyaml

import "testing"

func checkRuleHeadingDefaultLang(t *testing.T, languages []string, headings map[string]map[string]string, wantDiags int) {
	t.Helper()
	b := &Blog{Languages: languages, Structure: Structure{Headings: headings}}
	diags := ruleHeadingDefaultLang("blog.yaml", b, lineIndex{})
	if len(diags) != wantDiags {
		t.Fatalf("want %d diagnostics, got %d: %v", wantDiags, len(diags), diags)
	}
	for _, d := range diags {
		if d.Rule != "heading-default-lang" {
			t.Errorf("want rule heading-default-lang, got %q", d.Rule)
		}
	}
}
