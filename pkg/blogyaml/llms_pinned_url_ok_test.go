//ff:func feature=blogyaml type=rule control=iteration dimension=1
//ff:what llmsPinnedURLOK가 "/" 시작 경로와 호스트 있는 절대 http(s) URL만 참으로 판정하는지 검증
package blogyaml

import "testing"

func TestLlmsPinnedURLOK(t *testing.T) {
	cases := []struct {
		name, url string
		want      bool
	}{
		{"site-rooted path", "/reins.md", true},
		{"absolute https", "https://x.com/a", true},
		{"absolute http", "http://x.com", true},
		{"non-http scheme", "ftp://x.com/a", false},
		{"relative path", "docs/a.md", false},
		{"https without host", "https://", false},
		{"unparseable", "://bad", false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := llmsPinnedURLOK(tc.url); got != tc.want {
				t.Errorf("llmsPinnedURLOK(%q) = %v, want %v", tc.url, got, tc.want)
			}
		})
	}
}
