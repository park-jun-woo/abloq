//ff:func feature=gate type=parser control=iteration dimension=1 topic=evidence
//ff:what isClaimLine 케이스 — 수치+단위+단정 조합만 true, 연도/버전/단위 없는 수/인라인 코드는 false
package gate

import "testing"

func TestIsClaimLine(t *testing.T) {
	cases := []struct {
		name, line string
		want       bool
	}{
		{"percent with improved", "Cache hit rate improved by 38% after the rollout.", true},
		{"korean ms claim", "응답 시간이 120ms로 줄었다.", true},
		{"korean multiple claim", "전환율이 3배 증가했다.", true},
		{"spaced unit", "Memory usage dropped to 512 MB.", true},
		{"korean percent record", "시장 점유율 45%를 기록했다.", true},
		{"date and version are not claims", "Released on 2026-01-01 with version 2.4.1.", false},
		{"bare year", "The 2024 report covers earlier work.", false},
		{"bare number", "See chapter 3 for details.", false},
		{"inline code stripped", "Set `timeout: 30%` in the config and improved nothing numeric.", false},
		{"unit without assertion", "The server runs on 8 GB of RAM.", false},
		{"assertion without unit", "Users grew to 10,000 by the end of the year.", false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := isClaimLine(tc.line); got != tc.want {
				t.Errorf("isClaimLine(%q) = %v, want %v", tc.line, got, tc.want)
			}
		})
	}
}
