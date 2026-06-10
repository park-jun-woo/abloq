//ff:func feature=visibility type=parser control=iteration dimension=1 topic=crawl
//ff:what hourOfKey가 CF 키의 시간 프리픽스를 뽑고 비로그 키는 false인지 검증
package cflog

import "testing"

func TestHourOfKey(t *testing.T) {
	cases := []struct {
		key  string
		want string
		ok   bool
	}{
		{"E123ABC.2026-06-01-12.aaaa1111.gz", "2026-06-01-12", true},
		{"logs/E123ABC.2026-06-01-12.aaaa1111.gz", "2026-06-01-12", true},
		{"README.txt", "", false},
		{"E123ABC.gz", "", false},
	}
	for _, c := range cases {
		got, ok := hourOfKey(c.key)
		if got != c.want || ok != c.ok {
			t.Errorf("hourOfKey(%q) = (%q, %v), want (%q, %v)", c.key, got, ok, c.want, c.ok)
		}
	}
}
