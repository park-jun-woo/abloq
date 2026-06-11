//ff:func feature=quest type=rule control=sequence
//ff:what [slug-consistency] 스코프드 어댑터 — 아이템 단위로 원문↔번역 유효 slug(EffSlug) 일치만 검사
//ff:why 스톡 ruleSlugConsistency의 missing-lang 검출은 Target에 전 언어판이 없어 매 제출 FAIL→MaxTries 교착이므로 배제 — 언어 완전성은 Seed의 아이템 집합이 보장하고, 이 룰은 slug 동일성만 아이템 스코프로 본다 (Phase017 계획)
package translation

import (
	rgate "github.com/park-jun-woo/reins/pkg/gate"
	"github.com/park-jun-woo/reins/pkg/quest"

	agate "github.com/park-jun-woo/abloq/pkg/gate"
)

// ruleSlugScoped builds the scoped slug-consistency rule: this item's origin
// and translation must share one effective slug (front matter slug or file
// stem) so the language versions land on the same URL path.
func ruleSlugScoped() rgate.Rule {
	return rgate.Rule{
		Meta: rgate.RuleMeta{ID: "slug-consistency", Level: rgate.LevelFail,
			Desc: "origin and translation share one effective slug (scoped to this item's pair)"},
		Check: func(ctx rgate.Context) (bool, quest.Fact) {
			sub := ctx.Submission.(*Submission)
			want, got := agate.EffSlug(sub.Origin), agate.EffSlug(sub.Target.Articles[0])
			if got == want {
				return false, quest.Fact{}
			}
			return true, quest.Fact{Where: sub.Article + "#slug",
				Expected: "effective slug " + want + " (" + sub.OriginLang + " origin)",
				Actual:   got}
		},
	}
}
