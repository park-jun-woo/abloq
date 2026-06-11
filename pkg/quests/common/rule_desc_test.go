//ff:func feature=quest type=rule control=sequence
//ff:what RuleDesc 검증 — 카탈로그 룰ID의 설명 조회와 미등록 ID의 빈 문자열
package common

import "testing"

func TestRuleDesc(t *testing.T) {
	if d := RuleDesc("min-sources"); d == "" {
		t.Error("min-sources: want non-empty desc")
	}
	if d := RuleDesc("no-such-rule"); d != "" {
		t.Errorf("unknown id: desc = %q, want empty", d)
	}
}
