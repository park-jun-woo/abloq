//ff:func feature=blogyaml type=parser control=iteration dimension=1
//ff:what joinPath가 빈 prefix면 key만, 아니면 "prefix.key"를 반환하는지 검증
package blogyaml

import "testing"

func TestJoinPath(t *testing.T) {
	cases := []struct {
		name, prefix, key, want string
	}{
		{"empty prefix", "", "site", "site"},
		{"with prefix", "site", "baseURL", "site.baseURL"},
		{"nested prefix", "geo.crawlers", "gptbot", "geo.crawlers.gptbot"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := joinPath(tc.prefix, tc.key); got != tc.want {
				t.Errorf("joinPath(%q, %q): want %q, got %q", tc.prefix, tc.key, tc.want, got)
			}
		})
	}
}
