//ff:func feature=quest type=rule control=sequence
//ff:what [review-record] REVIEW 기록 존재 + reviewer 식별자 + match 미출현 claim 전건 disposition 커버리지의 결정적 검사
//ff:why 지지(출처·본문이 주장을 실제로 뒷받침하는가) 판정은 비결정이라 룰이 될 수 없다 — 게이트는 검토 기록의 존재와 커버리지만 잠그고, 판정 내용은 별도 컨텍스트 검토자(자가 REVIEW 금지 — context.md 규약)의 몫 (Phase016)
package writing

import (
	"strings"

	rgate "github.com/park-jun-woo/reins/pkg/gate"
	"github.com/park-jun-woo/reins/pkg/quest"
)

// ruleReviewRecord builds the review-record rule: the record must exist,
// carry a `reviewer:` context identifier, and dispose of every insight-match
// missing claim (addressed|revised|excluded). Coverage is deterministic; the
// judgment inside the record is the reviewer's, never the machine's.
func ruleReviewRecord() rgate.Rule {
	return rgate.Rule{
		Meta: rgate.RuleMeta{ID: "review-record", Level: rgate.LevelFail,
			Desc: "a REVIEW record exists with a reviewer context id and a disposition for every match-missing claim"},
		Check: func(ctx rgate.Context) (bool, quest.Fact) {
			sub := ctx.Submission.(*Submission)
			if strings.TrimSpace(sub.Review) == "" {
				return true, quest.Fact{Where: sub.ReviewPath,
					Expected: "REVIEW record file (reviewer: line + claim dispositions)",
					Actual:   "missing or empty"}
			}
			if reviewerOf(sub.Review) == "" {
				return true, quest.Fact{Where: sub.ReviewPath,
					Expected: "a 'reviewer: <context id>' line",
					Actual:   "no reviewer line"}
			}
			if ids := undisposed(sub.Review, sub.Missing); len(ids) > 0 {
				return true, quest.Fact{Where: sub.ReviewPath,
					Expected: "a '- <claim-id>: addressed|revised|excluded — ...' line per match-missing claim",
					Actual:   "undisposed claim(s): " + strings.Join(ids, ", ")}
			}
			return false, quest.Fact{}
		},
	}
}
