//ff:func feature=gate type=rule control=iteration dimension=1
//ff:what [heading-canonical] 픽스처 케이스 — ## 레벨 PASS 2건, ### 레벨 인식 헤딩 FAIL 검증
package gate

import "testing"

func TestRuleHeadingCanonical(t *testing.T) {
	cases := []struct {
		name, file  string
		wantDiags   int
		wantMsgPart string
	}{
		{"pass canonical", "articles/pass.md", 0, ""},
		{"pass no sections", "articles/pass-minimal.md", 0, ""},
		{"fail h3 sources", "articles/heading-h3.md", 1, "must be ## level"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			checkArticleRule(t, ruleHeadingCanonical, tc.file, "heading-canonical", tc.wantDiags, tc.wantMsgPart)
		})
	}
}
