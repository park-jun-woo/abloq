//ff:func feature=quest type=frame control=sequence topic=queue
//ff:what 게이트 FAIL 경로 검증 — 빈 diff(무작업)는 lastmod-advance가, 범위 밖 글 수정은 queue-scope가 root cause로 잡는지
package refresh

import (
	"strings"
	"testing"

	rgate "github.com/park-jun-woo/reins/pkg/gate"
	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestEvaluateFail(t *testing.T) {
	// Empty diff — no work done: lastmod-advance is the root cause.
	root := writeInstance(t)
	v := rgate.Evaluate(Definition{}.Rules(), subWith(t, root))
	if v.Outcome != quest.OutFail || v.RootCause != "lastmod-advance" {
		t.Fatalf("empty diff: outcome=%s root=%s — want FAIL/lastmod-advance", v.Outcome, v.RootCause)
	}
	// A proper refresh that ALSO touches another article: queue-scope fires.
	root2 := writeInstance(t)
	refreshed := strings.Replace(baseArticleMD, "lastmod: 2026-06-02", "lastmod: 2026-06-09", 1)
	refreshed = strings.Replace(refreshed,
		"This stale body sentence still describes the situation as of early 2025 in vendor terms.",
		"The refreshed body now reflects the mid 2026 landscape with current vendor guidance and revised context.", 1)
	writeFile(t, root2, "content/en/posts/a.md", refreshed)
	writeFile(t, root2, "content/en/posts/other.md", "---\ntitle: X\n---\nstray edit\n")
	v = rgate.Evaluate(Definition{}.Rules(), subWith(t, root2))
	if v.Outcome != quest.OutFail {
		t.Fatalf("out-of-scope edit: outcome=%s — want FAIL", v.Outcome)
	}
	found := false
	for _, f := range v.Facts {
		if f.Rule == "queue-scope" && strings.Contains(f.Actual, "content/en/posts/other.md") {
			found = true
		}
	}
	if !found {
		t.Errorf("queue-scope Fact missing: %+v", v.Facts)
	}
}
