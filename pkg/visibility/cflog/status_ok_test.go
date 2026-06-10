//ff:func feature=visibility type=parser control=iteration dimension=1 topic=crawl
//ff:what statusOK가 2xx·304만 통과시키는지 검증 — 404/500/000/빈 값 제외
package cflog

import "testing"

func TestStatusOK(t *testing.T) {
	cases := []struct {
		status string
		want   bool
	}{
		{"200", true}, {"206", true}, {"304", true},
		{"404", false}, {"301", false}, {"500", false}, {"000", false}, {"", false}, {"2", false},
	}
	for _, c := range cases {
		if got := statusOK(c.status); got != c.want {
			t.Errorf("statusOK(%q) = %v, want %v", c.status, got, c.want)
		}
	}
}
