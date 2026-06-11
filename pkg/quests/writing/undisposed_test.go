//ff:func feature=quest type=rule control=sequence
//ff:what undisposed 검증 — disposition 라인이 없는 미출현 claim ID만 반환, 전건 커버리지는 빈 목록
package writing

import (
	"testing"

	"github.com/park-jun-woo/abloq/pkg/insight"
)

func TestUndisposed(t *testing.T) {
	missing := []insight.Claim{{ID: "c1"}, {ID: "c2"}}
	review := "reviewer: x\n- c1: addressed — fine\n"
	got := undisposed(review, missing)
	if len(got) != 1 || got[0] != "c2" {
		t.Errorf("got %v, want [c2]", got)
	}
	review += "- c2: revised — fixed\n"
	if got := undisposed(review, missing); len(got) != 0 {
		t.Errorf("full coverage: got %v, want empty", got)
	}
}
