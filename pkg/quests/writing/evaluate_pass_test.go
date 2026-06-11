//ff:func feature=quest type=frame control=sequence
//ff:what 게이트 전 룰 통과 검증 — 클린 글 + 전건 disposition REVIEW 기록을 reins 레벨집계로 평가해 PASS
package writing

import (
	"testing"

	rgate "github.com/park-jun-woo/reins/pkg/gate"
	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestEvaluatePass(t *testing.T) {
	root := writeInstance(t)
	art, _ := passFixtures()
	sub := subWith(t, root, art, "")
	sub.Review = "reviewer: ctx-2\nall claims present, no non_goals drift\n"
	sub.ReviewPath = "quests/writing/reviews/fixture.md"
	v := rgate.Evaluate(Definition{}.Rules(), rgate.Context{Submission: sub})
	if v.Outcome != quest.OutPass {
		t.Fatalf("Outcome = %s, facts = %+v — want PASS", v.Outcome, v.Facts)
	}
}
