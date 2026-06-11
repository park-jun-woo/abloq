//ff:func feature=quest type=rule control=sequence topic=queue
//ff:what [cluster-resolved] pkg/scan/cluster.Scan을 인스턴스에 재실행해 payload가 지정한 위반 종류가 대상 글에서 소멸했는지 검증 — 잔존은 FAIL (작업 완수 강제)
//ff:why 빈 diff 무작업 통과 차단(2차 검수 D1) — 단, 스캔 경합으로 이미 해소된 아이템은 재스캔에 위반이 없어 빈 diff PASS가 된다(교착 아님, 정상 소비). 타 글의 위반은 무관하다 (Phase018 계획)
package cluster

import (
	rgate "github.com/park-jun-woo/reins/pkg/gate"
	"github.com/park-jun-woo/reins/pkg/quest"

	scancluster "github.com/park-jun-woo/abloq/pkg/scan/cluster"
)

// ruleClusterResolved builds the cluster work-completion rule: one fresh
// cluster scan over the working tree must no longer detect, on the target
// article, any of the violation kinds the queue payload listed. Detection is
// the Phase011 scanner itself — issuing and resolution can never diverge.
func ruleClusterResolved() rgate.Rule {
	return rgate.Rule{
		Meta: rgate.RuleMeta{ID: "cluster-resolved", Level: rgate.LevelFail,
			Desc: "the queued cluster violation kinds vanished from the target article (fresh scan)"},
		Check: func(ctx rgate.Context) (bool, quest.Fact) {
			sub := ctx.Submission.(*Submission)
			items := scancluster.Scan(sub.Target.Dir, sub.Target.Blog)
			rule, remains := remainingViolation(items, sub)
			if !remains {
				return false, quest.Fact{}
			}
			return true, quest.Fact{Where: sub.Article,
				Expected: "the queued violation kinds resolved on the target article",
				Actual:   "still detected by the cluster scan: " + rule}
		},
	}
}
