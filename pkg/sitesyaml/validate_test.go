//ff:func feature=sitesyaml type=rule control=iteration dimension=1
//ff:what Validate가 룰 5종의 진단을 모두 수집하고 적법 선언에는 0건인지 검증
package sitesyaml

import "testing"

func TestValidate(t *testing.T) {
	src := []byte("sites:\n  - name: a\n    repo_path: /blogs/a\n")
	s, idx, diags := Parse("sites.yaml", src)
	if len(diags) != 0 {
		t.Fatalf("parse: %v", diags)
	}
	if got := Validate("sites.yaml", s, idx); len(got) != 0 {
		t.Errorf("valid declaration must have 0 diagnostics, got %v", got)
	}

	src = []byte("sites:\n  - name: Bad Name\n    repo_path: relative\n  - name: a\n    repo_path: /x\n  - name: a\n    repo_path: /x\n    gsc:\n      site_url: ftp://nope\n")
	s, idx, diags = Parse("sites.yaml", src)
	if len(diags) != 0 {
		t.Fatalf("parse: %v", diags)
	}
	got := Validate("sites.yaml", s, idx)
	rules := map[string]bool{}
	for _, d := range got {
		rules[d.Rule] = true
	}
	for _, rule := range []string{"name-format", "name-unique", "repo-path", "gsc-site-url"} {
		if !rules[rule] {
			t.Errorf("missing %s diagnostic in %v", rule, got)
		}
	}

	empty := &Sites{}
	got = Validate("sites.yaml", empty, lineIndex{})
	if len(got) != 1 || got[0].Rule != "sites-empty" {
		t.Errorf("empty list = %v, want one sites-empty diagnostic", got)
	}
}
