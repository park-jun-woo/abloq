//ff:func feature=visibility type=parser control=sequence topic=crawl
//ff:what LogKeys가 시간 프리픽스 없는 키를 걸러내는지 검증
package cflog

import (
	"reflect"
	"testing"
)

func TestLogKeys(t *testing.T) {
	got := LogKeys([]string{"E.2026-06-01-12.a.gz", "README.txt", "E.2026-06-01-13.b.gz"})
	want := []string{"E.2026-06-01-12.a.gz", "E.2026-06-01-13.b.gz"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("LogKeys = %v, want %v", got, want)
	}
}
