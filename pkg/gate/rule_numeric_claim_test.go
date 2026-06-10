//ff:func feature=gate type=rule control=iteration dimension=1 topic=evidence
//ff:what [numeric-claim-sourced] 픽스처 케이스 — 무출처 FAIL, 출처·각주 PASS, 코드/인용 제외, claims_ignore 예외와 사유 누락
package gate

import "testing"

func TestRuleNumericClaimSourced(t *testing.T) {
	cases := []struct {
		name, file  string
		wantDiags   int
		wantMsgPart string
	}{
		{"fail unsourced claim", "evidence/claims-unsourced.md", 1, "no source link"},
		{"pass inline link and footnote", "evidence/claims-sourced.md", 0, ""},
		{"pass code and quote blocks excluded", "evidence/claims-codeblock.md", 0, ""},
		{"pass exempt with reason", "evidence/claims-ignore.md", 0, ""},
		{"fail exemption without reason", "evidence/claims-ignore-bad.md", 1, "reason"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			checkArticleRule(t, ruleNumericClaimSourced, tc.file, "numeric-claim-sourced", tc.wantDiags, tc.wantMsgPart)
		})
	}
}
