//ff:func feature=gate type=parser control=iteration dimension=1 topic=evidence
//ff:what ClaimsExempt 케이스 — 사유 목록만 true, 사유 누락(bad)·키 없음은 false (스캐너는 계속 스캔)
package gate

import "testing"

func TestClaimsExempt(t *testing.T) {
	cases := []struct {
		name, fm string
		want     bool
	}{
		{"key absent", "title: x", false},
		{"valid reason", "claims_ignore:\n  - \"own benchmark, raw data in repo\"", true},
		{"empty list is bad, not exempt", "claims_ignore: []", false},
		{"blank reason is bad, not exempt", "claims_ignore:\n  - \"   \"", false},
	}
	b := loadGateBlog(t)
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			a := artFromContent(t, b, "---\n"+tc.fm+"\n---\n\nBody.\n")
			if got := ClaimsExempt(a); got != tc.want {
				t.Errorf("ClaimsExempt = %v, want %v", got, tc.want)
			}
		})
	}
}
