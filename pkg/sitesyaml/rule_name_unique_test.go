//ff:func feature=sitesyaml type=rule control=sequence
//ff:what ruleNameUnique가 중복 name을 두 번째 항목부터 진단하고 빈 name은 스킵(name-format 소관)하는지 검증
package sitesyaml

import "testing"

func TestRuleNameUnique(t *testing.T) {
	src := []byte("sites:\n  - name: a\n  - name: b\n  - name: a\n")
	s, idx, diags := Parse("sites.yaml", src)
	if len(diags) != 0 {
		t.Fatalf("parse: %v", diags)
	}
	got := ruleNameUnique("sites.yaml", s, idx)
	if len(got) != 1 || got[0].Rule != "name-unique" || got[0].Line != 4 {
		t.Fatalf("duplicate = %v, want one name-unique diagnostic at line 4", got)
	}

	empties := &Sites{Sites: []Site{{Name: ""}, {Name: ""}}}
	if got := ruleNameUnique("sites.yaml", empties, lineIndex{}); len(got) != 0 {
		t.Errorf("empty names = %v, want 0 (name-format reports them)", got)
	}
}
