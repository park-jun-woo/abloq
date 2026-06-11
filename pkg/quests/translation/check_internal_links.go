//ff:func feature=quest type=rule control=iteration dimension=1 topic=lossless
//ff:what 패리티 ⑥ — (a) 번역 내부 글 링크의 자기 언어 프리픽스(/{lang}/) 필수 (b) 양쪽 프리픽스 제거 후 양방향 multiset 일치
//ff:why (a)가 없으면 자기모순: 기본 언어가 루트 서빙이라 정규화 비교(b)만으로는 미치환 링크(원문 그대로)가 비검출된다 — 프리픽스 강제가 치환 누락을 직접 잡는다 (Phase017 계획 ⑥)
package translation

import (
	"github.com/park-jun-woo/reins/pkg/quest"
)

// checkInternalLinks runs the two-stage internal article link check: every
// internal article link in the translation must carry the translation's own
// language prefix, and after stripping the prefix on BOTH sides the link
// multisets must match bidirectionally.
func checkInternalLinks(sub *Submission) []quest.Fact {
	tLinks := internalLinks(sub.Target.Articles[0].Doc)
	var facts []quest.Fact
	pre := "/" + sub.Lang + "/"
	for _, l := range tLinks {
		if stripLangPrefix(l, sub.Lang) != l {
			continue
		}
		facts = append(facts, quest.Fact{Where: sub.Article + "#internal-links",
			Expected: "internal article links prefixed with " + pre,
			Actual:   "unprefixed link: " + l})
		break
	}
	oNeutral := stripAll(internalLinks(sub.Origin.Doc), sub.OriginLang)
	tNeutral := stripAll(tLinks, sub.Lang)
	return append(facts, checkMultiset(sub.Article+"#internal-links",
		"internal article link multiset (language prefix stripped)", oNeutral, tNeutral)...)
}
