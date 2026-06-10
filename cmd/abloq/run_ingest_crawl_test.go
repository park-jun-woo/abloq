//ff:func feature=cli type=command control=sequence topic=crawl
//ff:what runIngestCrawl이 픽스처 로그를 무상태 집계해 히트 행·원시 카운터·합계를 출력하는지, 소스·저장소 결손은 에러인지 검증
package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunIngestCrawl(t *testing.T) {
	repo := writeBlogFixture(t)
	logs := writeCFLogFixture(t)
	var out bytes.Buffer
	if err := runIngestCrawl(&out, logs, repo); err != nil {
		t.Fatalf("runIngestCrawl: %v", err)
	}
	got := out.String()
	if !strings.Contains(got, "2026-06-01  GPTBot               ko/opinion/hello  2 1") {
		t.Errorf("hit row missing or wrong:\n%s", got)
	}
	if !strings.Contains(got, "GPTBot               3") {
		t.Errorf("raw counter missing (3 unfiltered GPTBot lines):\n%s", got)
	}
	if !strings.Contains(got, `"curl/8.5.0" 1`) {
		t.Errorf("unknown bot candidate missing:\n%s", got)
	}
	if !strings.Contains(got, "ingest: 1 file(s), 3 article hit(s) (html+md), 1 unknown bot ua(s)") {
		t.Errorf("summary missing:\n%s", got)
	}
	if err := runIngestCrawl(&out, logs+"/missing", repo); err == nil {
		t.Error("missing source accepted")
	}
	if err := runIngestCrawl(&out, logs, t.TempDir()); err == nil {
		t.Error("repo without blog.yaml accepted")
	}
}
