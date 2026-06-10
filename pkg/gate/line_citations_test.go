//ff:func feature=gate type=parser control=iteration dimension=1 topic=evidence
//ff:what lineCitations 케이스 — 외부 링크 추출, 이미지·내부 링크 제외, 한 줄 다중 링크
package gate

import "testing"

func TestLineCitations(t *testing.T) {
	cases := []struct {
		name, line string
		wantURLs   []string
	}{
		{"one external link", "See [bench](https://example.com/b).", []string{"https://example.com/b"}},
		{"image is not a citation", "![alt](https://example.com/i.webp)", nil},
		{"two links on one line", "[a](http://x.test/a) and [b](https://y.test/b)", []string{"http://x.test/a", "https://y.test/b"}},
		{"internal link ignored", "[other](/en/tech/other/)", nil},
		{"image then link", "![alt](https://x.test/i.png) per [src](https://x.test/s)", []string{"https://x.test/s"}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			checkCitationURLs(t, lineCitations(tc.line, 7), tc.wantURLs, 7)
		})
	}
}
