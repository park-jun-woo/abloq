//ff:func feature=gate type=parser control=iteration dimension=1 topic=evidence
//ff:what sourceCount 케이스 — -, *, n. 리스트 항목만 집계, 산문은 0, sources 섹션 없으면 found=false
package gate

import "testing"

func TestSourceCount(t *testing.T) {
	cases := []struct {
		name, body string
		wantN      int
		wantFound  bool
	}{
		{"dash items", "Body.\n\n## Sources\n\n- a\n- b\n", 2, true},
		{"star items", "Body.\n\n## Sources\n\n* a\n", 1, true},
		{"ordered items", "Body.\n\n## Sources\n\n1. a\n2. b\n3. c\n", 3, true},
		{"prose only", "Body.\n\n## Sources\n\nJust prose, no list.\n", 0, true},
		{"next section excluded", "Body.\n\n## Sources\n\n- a\n\n## Changelog\n\n- not a source\n", 1, true},
		{"section missing", "Body without any sources heading.\n", 0, false},
	}
	b := loadGateBlog(t)
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			n, _, found := sourceCount(ParseArticle(b, "en", tc.body))
			if n != tc.wantN || found != tc.wantFound {
				t.Errorf("sourceCount = (%d, %v), want (%d, %v)", n, found, tc.wantN, tc.wantFound)
			}
		})
	}
}
