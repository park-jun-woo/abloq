//ff:func feature=blogyaml type=rule control=iteration dimension=1
//ff:what rulePriorityWeights 케이스 하나를 실행해 진단 수와 룰ID를 검증
package blogyaml

import "testing"

func checkRulePriorityWeights(t *testing.T, fetcher, train, gsc, citation int64, wantDiags int) {
	t.Helper()
	b := &Blog{Geo: Geo{PriorityWeights: PriorityWeights{
		Fetcher: fetcher, Train: train, GSC: gsc, Citation: citation,
	}}}
	diags := rulePriorityWeights("blog.yaml", b, lineIndex{})
	if len(diags) != wantDiags {
		t.Fatalf("want %d diagnostics, got %d: %v", wantDiags, len(diags), diags)
	}
	for _, d := range diags {
		if d.Rule != "priority-weights-range" {
			t.Errorf("want rule priority-weights-range, got %q", d.Rule)
		}
	}
}
