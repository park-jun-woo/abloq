//ff:func feature=cli type=command control=sequence
//ff:what runScanEvidence가 무출처 주장 글을 큐 파일로 쓰고(깨끗한 글 제외) 죽은 인용을 rot-check 한 줄로 보고하는지 검증
package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunScanEvidence(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "dead") {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()
	dir := writeBlogFixture(t)
	claims := "---\ntitle: Claims\ndate: 2026-01-02\n---\n\n처리량이 40% 증가했다.\n"
	if err := os.WriteFile(filepath.Join(dir, "content", "ko", "opinion", "claims.md"), []byte(claims), 0o644); err != nil {
		t.Fatal(err)
	}
	rot := "---\ntitle: Rot\ndate: 2026-01-02\n---\n\n[참고](" + srv.URL + "/dead-1) 링크.\n"
	if err := os.WriteFile(filepath.Join(dir, "content", "ko", "opinion", "rot.md"), []byte(rot), 0o644); err != nil {
		t.Fatal(err)
	}
	var out bytes.Buffer
	if err := runScanEvidence(&out, dir); err != nil {
		t.Fatalf("runScanEvidence: %v", err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "quests", "queue", "evidence-ko-opinion-claims.yaml"))
	if err != nil {
		t.Fatalf("claims queue file missing: %v", err)
	}
	if !strings.Contains(string(data), "claims:") || !strings.Contains(string(data), "40%") {
		t.Errorf("claims payload missing: %s", data)
	}
	if _, err := os.Stat(filepath.Join(dir, "quests", "queue", "evidence-ko-opinion-rot.yaml")); !os.IsNotExist(err) {
		t.Error("a stateless pass must never queue rot (3-strike rule is backend state)")
	}
	if !strings.Contains(out.String(), "rot-check: "+srv.URL+"/dead-1 hard (ko/opinion/rot)") {
		t.Errorf("rot report line missing: %s", out.String())
	}
	if !strings.Contains(out.String(), "1 article(s) queued, 1 citation(s) checked, 1 failing") {
		t.Errorf("summary line missing: %s", out.String())
	}
	if err := runScanEvidence(&out, t.TempDir()); err == nil {
		t.Error("dir without blog.yaml must error")
	}
	blocked := writeBlogFixture(t)
	if err := os.WriteFile(filepath.Join(blocked, "quests"), []byte("file in the way"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := runScanEvidence(&out, blocked); err == nil {
		t.Error("unwritable quests/queue must error")
	}
}
