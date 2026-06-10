//ff:func feature=scan type=parser control=sequence
//ff:what SignalsMap이 (lang,section,slug) 합계를 게이트 키로 인덱싱해 Hits만 채우는지 검증
package freshness

import "testing"

func TestSignalsMap(t *testing.T) {
	m := SignalsMap([]HitSum{{Lang: "ko", Section: "tech", Slug: "post-a", Hits: 9}})
	if m["ko/tech/post-a"].Hits != 9 {
		t.Errorf("want Hits 9, got %d", m["ko/tech/post-a"].Hits)
	}
	if m["ko/tech/post-a"].FetcherHits != 0 || m["ko/tech/post-a"].GSCTrend != 0 {
		t.Errorf("base signals must carry Hits only: %+v", m["ko/tech/post-a"])
	}
	if len(SignalsMap(nil)) != 0 {
		t.Error("empty input must yield empty map")
	}
}
