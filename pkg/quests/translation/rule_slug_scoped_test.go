//ff:func feature=quest type=rule control=sequence
//ff:what 스코프드 slug-consistency 검증 — 번역 front matter slug 이탈 시 발동, 동일 어간 쌍은 무발동
package translation

import (
	"strings"
	"testing"

	rgate "github.com/park-jun-woo/reins/pkg/gate"
)

func TestRuleSlugScoped(t *testing.T) {
	r := ruleSlugScoped()
	if r.Meta.ID != "slug-consistency" || r.Meta.Level != rgate.LevelFail {
		t.Errorf("Meta = %+v", r.Meta)
	}
	root := writeInstance(t)
	origin, ko := passPair()
	if fired, fact := fireRule(t, r, subWith(t, root, origin, ko)); fired {
		t.Errorf("same stem: fired with %+v", fact)
	}
	drifted := strings.Replace(ko, "title:", "slug: andere\ntitle:", 1)
	fired, fact := fireRule(t, r, subWith(t, root, origin, drifted))
	if !fired || fact.Actual != "andere" {
		t.Errorf("drifted slug: fired=%v fact=%+v", fired, fact)
	}
}
