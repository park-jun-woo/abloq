//ff:func feature=gate type=rule control=iteration dimension=1
//ff:what [section-order] 픽스처 케이스 — 정규 순서 PASS 2건, Sources→Related·Changelog→Sources 역전 FAIL 검증
package gate

import "testing"

func TestRuleSectionOrder(t *testing.T) {
	cases := []struct {
		name, file  string
		wantDiags   int
		wantMsgPart string
	}{
		{"pass canonical", "articles/pass.md", 0, ""},
		{"pass no sections", "articles/pass-minimal.md", 0, ""},
		{"fail related after sources", "articles/order-bad.md", 1, "must come before"},
		{"fail sources after changelog", "articles/order-bad-changelog.md", 1, "must come before"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			checkArticleRule(t, ruleSectionOrder, tc.file, "section-order", tc.wantDiags, tc.wantMsgPart)
		})
	}
}
