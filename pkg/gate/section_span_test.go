//ff:func feature=gate type=parser control=iteration dimension=1 topic=evidence
//ff:what sectionSpan 케이스 — 중간 섹션은 다음 헤딩 직전까지, 마지막 섹션은 본문 끝까지, 미존재 키는 false
package gate

import "testing"

func TestSectionSpan(t *testing.T) {
	b := loadGateBlog(t)
	d := ParseArticle(b, "en", "Body.\n\n## Sources\n\n- a\n\n## Changelog\n\n- done\n")
	cases := []struct {
		name, key          string
		wantStart, wantEnd int
		wantOK             bool
	}{
		{"middle section ends at next heading", "sources", 2, 6, true},
		{"last section ends at body end", "changelog", 6, len(d.BodyLines), true},
		{"absent key", "related", 0, 0, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			start, end, ok := sectionSpan(d, tc.key)
			if start != tc.wantStart || end != tc.wantEnd || ok != tc.wantOK {
				t.Errorf("sectionSpan(%s) = (%d, %d, %v), want (%d, %d, %v)",
					tc.key, start, end, ok, tc.wantStart, tc.wantEnd, tc.wantOK)
			}
		})
	}
}
