//ff:func feature=visibility type=parser control=sequence topic=report
//ff:what CiteHits가 글 키별 cited 적중 수를 합산하고(중복 키 가산) total은 무시하는지 검증
package report

import "testing"

func TestCiteHits(t *testing.T) {
	m := CiteHits([]CiteSum{
		{Lang: "ko", Section: "tech", Slug: "post-a", Cited: 2, Total: 3},
		{Lang: "ko", Section: "tech", Slug: "post-a", Cited: 1, Total: 5},
		{Lang: "ko", Section: "tech", Slug: "post-b", Cited: 0, Total: 4},
	})
	if m["ko/tech/post-a"] != 3 {
		t.Errorf("want 3 cited, got %d", m["ko/tech/post-a"])
	}
	if m["ko/tech/post-b"] != 0 {
		t.Errorf("want 0 cited, got %d", m["ko/tech/post-b"])
	}
	if len(CiteHits(nil)) != 0 {
		t.Error("empty input must yield empty map")
	}
}
