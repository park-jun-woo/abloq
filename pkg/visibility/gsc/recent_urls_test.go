//ff:func feature=visibility type=client control=sequence topic=gsc
//ff:what RecentURLs가 lastmod 최근 N일 경계(포함)·상한·비정상 lastmod 제외를 지키는지 검증
package gsc

import (
	"reflect"
	"testing"
	"time"
)

func TestRecentURLs(t *testing.T) {
	today := time.Date(2026, 6, 11, 3, 0, 0, 0, time.UTC)
	pages := []PageMod{
		{URL: "https://blog.test/a/", Lastmod: "2026-06-10"},
		{URL: "https://blog.test/b/", Lastmod: "2026-06-04"},
		{URL: "https://blog.test/c/", Lastmod: "2026-06-03"},
		{URL: "https://blog.test/d/", Lastmod: "bad"},
		{URL: "https://blog.test/e/", Lastmod: "2026-06-11T09:00:00+09:00"},
	}

	got := RecentURLs(pages, today, 7, 10)
	want := []string{"https://blog.test/a/", "https://blog.test/b/", "https://blog.test/e/"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("recent = %v, want %v", got, want)
	}

	if got := RecentURLs(pages, today, 7, 2); !reflect.DeepEqual(got, want[:2]) {
		t.Errorf("capped = %v, want %v", got, want[:2])
	}

	if got := RecentURLs(pages, today, 0, 10); !reflect.DeepEqual(got, []string{"https://blog.test/e/"}) {
		t.Errorf("zero-day window = %v, want only today's lastmod", got)
	}
}
