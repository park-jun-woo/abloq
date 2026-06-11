//ff:func feature=insight type=rule control=sequence
//ff:what Validate 통합 검증 — 에러 룰 3종 수집과 anchors 경고의 분리 반환
package insight

import "testing"

func TestValidate(t *testing.T) {
	ins := &Insight{Claims: []Claim{
		{ID: "a", Text: "x", Kind: "claim", Anchors: []string{"k"}},
		{ID: "a", Text: "y", Kind: "bogus"},
	}}
	errs, warns := Validate("insight.yaml", ins, []int{3, 7})
	if len(errs) != 2 {
		t.Errorf("want 2 errors (id-unique + kind), got %v", errs)
	}
	if len(warns) != 1 || warns[0].Rule != "insight-claim-anchors-empty" {
		t.Errorf("want 1 anchors-empty warning, got %v", warns)
	}
	errs, warns = Validate("insight.yaml", &Insight{}, nil)
	if len(errs) != 1 || errs[0].Rule != "insight-claims-min" || len(warns) != 0 {
		t.Errorf("want only claims-min error for empty insight, got errs=%v warns=%v", errs, warns)
	}
}
