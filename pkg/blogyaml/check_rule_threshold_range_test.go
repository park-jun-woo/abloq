//ff:func feature=blogyaml type=rule control=iteration dimension=1
//ff:what ruleThresholdRange 케이스 하나를 실행해 진단 수와 룰ID를 검증
package blogyaml

import "testing"

func checkRuleThresholdRange(t *testing.T, freshnessDays, minSources, minInternalLinks, wantDiags int) {
	t.Helper()
	b := &Blog{Geo: Geo{
		FreshnessDays:    freshnessDays,
		MinSources:       minSources,
		MinInternalLinks: minInternalLinks,
	}}
	diags := ruleThresholdRange("blog.yaml", b, lineIndex{})
	if len(diags) != wantDiags {
		t.Fatalf("want %d diagnostics, got %d: %v", wantDiags, len(diags), diags)
	}
	for _, d := range diags {
		if d.Rule != "threshold-range" {
			t.Errorf("want rule threshold-range, got %q", d.Rule)
		}
	}
}
