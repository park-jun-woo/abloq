//ff:func feature=archive type=client control=iteration dimension=1
//ff:what planItems가 URL × kind 3종을 pending 항목으로 전개하고 deploy_id·빈 request/response를 채우는지 검증
package archive

import "testing"

func TestPlanItems(t *testing.T) {
	items := planItems("dep-1", []string{"https://a/", "https://b/"})
	if len(items) != 6 {
		t.Fatalf("len = %d, want 6 (2 URLs × 3 kinds)", len(items))
	}
	kinds := map[string]int{}
	for _, item := range items {
		kinds[item.Kind]++
		if item.DeployID != "dep-1" || item.Status != StatusPending {
			t.Errorf("item %+v: want deploy_id dep-1 and status pending", item)
		}
		if string(item.Request) != "{}" || string(item.Response) != "{}" {
			t.Errorf("item %+v: pending request/response must be {}", item)
		}
	}
	if kinds[KindWayback] != 2 || kinds[KindIndexNow] != 2 || kinds[KindGSCIndex] != 2 {
		t.Errorf("kind fan-out = %v, want 2 each", kinds)
	}
	if len(planItems("dep-1", nil)) != 0 {
		t.Error("no changed URLs must plan nothing")
	}
}
