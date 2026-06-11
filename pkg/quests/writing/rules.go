//ff:func feature=quest type=frame control=sequence
//ff:what Rules — 채택 11룰 어댑터(전부 LevelFail) + review-record 룰의 카탈로그, 배제 3룰 사유 명시
package writing

import (
	rgate "github.com/park-jun-woo/reins/pkg/gate"
)

// Rules is the writing gate's catalog: eleven abloq pkg/gate rules wrapped by
// adaptRule (no new detection logic) plus the review-record coverage rule.
// All rules are LevelFail — reins level aggregation (any Fail → FAIL) decides.
//
// Excluded abloq rules (unsuitable at the writing stage, by design):
//   - slug-consistency: presumes sibling language versions; the writing quest
//     gates exactly one default-language article.
//   - hreflang-complete: presumes built output (public/); nothing is built at
//     the writing stage.
//   - honest-lastmod: presumes a git baseline diff and queue membership; it
//     cannot coexist with the Base-nil (everything-is-new) convention.
//
// Note: section-preserved, body-lossless and front-matter-intact are inert
// under Base nil (new article) and engage only when a future caller attaches
// a baseline — kept in the chain so editing flows inherit them unchanged.
func (Definition) Rules() []rgate.Rule {
	return []rgate.Rule{
		adaptRule("image-first"),
		adaptRule("image-attribution"),
		adaptRule("section-order"),
		adaptRule("section-preserved"),
		adaptRule("body-lossless"),
		adaptRule("front-matter-intact"),
		adaptRule("heading-canonical"),
		adaptRule("front-matter-schema"),
		adaptRule("min-sources"),
		adaptRule("numeric-claim-sourced"),
		adaptRule("citation-exists"),
		ruleReviewRecord(),
	}
}
