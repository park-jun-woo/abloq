//ff:func feature=gate type=rule control=sequence topic=baseline
//ff:what 베이스라인 룰 케이스 실행 헬퍼 — base.md를 원본, 지정 픽스처를 현재본으로 Target을 만들어 룰 실행
package gate

import (
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func checkBaselineRule(t *testing.T, rule func(*Target) []blogyaml.Diagnostic, newFile, wantRule string, wantDiags int, wantMsgPart string) {
	t.Helper()
	b := loadGateBlog(t)
	a := artFromMD(t, b, "en", "tech", "base", newFile)
	a.Base = artFromMD(t, b, "en", "tech", "base", "baseline/base.md").Doc
	tgt := NewTarget("testdata", b, []*Article{a})
	checkDiags(t, rule(tgt), wantDiags, wantRule, wantMsgPart)
}
