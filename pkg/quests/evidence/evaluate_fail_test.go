//ff:func feature=quest type=frame control=sequence topic=queue
//ff:what 게이트 FAIL 경로 검증 — 빈 diff(무작업)는 claims-resolved가, 큐 밖 주장 변형은 claim-scope가 잡는지
package evidence

import (
	"strings"
	"testing"

	rgate "github.com/park-jun-woo/reins/pkg/gate"
	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestEvaluateFail(t *testing.T) {
	// Empty diff — no work done: the resolution re-check is the root cause.
	root := writeInstance(t)
	v := rgate.Evaluate(Definition{}.Rules(), subWith(t, root))
	if v.Outcome != quest.OutFail || v.RootCause != "claims-resolved" {
		t.Fatalf("empty diff: outcome=%s root=%s — want FAIL/claims-resolved", v.Outcome, v.RootCause)
	}
	// Proper sourcing that ALSO rewords an out-of-queue claim: claim-scope fires.
	root2 := writeInstance(t)
	tampered := strings.Replace(baseArticleMD,
		unsourcedClaim,
		unsourcedClaim+" [Migration report](https://example.org/spec)", 1)
	tampered = strings.Replace(tampered, rotURL, "https://example.org/live-study", 1)
	tampered = strings.Replace(tampered, "Latency dropped 120ms", "Latency dropped 980ms", 1)
	writeFile(t, root2, "content/en/posts/a.md", tampered)
	v = rgate.Evaluate(Definition{}.Rules(), subWith(t, root2))
	if v.Outcome != quest.OutFail {
		t.Fatalf("out-of-queue claim tampering: outcome=%s — want FAIL", v.Outcome)
	}
	found := false
	for _, f := range v.Facts {
		if f.Rule == "claim-scope" && strings.Contains(f.Actual, "120ms") {
			found = true
		}
	}
	if !found {
		t.Errorf("claim-scope Fact missing: %+v", v.Facts)
	}
}
