//ff:func feature=gate type=parser control=iteration dimension=1 topic=evidence
//ff:what claimsIgnore 케이스 — 사유 목록은 exempt, 빈 목록/빈 문자열/비목록 값은 bad, 키 없음·깨진 FM은 둘 다 false
package gate

import "testing"

func TestClaimsIgnore(t *testing.T) {
	cases := []struct {
		name, fm            string
		wantExempt, wantBad bool
	}{
		{"key absent", "title: x", false, false},
		{"valid reason", "claims_ignore:\n  - \"own benchmark, raw data in repo\"", true, false},
		{"two reasons", "claims_ignore:\n  - \"reason one\"\n  - \"reason two\"", true, false},
		{"empty list", "claims_ignore: []", false, true},
		{"empty string entry", "claims_ignore:\n  - \"\"", false, true},
		{"blank reason", "claims_ignore:\n  - \"   \"", false, true},
		{"non-list value", "claims_ignore: because", false, true},
		{"broken yaml", "claims_ignore: [unclosed", false, false},
	}
	b := loadGateBlog(t)
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			a := artFromContent(t, b, "---\n"+tc.fm+"\n---\n\nBody.\n")
			exempt, bad := claimsIgnore(a)
			if exempt != tc.wantExempt || bad != tc.wantBad {
				t.Errorf("claimsIgnore = (%v, %v), want (%v, %v)", exempt, bad, tc.wantExempt, tc.wantBad)
			}
		})
	}
}
