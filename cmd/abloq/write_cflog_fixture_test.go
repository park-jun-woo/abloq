//ff:func feature=cli type=command control=sequence topic=crawl
//ff:what 임시 CF 로그 디렉토리 픽스처 — GPTBot 페이지·.md 히트와 curl 미지 봇이 담긴 .gz 1개 (writeBlogFixture 글 대상)
package main

import (
	"compress/gzip"
	"os"
	"path/filepath"
	"testing"
)

// writeCFLogFixture writes one gzip CloudFront log whose hits target the
// writeBlogFixture article (/ko/opinion/hello/, default_lang_in_subdir
// default true): two GPTBot page hits, one GPTBot .md hit, one curl line.
func writeCFLogFixture(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	row := func(tm, uri, status, ua string) string {
		return "2026-06-01\t" + tm + "\tICN54\t1024\t203.0.113.10\tGET\tt.example.com\t" +
			uri + "\t" + status + "\t-\t" + ua + "\t-\t-\tHit\trid\n"
	}
	body := "#Version: 1.0\n" +
		row("10:00:01", "/ko/opinion/hello/", "200", "Mozilla/5.0%20(compatible;%20GPTBot/1.2)") +
		row("10:00:02", "/ko/opinion/hello/", "200", "Mozilla/5.0%20(compatible;%20GPTBot/1.2)") +
		row("10:00:03", "/ko/opinion/hello.md", "200", "Mozilla/5.0%20(compatible;%20GPTBot/1.2)") +
		row("10:00:04", "/ko/opinion/hello/", "200", "curl/8.5.0")
	f, err := os.Create(filepath.Join(dir, "E1.2026-06-01-10.abcd1234.gz"))
	if err != nil {
		t.Fatalf("create: %v", err)
	}
	zw := gzip.NewWriter(f)
	if _, err := zw.Write([]byte(body)); err != nil {
		t.Fatalf("gzip write: %v", err)
	}
	if err := zw.Close(); err != nil {
		t.Fatalf("gzip close: %v", err)
	}
	if err := f.Close(); err != nil {
		t.Fatalf("close: %v", err)
	}
	return dir
}
