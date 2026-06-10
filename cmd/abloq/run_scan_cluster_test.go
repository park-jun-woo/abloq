//ff:func feature=cli type=command control=sequence
//ff:what runScanCluster가 위반 글의 큐 파일(key·candidates 포함)을 quests/queue/에 쓰고 정상 글은 제외하는지 검증
package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunScanCluster(t *testing.T) {
	dir := writeBlogFixture(t)
	// hello.md has no tags and one outlink below the min of 2 — it queues.
	// linked.md links hello twice over, so hello also has an inlink; linked
	// itself stays isolated (nothing points at it) and queues too.
	linked := "---\ntitle: Linked\ndate: 2026-01-03\n---\n\n[h](/opinion/hello/) 링크.\n"
	if err := os.WriteFile(filepath.Join(dir, "content", "ko", "opinion", "linked.md"), []byte(linked), 0o644); err != nil {
		t.Fatal(err)
	}
	var out bytes.Buffer
	if err := runScanCluster(&out, dir); err != nil {
		t.Fatalf("runScanCluster: %v", err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "quests", "queue", "cluster-ko-opinion-hello.yaml"))
	if err != nil {
		t.Fatalf("queue file missing: %v", err)
	}
	if !strings.Contains(string(data), "ko/opinion/hello") {
		t.Errorf("queue file must carry the gate join key: %s", data)
	}
	if !strings.Contains(string(data), "violations:") {
		t.Errorf("queue file must carry violations payload: %s", data)
	}
	if !strings.Contains(out.String(), "2 article(s) queued") {
		t.Errorf("summary line missing: %s", out.String())
	}
	if err := runScanCluster(&out, t.TempDir()); err == nil {
		t.Error("missing blog.yaml must error")
	}
	blocked := writeBlogFixture(t)
	if err := os.WriteFile(filepath.Join(blocked, "quests"), []byte("file, not a dir"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := runScanCluster(&out, blocked); err == nil {
		t.Error("unwritable quests/queue must error")
	}
}
