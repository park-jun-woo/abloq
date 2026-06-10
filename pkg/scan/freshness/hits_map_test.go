//ff:func feature=scan type=parser control=sequence
//ff:what HitsMap이 (lang,section,slug) 합계를 게이트 키로 인덱싱하는지 검증
package freshness

import "testing"

func TestHitsMap(t *testing.T) {
	m := HitsMap([]HitSum{{Lang: "ko", Section: "tech", Slug: "post-a", Hits: 9}})
	if m["ko/tech/post-a"] != 9 {
		t.Errorf("want 9, got %d", m["ko/tech/post-a"])
	}
	if len(HitsMap(nil)) != 0 {
		t.Error("empty input must yield empty map")
	}
}
