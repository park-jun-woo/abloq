//ff:func feature=sitesyaml type=rule control=sequence
//ff:what ruleRepoPath가 repo_path 부재·상대경로를 거부하고 절대경로는 통과시키는지 검증
package sitesyaml

import "testing"

func TestRuleRepoPath(t *testing.T) {
	src := []byte("sites:\n  - name: a\n    repo_path: /blogs/a\n  - name: b\n  - name: c\n    repo_path: blogs/c\n")
	s, idx, diags := Parse("sites.yaml", src)
	if len(diags) != 0 {
		t.Fatalf("parse: %v", diags)
	}
	got := ruleRepoPath("sites.yaml", s, idx)
	if len(got) != 2 {
		t.Fatalf("want 2 diagnostics, got %v", got)
	}
	if got[0].Line != 4 || got[0].Message != "sites[1].repo_path is required" {
		t.Errorf("missing repo_path diag = %+v", got[0])
	}
	if got[1].Line != 6 || got[1].Rule != "repo-path" {
		t.Errorf("relative repo_path diag = %+v", got[1])
	}
}
