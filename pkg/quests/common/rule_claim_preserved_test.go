//ff:func feature=quest type=rule control=sequence topic=queue
//ff:what claim-preserved가 주장 건수 감소를 FAIL로 잡고, 동수 교체·증가는 통과하는지 검증
package common

import (
	"strings"
	"testing"

	rgate "github.com/park-jun-woo/reins/pkg/gate"
)

func TestRuleClaimPreserved(t *testing.T) {
	r := RuleClaimPreserved()
	if r.Meta.ID != "claim-preserved" || r.Meta.Level != rgate.LevelFail {
		t.Fatalf("Meta = %+v", r.Meta)
	}
	// Replacing a stale figure keeps the count — passes.
	replaced := strings.Replace(claimBaseMD, "Throughput grew 40%", "Throughput grew 55%", 1)
	c := consFixture(t, claimBaseMD, replaced)
	if fired, _ := r.Check(rgate.Context{Submission: consCarrierStub{c}}); fired {
		t.Error("1:1 claim replacement fired")
	}
	// Deleting a claim drops the count — fails.
	deleted := strings.Replace(claimBaseMD, "Latency dropped 120ms in the same run.\n", "", 1)
	c = consFixture(t, claimBaseMD, deleted)
	fired, fact := r.Check(rgate.Context{Submission: consCarrierStub{c}})
	if !fired || !strings.Contains(fact.Expected, ">= 2") {
		t.Errorf("claim deletion: fired=%v fact=%+v", fired, fact)
	}
	// Base nil (defensive) is inert.
	c = consFixture(t, claimBaseMD, deleted)
	c.Target.Articles[0].Base = nil
	if fired, _ := r.Check(rgate.Context{Submission: consCarrierStub{c}}); fired {
		t.Error("Base nil must be inert")
	}
}
