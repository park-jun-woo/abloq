//ff:func feature=insight type=rule control=sequence
//ff:what insight-claim-anchors-empty 검증 — anchors 빈 claim은 requires_source 무관 경고, 있으면 통과
package insight

import "testing"

func TestRuleAnchorsEmpty(t *testing.T) {
	ins := &Insight{Claims: []Claim{
		{ID: "a", Anchors: []string{"k"}},
		{ID: "b", RequiresSource: true},
	}}
	diags := ruleAnchorsEmpty("insight.yaml", ins, []int{3, 5})
	if len(diags) != 1 || diags[0].Rule != "insight-claim-anchors-empty" || diags[0].Line != 5 {
		t.Errorf("want one anchors-empty warning at line 5 (requires_source irrelevant), got %v", diags)
	}
	if diags := ruleAnchorsEmpty("insight.yaml", &Insight{Claims: []Claim{{ID: "a", Anchors: []string{"k"}}}}, nil); len(diags) != 0 {
		t.Errorf("want no warning when anchors present, got %v", diags)
	}
}
