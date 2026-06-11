//ff:func feature=quest type=rule control=sequence topic=queue
//ff:what [lastmod-advance] 제출본 lastmod가 기준선(git HEAD) lastmod보다 미래인지 검증 — 갱신 퀘스트의 작업 완수 강제 룰
//ff:why honest-lastmod 단독은 "변경 시의 정직성"만 보고 변경을 강제하지 않는다 — 빈 diff PASS→큐 삭제→재발급 무한 루프를 lastmod 전진 강제가 끊고, 전진하면 honest-lastmod가 의미 diff(≥min_meaningful_diff)+큐 등재를 연쇄 강제한다 (2차 검수 D1)
package refresh

import (
	rgate "github.com/park-jun-woo/reins/pkg/gate"
	"github.com/park-jun-woo/reins/pkg/quest"
)

// ruleLastmodAdvance builds the refresh work-completion rule: the submitted
// article's lastmod must parse and be strictly later than the baseline's.
// A missing baseline lastmod compares as the zero time, so any valid
// submitted lastmod advances past it.
func ruleLastmodAdvance() rgate.Rule {
	return rgate.Rule{
		Meta: rgate.RuleMeta{ID: "lastmod-advance", Level: rgate.LevelFail,
			Desc: "front matter lastmod advances strictly past the baseline (forces real refresh work)"},
		Check: func(ctx rgate.Context) (bool, quest.Fact) {
			sub := ctx.Submission.(*Submission)
			a := sub.Target.Articles[0]
			docT, ok := fmTime(a.Doc, "lastmod")
			if !ok {
				return true, quest.Fact{Where: sub.Article + "#lastmod",
					Expected: "a parseable lastmod later than the baseline's",
					Actual:   "lastmod missing or unparseable"}
			}
			baseT, _ := fmTime(a.Base, "lastmod")
			if docT.After(baseT) {
				return false, quest.Fact{}
			}
			return true, quest.Fact{Where: sub.Article + "#lastmod",
				Expected: "lastmod > " + baseT.Format("2006-01-02") + " (baseline)",
				Actual:   "lastmod " + docT.Format("2006-01-02") + " did not advance"}
		},
	}
}
