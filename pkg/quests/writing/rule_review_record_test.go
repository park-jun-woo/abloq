//ff:func feature=quest type=rule control=sequence
//ff:what review-record 룰 검증 — 기록 부재/reviewer 없음/disposition 누락은 발동, 전건 커버리지면 통과
package writing

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/insight"
)

func TestRuleReviewRecord(t *testing.T) {
	r := ruleReviewRecord()
	missing := []insight.Claim{{ID: "c1"}, {ID: "c2"}}
	sub := &Submission{Missing: missing, ReviewPath: "quests/writing/reviews/a.md"}
	fired, fact := fireRule(t, r, sub)
	if !fired || !strings.Contains(fact.Actual, "missing") {
		t.Errorf("absent record: fired=%v fact=%+v", fired, fact)
	}
	sub.Review = "- c1: addressed — ok\n- c2: excluded — scope\n"
	fired, fact = fireRule(t, r, sub)
	if !fired || !strings.Contains(fact.Actual, "reviewer") {
		t.Errorf("no reviewer: fired=%v fact=%+v", fired, fact)
	}
	sub.Review = "reviewer: ctx-2\n- c1: addressed — ok\n"
	fired, fact = fireRule(t, r, sub)
	if !fired || !strings.Contains(fact.Actual, "c2") {
		t.Errorf("undisposed c2: fired=%v fact=%+v", fired, fact)
	}
	sub.Review = "reviewer: ctx-2\n- c1: addressed — ok\n- c2: excluded — scope\n"
	if fired, _ = fireRule(t, r, sub); fired {
		t.Error("full coverage: rule fired")
	}
}
