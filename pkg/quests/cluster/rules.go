//ff:func feature=quest type=frame control=sequence topic=queue
//ff:what Rules — queue-scope(+candidates 확장) + cluster-resolved 2룰 카탈로그(전부 LevelFail), 배제 룰 사유 명시
//ff:why 전부 LevelFail — LevelReview 룰은 Phase016 D1에서 교착(REVIEW 잠금이 재제출을 막음)이 확인돼 금지 (Phase017 선례)
package cluster

import (
	rgate "github.com/park-jun-woo/reins/pkg/gate"

	"github.com/park-jun-woo/abloq/pkg/quests/common"
)

// Rules is the cluster gate's catalog: the queue-scope rule (its allowed set
// extended with the payload candidates — incoming links are how isolation
// resolves) and the resolution re-check. All rules are LevelFail — reins
// level aggregation decides.
//
// Excluded abloq rules (by design, one reason each):
//   - honest-lastmod / lastmod-advance: a candidate-article anchor line is a
//     sub-min_meaningful_diff change and must NOT advance lastmod (the
//     honesty convention — see context.md); nothing here forces an update.
//   - claim-scope / claim-preserved / numeric-claim-sourced: curation adds
//     tags and anchor links, never numeric claims; the repo/CI gate
//     re-checks claims on commit, and queue-scope bounds the blast radius.
//   - body-lossless / section rules / front-matter rules / min-sources /
//     citation-exists: an anchor line and a tags edit do not reshape the
//     article; the full structure re-check is the repo/CI `abloq gate`
//     path's job on commit.
//   - slug-consistency / hreflang-complete: language-set and built-output
//     rules — out of this submission's scope.
func (Definition) Rules() []rgate.Rule {
	return []rgate.Rule{
		common.RuleQueueScope(),
		ruleClusterResolved(),
	}
}
