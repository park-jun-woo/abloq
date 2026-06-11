//ff:func feature=sitesyaml type=rule control=sequence topic=gsc
//ff:what ruleGSCSiteURL이 형식 위반(스킴·빈 sc-domain)만 진단하고 빈 값·적법 형식 2종은 통과시키는지 검증
package sitesyaml

import "testing"

func TestRuleGSCSiteURL(t *testing.T) {
	src := []byte("sites:\n  - name: a\n    repo_path: /x\n  - name: b\n    repo_path: /x\n    gsc:\n      site_url: sc-domain:b.com\n  - name: c\n    repo_path: /x\n    gsc:\n      site_url: https://c.com/\n  - name: d\n    repo_path: /x\n    gsc:\n      site_url: ftp://d.com\n")
	s, idx, diags := Parse("sites.yaml", src)
	if len(diags) != 0 {
		t.Fatalf("parse: %v", diags)
	}
	got := ruleGSCSiteURL("sites.yaml", s, idx)
	if len(got) != 1 {
		t.Fatalf("want 1 diagnostic, got %v", got)
	}
	if got[0].Rule != "gsc-site-url" || got[0].Line != 15 {
		t.Errorf("diag = %+v, want gsc-site-url at the site_url line", got[0])
	}
}
