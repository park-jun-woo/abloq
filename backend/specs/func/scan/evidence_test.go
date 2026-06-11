package scan

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	pqueueio "github.com/park-jun-woo/abloq/pkg/queueio"
	pevidence "github.com/park-jun-woo/abloq/pkg/scan/evidence"
)

func TestEvidence(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "dead") {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()
	dir := t.TempDir()
	blogYAML := "site:\n  baseURL: https://t.example.com\n  title: T\n  author: A\n" +
		"languages: [ko]\nsections: [tech]\n"
	if err := os.WriteFile(filepath.Join(dir, "blog.yaml"), []byte(blogYAML), 0o644); err != nil {
		t.Fatal(err)
	}
	postDir := filepath.Join(dir, "content", "ko", "tech")
	if err := os.MkdirAll(postDir, 0o755); err != nil {
		t.Fatal(err)
	}
	claims := "---\ntitle: Claims\ndate: 2026-06-01\n---\n\n처리량이 40% 증가했다.\n"
	if err := os.WriteFile(filepath.Join(postDir, "post-claims.md"), []byte(claims), 0o644); err != nil {
		t.Fatal(err)
	}
	rot := "---\ntitle: Rot\ndate: 2026-06-02\n---\n\n[참고](https://example.org/dead-1) 링크.\n"
	if err := os.WriteFile(filepath.Join(postDir, "post-rot.md"), []byte(rot), 0o644); err != nil {
		t.Fatal(err)
	}
	t.Setenv("LINKCHECK_HOST_OVERRIDE", srv.URL)

	// First scan: the claims item only; the dead citation is at 1 failure.
	resp, err := Evidence(EvidenceRequest{RepoPath: dir, ChecksJSON: "[]"})
	if err != nil {
		t.Fatalf("Evidence: %v", err)
	}
	if resp.Detected != 1 {
		t.Fatalf("want 1 detected, got %d", resp.Detected)
	}
	rows, err := pqueueio.DecodeRows(resp.ItemsJSON)
	if err != nil || len(rows) != 1 || rows[0].Kind != "evidence" || rows[0].Slug != "post-claims" {
		t.Fatalf("ItemsJSON must hold the claims item: %v %+v", err, rows)
	}
	var checks []pevidence.Check
	if err := json.Unmarshal(resp.ChecksJSON, &checks); err != nil || len(checks) != 1 {
		t.Fatalf("ChecksJSON must hold 1 probe state: %v %+v", err, checks)
	}
	if checks[0].Status != "hard" || checks[0].ConsecutiveFailures != 1 {
		t.Errorf("dead citation state: %+v", checks[0])
	}

	// Third consecutive failure (previous state says 2): rot confirms.
	checks[0].ConsecutiveFailures = 2
	prevJSON, _ := json.Marshal(checks)
	resp, err = Evidence(EvidenceRequest{RepoPath: dir, ChecksJSON: string(prevJSON)})
	if err != nil {
		t.Fatal(err)
	}
	if resp.Detected != 2 {
		t.Fatalf("rot must confirm at 3 failures: detected = %d", resp.Detected)
	}
	rows, _ = pqueueio.DecodeRows(resp.ItemsJSON)
	if rows[1].Slug != "post-rot" || !strings.Contains(rows[1].Payload["rot_urls"], "dead-1") {
		t.Errorf("rot item: %+v", rows[1])
	}

	if _, err := Evidence(EvidenceRequest{RepoPath: dir, ChecksJSON: "not-json"}); err == nil {
		t.Error("broken previous-state JSON must error")
	}

	if _, err := Evidence(EvidenceRequest{RepoPath: t.TempDir(), ChecksJSON: "[]"}); err == nil {
		t.Error("missing blog.yaml must error")
	}

	badDir := t.TempDir()
	badYAML := "site:\n  baseURL: not-a-url\n  title: T\n  author: A\nlanguages: [ko]\nsections: [tech]\n"
	if err := os.WriteFile(filepath.Join(badDir, "blog.yaml"), []byte(badYAML), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := Evidence(EvidenceRequest{RepoPath: badDir, ChecksJSON: "[]"}); err == nil {
		t.Error("invalid blog.yaml must error")
	}

	if _, err := Evidence(EvidenceRequest{ChecksJSON: "[]"}); err == nil {
		t.Error("missing repo_path must error")
	}
}
