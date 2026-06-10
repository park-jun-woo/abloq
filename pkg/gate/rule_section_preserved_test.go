//ff:func feature=gate type=rule control=iteration dimension=1 topic=baseline
//ff:what [section-preserved] 베이스라인 픽스처 케이스 — 재배열 PASS, 섹션 삭제 FAIL, 원본 없음·동일 스킵 검증
package gate

import "testing"

func TestRuleSectionPreserved(t *testing.T) {
	cases := []struct {
		name, newFile string
		wantDiags     int
		wantMsgPart   string
	}{
		{"pass reordered", "baseline/reordered.md", 0, ""},
		{"pass identical", "baseline/base.md", 0, ""},
		{"fail dropped section", "baseline/dropped-section.md", 1, "was removed"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			checkBaselineRule(t, ruleSectionPreserved, tc.newFile, "section-preserved", tc.wantDiags, tc.wantMsgPart)
		})
	}
}
