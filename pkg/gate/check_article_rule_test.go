//ff:func feature=gate type=rule control=sequence
//ff:what 글 단위 룰 케이스 실행 헬퍼 — 픽스처 1편으로 Target을 만들어 룰을 실행하고 진단을 검증
package gate

import (
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func checkArticleRule(t *testing.T, rule func(*Target) []blogyaml.Diagnostic, file, wantRule string, wantDiags int, wantMsgPart string) {
	t.Helper()
	b := loadGateBlog(t)
	a := artFromMD(t, b, "en", "tech", "fixture", file)
	tgt := NewTarget("testdata", b, []*Article{a})
	checkDiags(t, rule(tgt), wantDiags, wantRule, wantMsgPart)
}
