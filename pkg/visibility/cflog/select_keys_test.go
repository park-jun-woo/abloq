//ff:func feature=visibility type=parser control=sequence topic=crawl
//ff:what selectKeys가 (커서, 마지막 닫힌 시간대] 반개구간의 로그 키만 남기는지 검증 — 비로그 키 제외
package cflog

import (
	"reflect"
	"testing"
)

func TestSelectKeys(t *testing.T) {
	keys := []string{
		"E.2026-06-01-11.x.gz",
		"E.2026-06-01-12.x.gz",
		"E.2026-06-01-13.x.gz",
		"E.2026-06-01-14.x.gz",
		"README.txt",
	}
	got := selectKeys(keys, "2026-06-01-11", "2026-06-01-13")
	want := []string{"E.2026-06-01-12.x.gz", "E.2026-06-01-13.x.gz"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("selectKeys = %v, want %v", got, want)
	}
	if got := selectKeys(keys, "", "2026-06-01-11"); len(got) != 1 {
		t.Errorf("empty cursor: %v, want only the 11 hour", got)
	}
}
