//ff:type feature=quest type=schema
//ff:what 게이트 Context.Submission — 단일 글 Target(Base nil 규약), match 미출현 claim, REVIEW 기록·작업 로그 본문
package writing

import (
	agate "github.com/park-jun-woo/abloq/pkg/gate"
	"github.com/park-jun-woo/abloq/pkg/insight"
)

// Submission is what one writing-quest submit carries through the gate: the
// assembled single-article target (Base nil — everything is judged as new),
// the insight-match missing claims (the REVIEW coverage input), and the
// review record / work log contents ("" when the file is absent — the
// review-record rule, not Prepare, judges absence).
type Submission struct {
	Target      *agate.Target
	Article     string
	Missing     []insight.Claim
	Review      string
	ReviewPath  string
	Worklog     string
	WorklogPath string
}
