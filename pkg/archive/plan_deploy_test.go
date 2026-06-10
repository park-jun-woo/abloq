//ff:func feature=archive type=client control=iteration dimension=1
//ff:what PlanDeploy가 이미 영수증 있는 (kind, target)만 건너뛰고 나머지를 pending으로 계획하는지(멱등 재웹훅 = 0건) 검증
package archive

import "testing"

func TestPlanDeploy(t *testing.T) {
	changed := []string{"https://a/", "https://b/"}
	existing := []Existing{
		{Kind: KindWayback, Target: "https://a/"},
		{Kind: KindIndexNow, Target: "https://a/"},
	}
	planned := PlanDeploy("dep-1", changed, existing)
	if len(planned) != 4 {
		t.Fatalf("len = %d, want 4 (6 pairs − 2 existing)", len(planned))
	}
	pairs := make([]Existing, 0, 6)
	for _, item := range planned {
		if item.Target == "https://a/" && item.Kind != KindGSCIndex {
			t.Errorf("pair (%s, %s) was already receipted and must be skipped", item.Kind, item.Target)
		}
		pairs = append(pairs, Existing{Kind: item.Kind, Target: item.Target})
	}
	pairs = append(pairs, existing...)

	rerun := PlanDeploy("dep-1", changed, pairs)
	if len(rerun) != 0 {
		t.Errorf("idempotent re-webhook planned %d items, want 0", len(rerun))
	}
}
