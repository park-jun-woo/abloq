//ff:func feature=quest type=frame control=sequence topic=queue
//ff:what Rules — lastmod-advance + 공용 큐 룰 2종(claim-preserved·queue-scope) + 채택 11룰 어댑터(전부 LevelFail), 배제 룰 사유 명시
//ff:why 전부 LevelFail — LevelReview 룰은 Phase016 D1에서 교착(REVIEW 잠금이 재제출을 막음)이 확인돼 금지 (Phase017 선례)
package refresh

import (
	rgate "github.com/park-jun-woo/reins/pkg/gate"

	"github.com/park-jun-woo/abloq/pkg/quests/common"
)

// Rules is the refresh gate's catalog. lastmod-advance leads (the
// work-completion forcing rule), then the shared queue rules, then the
// adopted abloq rules wrapped by the shared adapter. All rules are LevelFail
// — reins level aggregation decides.
//
// Adopted abloq rules (baseline-attached, so the comparison rules engage):
//   - honest-lastmod: the chain's core — an advanced lastmod requires a
//     meaningful body diff (>= geo.min_meaningful_diff) and queue membership
//     (the item's own queue file, present until the ② deletion commit).
//   - front-matter-intact: only lastmod may change vs HEAD.
//   - section-preserved / section-order / heading-canonical / image-first /
//     image-attribution / front-matter-schema: the refreshed article keeps
//     the structural shape.
//   - min-sources: the sources section never shrinks below geo.min_sources.
//   - numeric-claim-sourced: claims added since HEAD carry a source link.
//   - citation-exists: new citation URLs answer 200 and match the title.
//
// Excluded abloq rules (by design, one reason each):
//   - body-lossless: refreshing rewrites stale lines in place — a verbatim
//     line-multiset floor would make any refresh impossible. claim-preserved
//     (count floor) plus section-preserved carry the preservation duty.
//   - claim-scope: the refresh queue payload authorizes no specific claim
//     hashes — stale-figure replacement is the work itself; the count floor
//     (claim-preserved) is its deterministic proxy.
//   - slug-consistency: presumes sibling language versions in the target;
//     this gate inspects exactly one article (language companions are the
//     translation quest's job after ①).
//   - hreflang-complete: inspects built output (public/); nothing is built
//     at the refresh stage — the repo/CI `abloq gate` path keeps it.
func (Definition) Rules() []rgate.Rule {
	return []rgate.Rule{
		ruleLastmodAdvance(),
		common.RuleClaimPreserved(),
		common.RuleQueueScope(),
		common.AdaptRule("honest-lastmod"),
		common.AdaptRule("front-matter-intact"),
		common.AdaptRule("front-matter-schema"),
		common.AdaptRule("section-order"),
		common.AdaptRule("section-preserved"),
		common.AdaptRule("heading-canonical"),
		common.AdaptRule("image-first"),
		common.AdaptRule("image-attribution"),
		common.AdaptRule("min-sources"),
		common.AdaptRule("numeric-claim-sourced"),
		common.AdaptRule("citation-exists"),
	}
}
