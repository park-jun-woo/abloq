//ff:func feature=gate type=parser control=iteration dimension=1 topic=evidence
//ff:what hasSourceLink 케이스 — 인라인 http(s) 링크와 각주 참조는 true, 내부 링크와 맨몸 URL은 false
package gate

import "testing"

func TestHasSourceLink(t *testing.T) {
	cases := []struct {
		name, text string
		want       bool
	}{
		{"inline https link", "Improved 42% ([bench](https://example.com/b)).", true},
		{"inline http link", "Improved 42% ([bench](http://example.com/b)).", true},
		{"footnote reference", "Latency dropped to 12ms.[^1]", true},
		{"internal link is not a source", "Improved 42% ([post](/en/tech/other/)).", false},
		{"bare url is not a markdown link", "Improved 42%, see https://example.com/b.", false},
		{"no link at all", "Improved 42% with no evidence.", false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := hasSourceLink(tc.text); got != tc.want {
				t.Errorf("hasSourceLink(%q) = %v, want %v", tc.text, got, tc.want)
			}
		})
	}
}
