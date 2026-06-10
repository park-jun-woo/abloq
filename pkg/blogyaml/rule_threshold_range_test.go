//ff:func feature=blogyaml type=rule control=iteration dimension=1
//ff:what ruleThresholdRange가 freshness_days<1, min_sources<0, min_internal_links<0을 각각 거부하는지 검증
package blogyaml

import "testing"

func TestRuleThresholdRange(t *testing.T) {
	cases := []struct {
		name                                        string
		freshnessDays, minSources, minInternalLinks int
		wantDiags                                   int
	}{
		{"all valid", 90, 1, 2, 0},
		{"boundary valid", 1, 0, 0, 0},
		{"freshness_days zero", 0, 1, 2, 1},
		{"min_sources negative", 90, -1, 2, 1},
		{"min_internal_links negative", 90, 1, -1, 1},
		{"all invalid", 0, -1, -1, 3},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			checkRuleThresholdRange(t, tc.freshnessDays, tc.minSources, tc.minInternalLinks, tc.wantDiags)
		})
	}
}
