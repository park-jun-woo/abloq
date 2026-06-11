//ff:func feature=insight type=rule control=sequence
//ff:what insight-claims-min 검증 — claims 0개는 에러, 1개 이상은 통과
package insight

import "testing"

func TestRuleClaimsMin(t *testing.T) {
	diags := ruleClaimsMin("insight.yaml", &Insight{})
	if len(diags) != 1 || diags[0].Rule != "insight-claims-min" {
		t.Errorf("want insight-claims-min diagnostic for zero claims, got %v", diags)
	}
	if diags := ruleClaimsMin("insight.yaml", &Insight{Claims: []Claim{{ID: "a"}}}); len(diags) != 0 {
		t.Errorf("want no diagnostic with 1 claim, got %v", diags)
	}
}
