//ff:func feature=visibility type=parser control=sequence topic=report
//ff:what Tally.add가 category별 카운터를 나누고 md를 분류 불문 누적, 미지 category는 hits를 버리는지 검증
package report

import "testing"

func TestTallyAdd(t *testing.T) {
	var ta Tally
	ta.add("training", 7, 2)
	ta.add("search", 5, 0)
	ta.add("fetch", 3, 1)
	ta.add("unknown-category", 9, 4)
	want := Tally{Training: 7, Search: 5, Fetch: 3, MD: 7}
	if ta != want {
		t.Errorf("want %+v, got %+v", want, ta)
	}
}
