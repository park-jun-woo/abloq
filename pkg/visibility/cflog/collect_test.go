//ff:func feature=visibility type=parser control=sequence topic=crawl
//ff:what 골든 수집 테스트 — 픽스처 .gz 4파일의 기대 집계 전행 비교, 커서 시간 경계 전진, 재수집 0건(멱등)
package cflog

import (
	"reflect"
	"testing"
	"time"
)

// TestCollect runs one full collect over the committed testdata/logs
// fixtures and compares every aggregated row, the unknown-bot candidates,
// the advanced cursor and the totals; a re-collect from the returned cursor
// must ingest nothing (the Hurl duplicate-accumulation guarantee).
func TestCollect(t *testing.T) {
	root, b := writeRepoFixture(t)
	urls, err := BuildURLMap(root, b)
	if err != nil {
		t.Fatalf("BuildURLMap: %v", err)
	}
	src := DirSource{Root: "testdata/logs"}
	now := time.Date(2026, 6, 2, 4, 0, 0, 0, time.UTC)
	res, err := Collect(src, urls, nil, now, 2*time.Hour)
	if err != nil {
		t.Fatalf("Collect: %v", err)
	}
	wantHits := []HitRow{
		{HitDate: "2026-06-01", Bot: "Amazonbot", Lang: "ko", Section: "tech", Slug: "post-b", Hits: 1},
		{HitDate: "2026-06-01", Bot: "Bytespider", Lang: "ko", Section: "tech", Slug: "post-b", MDHits: 1},
		{HitDate: "2026-06-01", Bot: "ChatGPT-User", Lang: "ko", Section: "tech", Slug: "post-b", Hits: 1},
		{HitDate: "2026-06-01", Bot: "ClaudeBot", Lang: "en", Section: "tech", Slug: "post-a", Hits: 1},
		{HitDate: "2026-06-01", Bot: "ClaudeBot", Lang: "ko", Section: "tech", Slug: "post-a", Hits: 1},
		{HitDate: "2026-06-01", Bot: "GPTBot", Lang: "ko", Section: "tech", Slug: "post-a", Hits: 1, MDHits: 1},
		{HitDate: "2026-06-01", Bot: "OAI-SearchBot", Lang: "ko", Section: "tech", Slug: "post-a", Hits: 1},
		{HitDate: "2026-06-02", Bot: "ChatGPT-User", Lang: "en", Section: "tech", Slug: "post-a", MDHits: 1},
		{HitDate: "2026-06-02", Bot: "ClaudeBot", Lang: "ko", Section: "tech", Slug: "post-b", Hits: 1},
		{HitDate: "2026-06-02", Bot: "GPTBot", Lang: "ko", Section: "tech", Slug: "post-a", Hits: 2},
		{HitDate: "2026-06-02", Bot: "GoogleOther", Lang: "ko", Section: "tech", Slug: "post-a", Hits: 1},
		{HitDate: "2026-06-02", Bot: "PerplexityBot", Lang: "ko", Section: "tech", Slug: "post-b", Hits: 1},
	}
	if !reflect.DeepEqual(res.Hits, wantHits) {
		t.Errorf("Hits = %+v, want %+v", res.Hits, wantHits)
	}
	wantUnknown := []UnknownRow{
		{UA: "Mozilla/5.0 (compatible;PetalBot;+https://webmaster.petalsearch.com/site/petalbot)",
			Hits: 1, FirstSeen: "2026-06-01T12:00:08Z", LastSeen: "2026-06-01T12:00:08Z"},
		{UA: "curl/8.5.0", Hits: 1, FirstSeen: "2026-06-01T22:00:04Z", LastSeen: "2026-06-01T22:00:04Z"},
	}
	if !reflect.DeepEqual(res.Unknown, wantUnknown) {
		t.Errorf("Unknown = %+v, want %+v", res.Unknown, wantUnknown)
	}
	wantCursors := []Cursor{{Source: CursorSource, CursorHour: "2026-06-02-01"}}
	if !reflect.DeepEqual(res.Cursors, wantCursors) {
		t.Errorf("Cursors = %+v, want %+v", res.Cursors, wantCursors)
	}
	if res.Files != 4 || res.Total != 14 {
		t.Errorf("Files, Total = %d, %d, want 4, 14", res.Files, res.Total)
	}
	again, err := Collect(src, urls, res.Cursors, now, 2*time.Hour)
	if err != nil {
		t.Fatalf("re-Collect: %v", err)
	}
	if again.Files != 0 || again.Total != 0 || len(again.Hits) != 0 || len(again.Unknown) != 0 {
		t.Errorf("re-collect ingested something: files=%d total=%d hits=%d unknown=%d",
			again.Files, again.Total, len(again.Hits), len(again.Unknown))
	}
}
