//ff:func feature=blogyaml type=rule control=iteration dimension=1
//ff:what ruleTaxonomyUnique가 중복 태그만 거부하고 부재·유일 목록은 통과시키는지 검증
package blogyaml

import "testing"

func TestRuleTaxonomyUnique(t *testing.T) {
	cases := []struct {
		name      string
		taxonomy  []string
		wantDiags int
	}{
		{"absent", nil, 0},
		{"unique", []string{"geo", "abloq"}, 0},
		{"one duplicate", []string{"geo", "abloq", "geo"}, 1},
		{"triple counts twice", []string{"geo", "geo", "geo"}, 2},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) { checkRuleTaxonomyUnique(t, tc.taxonomy, tc.wantDiags) })
	}
}
