//ff:func feature=cli type=command control=sequence topic=report
//ff:what collectLogBotSums가 ym 형식 위반·URL맵 빌드 실패·로그 소스 부재·손상 .gz를 각각 에러로 내는지 검증
package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func TestCollectLogBotSumsErrors(t *testing.T) {
	b, _, err := blogyaml.Load("../../backend/fixtures/blog/blog.yaml")
	if err != nil {
		t.Fatal(err)
	}
	repo := "../../backend/fixtures/blog"
	logs := "../../backend/fixtures/cflogs"
	if _, _, err := collectLogBotSums(repo, logs, "not-a-ym", b); err == nil {
		t.Error("malformed ym must error")
	}
	if _, _, err := collectLogBotSums(t.TempDir()+"/missing", logs, "2026-06", b); err == nil {
		t.Error("a repo without content must error the URL map build")
	}
	if _, _, err := collectLogBotSums(repo, t.TempDir()+"/missing", "2026-06", b); err == nil {
		t.Error("a nonexistent log source must error")
	}
	// A corrupt .gz log object fails the wholesale ingest.
	badSrc := t.TempDir()
	if err := os.WriteFile(filepath.Join(badSrc, "E2X.2026-06-01-10.dead.gz"), []byte("not gzip"), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, _, err := collectLogBotSums(repo, badSrc, "2026-06", b); err == nil {
		t.Error("a corrupt log object must error")
	}
}
