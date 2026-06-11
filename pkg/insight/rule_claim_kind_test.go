//ff:func feature=insight type=rule control=iteration dimension=1
//ff:what insight-claim-kind 검증 — 사전 4종은 통과, 사전 밖·빈 kind는 에러
package insight

import "testing"

func TestRuleClaimKind(t *testing.T) {
	for _, kind := range []string{"claim", "rebuttal", "prediction", "definition"} {
		ins := &Insight{Claims: []Claim{{ID: "a", Kind: kind}}}
		if diags := ruleClaimKind("insight.yaml", ins, nil); len(diags) != 0 {
			t.Errorf("want kind %q accepted, got %v", kind, diags)
		}
	}
	ins := &Insight{Claims: []Claim{{ID: "a", Kind: "opinion"}, {ID: "b"}}}
	diags := ruleClaimKind("insight.yaml", ins, []int{3, 6})
	if len(diags) != 2 || diags[0].Rule != "insight-claim-kind" || diags[1].Line != 6 {
		t.Errorf("want 2 kind diagnostics (unknown + empty), got %v", diags)
	}
}
