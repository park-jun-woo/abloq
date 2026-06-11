//ff:func feature=quest type=rule control=sequence topic=queue
//ff:what claim-scope가 큐 밖 주장 변형을 FAIL로 잡고, 큐에 든 주장의 변경·무변경 제출은 통과하는지 검증
package common

import (
	"strings"
	"testing"

	rgate "github.com/park-jun-woo/reins/pkg/gate"

	agate "github.com/park-jun-woo/abloq/pkg/gate"
)

const claimBaseMD = `---
title: "A"
date: 2026-06-01
lastmod: 2026-06-02
---

Throughput grew 40% after the change.

Latency dropped 120ms in the same run.

## Sources

- [Spec](https://example.org/spec)
`

func TestRuleClaimScope(t *testing.T) {
	r := RuleClaimScope()
	if r.Meta.ID != "claim-scope" || r.Meta.Level != rgate.LevelFail {
		t.Fatalf("Meta = %+v", r.Meta)
	}
	// Unchanged claims pass.
	c := consFixture(t, claimBaseMD, claimBaseMD)
	if fired, _ := r.Check(rgate.Context{Submission: consCarrierStub{c}}); fired {
		t.Error("verbatim claims fired")
	}
	// Altering a claim that is NOT in the queue payload fails.
	tampered := strings.Replace(claimBaseMD, "Latency dropped 120ms", "Latency dropped 980ms", 1)
	c = consFixture(t, claimBaseMD, tampered)
	fired, fact := r.Check(rgate.Context{Submission: consCarrierStub{c}})
	if !fired || !strings.Contains(fact.Actual, "120ms") {
		t.Errorf("out-of-queue claim tampering: fired=%v fact=%+v", fired, fact)
	}
	// The same change passes once the queue payload authorizes that claim.
	c = consFixture(t, claimBaseMD, tampered)
	c.QueuedClaims[agate.HashText("Latency dropped 120ms in the same run.")] = true
	if fired, fact := r.Check(rgate.Context{Submission: consCarrierStub{c}}); fired {
		t.Errorf("queued claim change fired: %+v", fact)
	}
	// Base nil (defensive) is inert.
	c = consFixture(t, claimBaseMD, claimBaseMD)
	c.Target.Articles[0].Base = nil
	if fired, _ := r.Check(rgate.Context{Submission: consCarrierStub{c}}); fired {
		t.Error("Base nil must be inert")
	}
}
