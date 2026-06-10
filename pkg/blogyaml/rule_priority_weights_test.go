//ff:func feature=blogyaml type=rule control=iteration dimension=1
//ff:what rulePriorityWeights가 음수 가중치만 거부하고(0은 유효) 가중치별 진단을 내는지 검증
package blogyaml

import "testing"

func TestRulePriorityWeights(t *testing.T) {
	cases := []struct {
		name                          string
		fetcher, train, gsc, citation int64
		wantDiags                     int
	}{
		{"defaults valid", 3, 1, 1, 2, 0},
		{"all zero valid", 0, 0, 0, 0, 0},
		{"fetcher negative", -1, 1, 1, 2, 1},
		{"train negative", 3, -1, 1, 2, 1},
		{"gsc negative", 3, 1, -1, 2, 1},
		{"citation negative", 3, 1, 1, -2, 1},
		{"all negative", -1, -1, -1, -1, 4},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			checkRulePriorityWeights(t, tc.fetcher, tc.train, tc.gsc, tc.citation, tc.wantDiags)
		})
	}
}
