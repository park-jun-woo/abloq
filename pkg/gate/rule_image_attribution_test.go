//ff:func feature=gate type=rule control=iteration dimension=1
//ff:what [image-attribution] 픽스처 케이스 — 표기 PASS 2건, 표기 누락·이미지 없음 FAIL, order 미선언 스킵 검증
package gate

import "testing"

func TestRuleImageAttribution(t *testing.T) {
	skip := loadGateBlog(t)
	skip.Structure.Order = []string{"image", "body", "sources"}
	bad := artFromMD(t, skip, "en", "tech", "no-attrib", "articles/no-attrib.md")
	if diags := ruleImageAttribution(NewTarget("testdata", skip, []*Article{bad})); len(diags) != 0 {
		t.Errorf("attribution undeclared: want 0 diagnostics, got %v", diags)
	}
	cases := []struct {
		name, file  string
		wantDiags   int
		wantMsgPart string
	}{
		{"pass canonical", "articles/pass.md", 0, ""},
		{"pass korean", "articles/pass-ko.md", 0, ""},
		{"fail missing attribution", "articles/no-attrib.md", 1, "must follow the main image"},
		{"fail no image", "articles/no-image.md", 1, "no main image to attribute"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			checkArticleRule(t, ruleImageAttribution, tc.file, "image-attribution", tc.wantDiags, tc.wantMsgPart)
		})
	}
}
