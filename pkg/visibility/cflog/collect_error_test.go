//ff:func feature=visibility type=parser control=sequence topic=crawl
//ff:what Collect의 에러 경로 — List 실패와 파일 수집 실패가 그대로 전파되는지 검증
package cflog

import (
	"errors"
	"testing"
	"time"
)

func TestCollectErrors(t *testing.T) {
	now := time.Date(2026, 6, 2, 4, 0, 0, 0, time.UTC)
	boom := errors.New("boom")
	if _, err := Collect(fakeSource{listErr: boom}, nil, nil, now, 0); !errors.Is(err, boom) {
		t.Errorf("List error not propagated: %v", err)
	}
	src := fakeSource{keys: []string{"E.2026-06-01-12.a.gz"}, getErr: boom}
	if _, err := Collect(src, nil, nil, now, 0); !errors.Is(err, boom) {
		t.Errorf("Get error not propagated: %v", err)
	}
}
