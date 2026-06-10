//ff:func feature=visibility type=client control=sequence topic=citation
//ff:what runEngine이 budget 상한으로 id 순 실행하고 인용 매칭은 cited=true, 엔진 에러는 cited=false + 에러 근거 샘플로 남기는지 검증
package citation

import (
	"errors"
	"testing"
)

func TestRunEngine(t *testing.T) {
	queries := []Query{{ID: 1, Text: "q-one"}, {ID: 2, Text: "q-two"}, {ID: 3, Text: "q-three"}}
	engine := Engine{Name: "fake", Ask: func(q string) ([]string, error) {
		if q == "q-two" {
			return nil, errors.New("rate limited")
		}
		return []string{"https://blog.test/" + q + "/", "https://other.example.org/"}, nil
	}}

	samples := runEngine(engine, queries, 2, "blog.test", 0)
	if len(samples) != 2 {
		t.Fatalf("samples = %d, want budget 2", len(samples))
	}
	if !samples[0].Cited || samples[0].CitationQueriesID != 1 || samples[0].Engine != "fake" ||
		samples[0].Evidence != `{"matched":["https://blog.test/q-one/"]}` ||
		samples[0].ExtractorVersion != ExtractorVersion {
		t.Errorf("sample[0] = %+v", samples[0])
	}
	if samples[1].Cited || samples[1].CitationQueriesID != 2 ||
		samples[1].Evidence != `{"error":"rate limited"}` {
		t.Errorf("sample[1] = %+v", samples[1])
	}

	if got := runEngine(engine, queries, 99, "blog.test", 0); len(got) != 3 {
		t.Errorf("budget beyond queries = %d samples, want 3", len(got))
	}
}
