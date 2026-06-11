//ff:func feature=quest type=frame control=sequence topic=queue
//ff:what Rules — 해소 재검 쌍(claims-resolved·rot-resolved) + 공용 큐 룰 2종(claim-scope·queue-scope) + 채택 3룰 어댑터(전부 LevelFail), 배제 룰 사유 명시
//ff:why 전부 LevelFail — LevelReview 룰은 Phase016 D1에서 교착(REVIEW 잠금이 재제출을 막음)이 확인돼 금지 (Phase017 선례)
package evidence

import (
	rgate "github.com/park-jun-woo/reins/pkg/gate"

	"github.com/park-jun-woo/abloq/pkg/quests/common"
)

// Rules is the evidence gate's catalog. The resolution re-check pair leads
// (work-completion forcing — an untouched article keeps its queued unsourced
// claims and rot citations, so an empty diff can never PASS), then the
// shared queue rules, then the adopted abloq rules wrapped by the shared
// adapter. All rules are LevelFail — reins level aggregation decides.
//
// Adopted abloq rules:
//   - numeric-claim-sourced: blocks the claim-rewording bypass — with the
//     HEAD baseline attached, a reworded claim is detected as a new claim
//     and must carry a source in its paragraph.
//   - min-sources: the sources section never shrinks below geo.min_sources.
//   - citation-exists: the replacement URLs must answer 200 and match the
//     cited title — sourcing with dead links is no sourcing.
//
// Excluded abloq rules (by design, one reason each):
//   - honest-lastmod / lastmod-advance: this quest does not force a lastmod
//     update — a sub-threshold change (one or two source links) must NOT
//     advance lastmod (the honesty convention; see context.md). When the
//     change is meaningful the repo-level honest-lastmod still governs the
//     update via the commit path.
//   - claim-preserved: claim-scope is stricter here — claims outside the
//     queue payload must survive verbatim, not merely in count.
//   - body-lossless / section-preserved / section-order / front-matter-* /
//     image-* rules: sourcing touches claim paragraphs and the sources
//     section only; claim-scope + queue-scope bound the blast radius, and
//     the repo/CI gate re-checks the full structure on commit.
//   - slug-consistency / hreflang-complete: language-set and built-output
//     rules — out of a single-article submission's scope.
func (Definition) Rules() []rgate.Rule {
	return []rgate.Rule{
		ruleClaimsResolved(),
		ruleRotResolved(),
		common.RuleClaimScope(),
		common.RuleQueueScope(),
		common.AdaptRule("numeric-claim-sourced"),
		common.AdaptRule("min-sources"),
		common.AdaptRule("citation-exists"),
	}
}
