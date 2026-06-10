//ff:func feature=blogyaml type=rule control=iteration dimension=1
//ff:what ruleSectionsEmpty가 빈 sections만 거부하는지 검증
package blogyaml

import "testing"

func TestRuleSectionsEmpty(t *testing.T) {
	cases := []struct {
		name      string
		sections  []string
		wantDiags int
	}{
		{"empty", nil, 1},
		{"non-empty", []string{"tech"}, 0},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) { checkRuleSectionsEmpty(t, tc.sections, tc.wantDiags) })
	}
}
