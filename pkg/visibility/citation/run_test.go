//ff:func feature=visibility type=client control=sequence topic=citation
//ff:what Run이 엔진 순서대로 샘플을 합산하고 budget 0이면 no-op(nil)인지 검증
package citation

import "testing"

func TestRun(t *testing.T) {
	queries := []Query{{ID: 1, Text: "q1"}, {ID: 2, Text: "q2"}}
	cite := func(q string) ([]string, error) { return []string{"https://blog.test/p/"}, nil }
	engines := []Engine{{Name: "e1", Ask: cite}, {Name: "e2", Ask: cite}}

	samples := Run(engines, queries, 2, "blog.test", 0)
	if len(samples) != 4 {
		t.Fatalf("samples = %d, want 2 engines x 2 queries", len(samples))
	}
	if samples[0].Engine != "e1" || samples[1].Engine != "e1" ||
		samples[2].Engine != "e2" || samples[3].Engine != "e2" {
		t.Errorf("engine order = %s %s %s %s", samples[0].Engine, samples[1].Engine, samples[2].Engine, samples[3].Engine)
	}

	if got := Run(engines, queries, 0, "blog.test", 0); got != nil {
		t.Errorf("budget 0 = %v, want nil no-op", got)
	}
	if got := Run(nil, queries, 2, "blog.test", 0); got != nil {
		t.Errorf("no engines = %v, want nil", got)
	}
}
