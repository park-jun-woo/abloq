//ff:func feature=quest type=frame control=sequence
//ff:what Rules — translation-parity + 스코프드 slug-consistency + 채택 6룰 어댑터 + hugo-build 카탈로그(전부 LevelFail), 배제 룰 사유 명시
//ff:why 전부 LevelFail — LevelReview 룰은 Phase016 D1에서 교착(REVIEW 잠금이 재제출을 막음)이 확인돼 금지 (Phase017 계획)
package translation

import (
	rgate "github.com/park-jun-woo/reins/pkg/gate"

	"github.com/park-jun-woo/abloq/pkg/quests/common"
)

// Rules is the translation gate's catalog. translation-parity leads (root
// cause priority), then the scoped slug rule, then the translation-meaningful
// abloq rules wrapped by the shared adapter, then the whole-instance hugo
// build. All rules are LevelFail — reins level aggregation decides.
//
// Adopted abloq rules (translation-meaningful — they judge the submitted
// translation in its own language):
//   - front-matter-schema: the translation is a full article; required fields
//     and types apply as-is.
//   - image-first / image-attribution / section-order / heading-canonical:
//     per-article shape rules driven by blog.yaml structure and the target
//     language's heading map.
//   - min-sources: the translated sources section must list >= min entries
//     under the target language's recognized heading.
//
// Excluded abloq rules (by design, one reason each):
//   - section-preserved / body-lossless / front-matter-intact: baseline
//     (git HEAD same-file) comparisons — inert under the Base-nil convention;
//     origin-vs-translation preservation is translation-parity's job (①~⑥).
//   - slug-consistency (stock): its missing-lang check needs every language
//     version in the Target — a per-item submit would FAIL forever into the
//     MaxTries deadlock. Replaced by the scoped pair rule (ruleSlugScoped);
//     language completeness is guaranteed by Seed's item matrix.
//   - honest-lastmod: presumes a git baseline diff and queue membership; the
//     translation mirrors the origin's lastmod instead — fm-mirror (⑦)
//     replaces it (the Phase011 delegation, resolved for this chain).
//   - hreflang-complete: inspects built output (public/); the quest builds to
//     a throwaway dir and only locks build success — built-page checks remain
//     the repo/CI-level `abloq gate` path's job.
//   - numeric-claim-sourced: the claim detector is default-language-pattern
//     based — false hits/misses on CJK/RTL prose; the claim-source pairing
//     was already gated on the origin and link parity (⑤) preserves it.
//   - citation-exists: parity ⑤ forces URL equality with the origin, whose
//     citations the writing gate already verified — re-checking per language
//     multiplies network cost for zero new information.
func (Definition) Rules() []rgate.Rule {
	return []rgate.Rule{
		ruleTranslationParity(),
		ruleSlugScoped(),
		common.AdaptRule("front-matter-schema"),
		common.AdaptRule("image-first"),
		common.AdaptRule("image-attribution"),
		common.AdaptRule("section-order"),
		common.AdaptRule("heading-canonical"),
		common.AdaptRule("min-sources"),
		ruleHugoBuild(),
	}
}
