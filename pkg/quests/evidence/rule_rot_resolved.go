//ff:func feature=quest type=rule control=iteration dimension=1 topic=queue
//ff:what [rot-resolved] payload rot_urls가 Doc 인용에 잔존하는지 검증 — 잔존은 FAIL (작업 완수 강제, claims-resolved와 쌍)
package evidence

import (
	rgate "github.com/park-jun-woo/reins/pkg/gate"
	"github.com/park-jun-woo/reins/pkg/quest"

	agate "github.com/park-jun-woo/abloq/pkg/gate"
)

// ruleRotResolved builds the rot half of the evidence work-completion pair:
// none of the queue payload's confirmed-rot URLs may survive among the
// submitted article's citations — each must be replaced with a live source.
func ruleRotResolved() rgate.Rule {
	return rgate.Rule{
		Meta: rgate.RuleMeta{ID: "rot-resolved", Level: rgate.LevelFail,
			Desc: "no queued rot URL remains among the article's citations"},
		Check: func(ctx rgate.Context) (bool, quest.Fact) {
			sub := ctx.Submission.(*Submission)
			rotten := make(map[string]bool, len(sub.RotURLs))
			for _, u := range sub.RotURLs {
				rotten[u] = true
			}
			for _, c := range agate.Citations(sub.Target.Articles[0].Doc) {
				if !rotten[c.URL] {
					continue
				}
				return true, quest.Fact{Where: sub.Article,
					Expected: "the rotten citation replaced with a live source",
					Actual:   "confirmed-rot URL still cited: " + c.URL}
			}
			return false, quest.Fact{}
		},
	}
}
