//ff:func feature=cli type=command control=sequence topic=report
//ff:what collectLogBotSums가 CF 로그 픽스처를 현·전월 윈도로 나눠 봇별 합계를 내고 source 비면 빈 집계인지 검증
package main

import (
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func TestCollectLogBotSums(t *testing.T) {
	b, diags, err := blogyaml.Load("../../backend/fixtures/blog/blog.yaml")
	if err != nil || len(diags) > 0 {
		t.Fatalf("fixture blog.yaml: %v %v", err, diags)
	}
	repo := "../../backend/fixtures/blog"
	// Fixture log hits are all on 2026-06-01 — inside the 2026-06 window
	// ([2026-06-01 .. 2026-06-30]) and inside 2026-07's previous window.
	bots, prevBots, err := collectLogBotSums(repo, "../../backend/fixtures/cflogs", "2026-06", b)
	if err != nil {
		t.Fatalf("collectLogBotSums: %v", err)
	}
	if len(bots) != 3 {
		t.Fatalf("want 3 bot rows in the 2026-06 window, got %d: %+v", len(bots), bots)
	}
	if len(prevBots) != 0 {
		t.Errorf("2026-05 window must be empty: %+v", prevBots)
	}
	// One month later the same rows land in the previous window.
	bots, prevBots, err = collectLogBotSums(repo, "../../backend/fixtures/cflogs", "2026-07", b)
	if err != nil {
		t.Fatal(err)
	}
	if len(bots) != 0 || len(prevBots) != 3 {
		t.Errorf("want 0 current / 3 prev rows for 2026-07: %d %d", len(bots), len(prevBots))
	}
	// No source — the crawl layer stays empty.
	bots, prevBots, err = collectLogBotSums(repo, "", "2026-06", b)
	if err != nil || bots != nil || prevBots != nil {
		t.Errorf("empty source must yield empty sums: %v %v %v", bots, prevBots, err)
	}
}
