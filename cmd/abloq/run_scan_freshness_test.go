//ff:func feature=cli type=command control=sequence
//ff:what runScanFreshness가 freshness_days 초과 글의 큐 파일(key 필드 포함)을 quests/queue/에 쓰고 신선 글은 제외하는지 검증
package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestRunScanFreshness(t *testing.T) {
	dir := writeBlogFixture(t)
	// The fixture post (lastmod 2026-01-02) is far past any threshold, but the
	// default freshness_days is 90 — pin a 1-day threshold to be explicit.
	blogYAML := "site:\n  baseURL: https://t.example.com\n  title: T\n  author: A\n" +
		"languages: [ko]\nsections: [opinion]\ngeo:\n  freshness_days: 1\n"
	if err := os.WriteFile(filepath.Join(dir, "blog.yaml"), []byte(blogYAML), 0o644); err != nil {
		t.Fatal(err)
	}
	fresh := "---\ntitle: Fresh\ndate: " + time.Now().UTC().Format("2006-01-02") + "\n---\nbody\n"
	if err := os.WriteFile(filepath.Join(dir, "content", "ko", "opinion", "fresh.md"), []byte(fresh), 0o644); err != nil {
		t.Fatal(err)
	}
	var out bytes.Buffer
	if err := runScanFreshness(&out, dir); err != nil {
		t.Fatalf("runScanFreshness: %v", err)
	}
	queueFile := filepath.Join(dir, "quests", "queue", "refresh-ko-opinion-hello.yaml")
	data, err := os.ReadFile(queueFile)
	if err != nil {
		t.Fatalf("queue file missing: %v", err)
	}
	if !strings.Contains(string(data), "ko/opinion/hello") {
		t.Errorf("queue file must carry the gate join key: %s", data)
	}
	if _, err := os.Stat(filepath.Join(dir, "quests", "queue", "refresh-ko-opinion-fresh.yaml")); !os.IsNotExist(err) {
		t.Error("fresh article must not be queued")
	}
	if !strings.Contains(out.String(), "1 stale article(s)") {
		t.Errorf("summary line missing: %s", out.String())
	}
}
