//ff:func feature=quest type=rule control=sequence
//ff:what translation-parity 클린 경로 검증 — 패리티 통과 픽스처 쌍에서 미발동, 메타(ID·LevelFail·Desc) 확인
package translation

import (
	"testing"

	rgate "github.com/park-jun-woo/reins/pkg/gate"
)

func TestRuleTranslationParity(t *testing.T) {
	r := ruleTranslationParity()
	if r.Meta.ID != "translation-parity" || r.Meta.Level != rgate.LevelFail || r.Meta.Desc == "" {
		t.Errorf("Meta = %+v", r.Meta)
	}
	origin, ko := passPair()
	fired, fact := fireRule(t, r, subWith(t, writeInstance(t), origin, ko))
	if fired {
		t.Errorf("clean pair: fired with %+v", fact)
	}
}
