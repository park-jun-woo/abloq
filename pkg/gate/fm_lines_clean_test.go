//ff:func feature=gate type=rule control=sequence topic=baseline
//ff:what fmLinesClean이 lastmod 라인 제외/우측 공백 제거/말미 빈 줄 제거를 수행하는지 검증
package gate

import (
	"reflect"
	"testing"
)

func TestFMLinesClean(t *testing.T) {
	got := fmLinesClean("title: x  \nlastmod: 2026-01-01\ndate: 1\n\n")
	want := []string{"title: x", "date: 1"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("fmLinesClean = %v, want %v", got, want)
	}
	if got := fmLinesClean("lastmod: 1\n"); len(got) != 0 {
		t.Errorf("lastmod-only: want empty, got %v", got)
	}
}
