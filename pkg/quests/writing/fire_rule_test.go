//ff:func feature=quest type=rule control=sequence
//ff:what 테스트 헬퍼 — 어댑터 룰 1개를 Submission에 평가해 (발동 여부, Fact) 반환
package writing

import (
	"testing"

	rgate "github.com/park-jun-woo/reins/pkg/gate"
	"github.com/park-jun-woo/reins/pkg/quest"
)

func fireRule(t *testing.T, r rgate.Rule, sub *Submission) (bool, quest.Fact) {
	t.Helper()
	return r.Check(rgate.Context{Submission: sub})
}
