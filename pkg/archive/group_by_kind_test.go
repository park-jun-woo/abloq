//ff:func feature=archive type=client control=sequence
//ff:what groupByKind가 kind별로 나누고 그룹 안의 입력 순서를 보존하는지 검증
package archive

import "testing"

func TestGroupByKind(t *testing.T) {
	groups := groupByKind([]Pending{
		{Kind: KindWayback, Target: "https://a/"},
		{Kind: KindGSCIndex, Target: "https://a/"},
		{Kind: KindWayback, Target: "https://b/"},
	})
	if len(groups[KindWayback]) != 2 || len(groups[KindGSCIndex]) != 1 || len(groups[KindIndexNow]) != 0 {
		t.Fatalf("group sizes = %d/%d/%d, want 2/1/0",
			len(groups[KindWayback]), len(groups[KindGSCIndex]), len(groups[KindIndexNow]))
	}
	if groups[KindWayback][0].Target != "https://a/" || groups[KindWayback][1].Target != "https://b/" {
		t.Errorf("wayback group order = %v, want input order preserved", groups[KindWayback])
	}
}
