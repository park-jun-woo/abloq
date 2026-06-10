//ff:func feature=visibility type=client control=sequence topic=gsc
//ff:what Dates가 빈 커서=lookback일, 커서 이후=다음 날부터 마진 경계까지, 최신 커서=빈 목록을 내는지 검증
package gsc

import (
	"reflect"
	"testing"
	"time"
)

func TestDates(t *testing.T) {
	today := time.Date(2026, 6, 11, 7, 30, 0, 0, time.UTC)

	got := Dates("", today, 2, 3)
	want := []string{"2026-06-07", "2026-06-08", "2026-06-09"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("first run = %v, want %v", got, want)
	}

	got = Dates("2026-06-07", today, 2, 28)
	want = []string{"2026-06-08", "2026-06-09"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("cursor advance = %v, want %v", got, want)
	}

	if got := Dates("2026-06-09", today, 2, 28); got != nil {
		t.Errorf("up-to-date cursor = %v, want nil", got)
	}
	if got := Dates("2026-06-10", today, 2, 28); got != nil {
		t.Errorf("cursor past the margin = %v, want nil", got)
	}
	if got := Dates("not-a-date", today, 2, 28); got != nil {
		t.Errorf("malformed cursor = %v, want nil", got)
	}
}
