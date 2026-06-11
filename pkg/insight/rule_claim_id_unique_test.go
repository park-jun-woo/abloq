//ff:func feature=insight type=rule control=sequence
//ff:what insight-claim-id-unique 검증 — 중복 id는 라인 포함 에러, 유니크 id는 통과
package insight

import "testing"

func TestRuleClaimIDUnique(t *testing.T) {
	ins := &Insight{Claims: []Claim{{ID: "a"}, {ID: "b"}, {ID: "a"}}}
	diags := ruleClaimIDUnique("insight.yaml", ins, []int{3, 5, 7})
	if len(diags) != 1 || diags[0].Rule != "insight-claim-id-unique" || diags[0].Line != 7 {
		t.Errorf("want one duplicate-id diagnostic at line 7, got %v", diags)
	}
	if diags := ruleClaimIDUnique("insight.yaml", &Insight{Claims: []Claim{{ID: "a"}, {ID: "b"}}}, nil); len(diags) != 0 {
		t.Errorf("want no diagnostic for unique ids, got %v", diags)
	}
}
