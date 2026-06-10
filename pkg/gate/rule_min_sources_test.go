//ff:func feature=gate type=rule control=iteration dimension=1 topic=evidence
//ff:what [min-sources] 픽스처 케이스 — 출처 2건 PASS, 빈 섹션 FAIL, 섹션 누락 FAIL (임계 기본값 1)
package gate

import "testing"

func TestRuleMinSources(t *testing.T) {
	cases := []struct {
		name, file  string
		wantDiags   int
		wantMsgPart string
	}{
		{"pass two sources", "evidence/sources-ok.md", 0, ""},
		{"fail empty section", "evidence/sources-empty.md", 1, "lists 0 source(s)"},
		{"fail missing section", "evidence/sources-missing.md", 1, "sources section missing"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			checkArticleRule(t, ruleMinSources, tc.file, "min-sources", tc.wantDiags, tc.wantMsgPart)
		})
	}
}
