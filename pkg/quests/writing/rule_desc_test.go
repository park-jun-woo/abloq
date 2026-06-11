//ff:func feature=quest type=rule control=sequence
//ff:what ruleDesc 검증 — abloq 카탈로그 룰ID는 설명을 돌려주고 미지 ID는 빈 문자열
package writing

import "testing"

func TestRuleDesc(t *testing.T) {
	if got := ruleDesc("min-sources"); got == "" {
		t.Error("min-sources: want non-empty desc")
	}
	if got := ruleDesc("no-such-rule"); got != "" {
		t.Errorf("unknown id: desc = %q, want empty", got)
	}
}
