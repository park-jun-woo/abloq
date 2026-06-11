//ff:func feature=quest type=rule control=iteration dimension=1 topic=queue
//ff:what [claims-resolved] Doc에 무출처 주장 검출을 재실행해 payload claims 해시와의 교집합이 ∅인지 검증 — 잔존은 FAIL (작업 완수 강제)
//ff:why 빈 diff 무작업 통과 차단(2차 검수 D1) — 큐가 지목한 주장이 여전히 무출처면 보강이 일어나지 않은 것이다. 검출기는 Phase010 스캐너와 동일(gate.DetectClaims+ClaimsExempt+HashText)이라 발급과 해소 판정이 갈라질 수 없다
package evidence

import (
	rgate "github.com/park-jun-woo/reins/pkg/gate"
	"github.com/park-jun-woo/reins/pkg/quest"

	agate "github.com/park-jun-woo/abloq/pkg/gate"
)

// ruleClaimsResolved builds the evidence work-completion rule: re-running
// the unsourced-claim detector over the submitted article must find no claim
// whose hash the queue payload listed. A claims_ignore exemption (reasons
// required, repo-gate audited) resolves like the scanner would stop issuing.
func ruleClaimsResolved() rgate.Rule {
	return rgate.Rule{
		Meta: rgate.RuleMeta{ID: "claims-resolved", Level: rgate.LevelFail,
			Desc: "no queued claim hash remains unsourced in the submitted article"},
		Check: func(ctx rgate.Context) (bool, quest.Fact) {
			sub := ctx.Submission.(*Submission)
			a := sub.Target.Articles[0]
			if agate.ClaimsExempt(a) {
				return false, quest.Fact{}
			}
			for _, c := range agate.DetectClaims(a.Doc) {
				if c.Sourced || !sub.QueuedClaims[agate.HashText(c.Text)] {
					continue
				}
				return true, quest.Fact{Where: sub.Article,
					Expected: "a source link in the claim's paragraph (or a justified replacement)",
					Actual:   "queued claim still unsourced: " + c.Text}
			}
			return false, quest.Fact{}
		},
	}
}
