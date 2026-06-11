//ff:func feature=quest type=rule control=sequence topic=lossless
//ff:what [translation-parity] 원문↔번역 구조 패리티 룰 — ① 헤딩 시퀀스 ② 문단 수 ③ 이미지 ④ 코드블록 ⑤ 외부 링크 ⑥ 내부 링크 2단 ⑦ fm-mirror
//ff:why 기존 body-lossless는 동일 문서 라인 multiset(Base=HEAD 같은 파일)이라 원문↔번역 비교에 부적합 — 언어 독립 구조 룰을 신설한다. ContentLines는 정규화 텍스트라 언어 간 비교 실효용 없음, MultisetSubset만 재사용 (Phase017 계획)
package translation

import (
	rgate "github.com/park-jun-woo/reins/pkg/gate"
	"github.com/park-jun-woo/reins/pkg/quest"
)

// ruleTranslationParity builds the translation-parity rule: the language-
// independent structural comparison between the Prepare-loaded origin and the
// submitted translation. Any of the seven aspects firing is a FAIL.
func ruleTranslationParity() rgate.Rule {
	return rgate.Rule{
		Meta: rgate.RuleMeta{ID: "translation-parity", Level: rgate.LevelFail,
			Desc: "translation mirrors the origin's structure: headings, paragraph blocks, images, code blocks, links, date/lastmod"},
		Check: func(ctx rgate.Context) (bool, quest.Fact) {
			sub := ctx.Submission.(*Submission)
			facts := parityFacts(sub)
			if len(facts) == 0 {
				return false, quest.Fact{}
			}
			return true, firstFact(facts)
		},
	}
}
