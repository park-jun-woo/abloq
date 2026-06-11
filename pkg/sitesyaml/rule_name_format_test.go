//ff:func feature=sitesyaml type=rule control=iteration dimension=1
//ff:what ruleNameFormat이 name 부재와 비슬러그(대문자·공백·언더스코어·하이픈 가장자리)를 거부하고 적법 슬러그는 통과시키는지 검증
package sitesyaml

import "testing"

func TestRuleNameFormat(t *testing.T) {
	ok := &Sites{Sites: []Site{{Name: "a"}, {Name: "park-jun-woo"}, {Name: "site2"}}}
	if got := ruleNameFormat("sites.yaml", ok, lineIndex{}); len(got) != 0 {
		t.Errorf("valid slugs = %v, want 0 diagnostics", got)
	}

	src := []byte("sites:\n  - repo_path: /x\n  - name: Upper\n  - name: has space\n  - name: under_score\n  - name: -edge\n  - name: edge-\n")
	s, idx, diags := Parse("sites.yaml", src)
	if len(diags) != 0 {
		t.Fatalf("parse: %v", diags)
	}
	got := ruleNameFormat("sites.yaml", s, idx)
	if len(got) != 6 {
		t.Fatalf("want 6 diagnostics, got %d: %v", len(got), got)
	}
	if got[0].Line != 2 || got[0].Message != "sites[0].name is required" {
		t.Errorf("missing name diag = %+v, want required at the item line", got[0])
	}
	for _, d := range got[1:] {
		if d.Rule != "name-format" {
			t.Errorf("diag = %+v, want name-format", d)
		}
	}
}
