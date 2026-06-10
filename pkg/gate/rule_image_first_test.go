//ff:func feature=gate type=rule control=iteration dimension=1
//ff:what [image-first] 픽스처 케이스 — 이미지 선두 PASS 2건, 텍스트 선두·이미지 후행 FAIL, order 미선언 스킵 검증
package gate

import "testing"

func TestRuleImageFirst(t *testing.T) {
	cases := []struct {
		name, file  string
		wantDiags   int
		wantMsgPart string
	}{
		{"pass canonical", "articles/pass.md", 0, ""},
		{"pass minimal", "articles/pass-minimal.md", 0, ""},
		{"fail text first", "articles/no-image.md", 1, "first content line"},
		{"fail image later", "articles/image-later.md", 1, "first content line"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			checkArticleRule(t, ruleImageFirst, tc.file, "image-first", tc.wantDiags, tc.wantMsgPart)
		})
	}
}
