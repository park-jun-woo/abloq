//ff:func feature=visibility type=parser control=sequence topic=crawl
//ff:what NewAgg가 비어 있는 카운터 맵들과 URI 역매핑을 갖춘 누적기를 만드는지 검증
package cflog

import "testing"

func TestNewAgg(t *testing.T) {
	urls := map[string]Article{"/a/": {Lang: "ko", Section: "tech", Slug: "a"}}
	agg := NewAgg(urls)
	if agg.URLs == nil || agg.hits == nil || agg.unknown == nil || agg.Raw == nil {
		t.Fatalf("NewAgg left a nil map: %+v", agg)
	}
	if len(agg.HitRows()) != 0 || len(agg.UnknownRows()) != 0 {
		t.Errorf("fresh accumulator is not empty")
	}
}
