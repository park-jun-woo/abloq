//ff:func feature=gate type=parser control=iteration dimension=1
//ff:what headingHit이 정확 일치만 인식(어간/포함 매칭 거부)하고 레벨·키를 채우는지 검증
package gate

import "testing"

func TestHeadingHit(t *testing.T) {
	hi := buildHeadingIndex(loadGateBlog(t))
	cases := []struct {
		name, lang, ln, wantKey string
		wantLevel               int
		wantOK                  bool
	}{
		{"exact en", "en", "## Sources", "sources", 2, true},
		{"case fold", "en", "## SOURCES", "sources", 2, true},
		{"h3 recognized", "en", "### Related", "related", 3, true},
		{"ko exact", "ko", "## 출처", "sources", 2, true},
		{"stem must not match", "en", "## When Sources Are Attacked", "", 0, false},
		{"not a heading", "en", "Sources", "", 0, false},
		{"unknown lang", "fr", "## Sources", "", 0, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			hit, ok := headingHit(hi, tc.lang, tc.ln, 3)
			if ok != tc.wantOK || hit.Key != tc.wantKey || hit.Level != tc.wantLevel {
				t.Errorf("got (%+v, %v), want key=%s level=%d ok=%v", hit, ok, tc.wantKey, tc.wantLevel, tc.wantOK)
			}
		})
	}
}
