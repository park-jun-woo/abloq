//ff:func feature=gate type=frame control=iteration dimension=1
//ff:what special이 layout 키 보유 페이지만 특수로 분류하고 일반 글·front matter 없는 글·깨진 YAML을 일반으로 두는지 검증
package gate

import "testing"

func TestSpecial(t *testing.T) {
	b := loadGateBlog(t)
	hi := buildHeadingIndex(b)
	cases := []struct {
		name    string
		content string
		want    bool
	}{
		{"layout page", "---\nlayout: \"about\"\ntitle: t\n---\n\nbody\n", true},
		{"regular article", "---\ntitle: t\ndate: 2026-01-01\n---\n\nbody\n", false},
		{"no front matter", "just body\n", false},
		{"broken yaml", "---\ntitle: [unterminated\n---\n\nbody\n", false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			a := &Article{Lang: "en", Doc: parseDoc(hi, "en", tc.content)}
			if got := special(a); got != tc.want {
				t.Errorf("special = %v, want %v", got, tc.want)
			}
		})
	}
}
