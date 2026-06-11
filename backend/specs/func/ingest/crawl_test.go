package ingest

import (
	"compress/gzip"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/visibility/cflog"
)

// writeFixtures builds a blog repository (one ko/tech article, root-served)
// and a log directory with one closed-hour .gz: two GPTBot page hits, one
// GPTBot .md hit and one PetalBot unknown candidate.
func writeFixtures(t *testing.T) (string, string) {
	t.Helper()
	repo := t.TempDir()
	blogYAML := "site:\n  baseURL: https://t.example.com\n  title: T\n  author: A\n" +
		"  default_lang_in_subdir: false\nlanguages: [ko]\nsections: [tech]\n"
	if err := os.WriteFile(filepath.Join(repo, "blog.yaml"), []byte(blogYAML), 0o644); err != nil {
		t.Fatal(err)
	}
	postDir := filepath.Join(repo, "content", "ko", "tech")
	if err := os.MkdirAll(postDir, 0o755); err != nil {
		t.Fatal(err)
	}
	post := "---\ntitle: A\ndate: 2026-05-01\n---\nbody\n"
	if err := os.WriteFile(filepath.Join(postDir, "post-a.md"), []byte(post), 0o644); err != nil {
		t.Fatal(err)
	}

	logs := t.TempDir()
	row := func(tm, uri, ua string) string {
		return "2026-06-01\t" + tm + "\tICN\t10\tip\tGET\thost\t" + uri + "\t200\t-\t" + ua + "\t-\t-\tHit\trid\n"
	}
	f, err := os.Create(filepath.Join(logs, "E1.2026-06-01-10.aaaa1111.gz"))
	if err != nil {
		t.Fatal(err)
	}
	zw := gzip.NewWriter(f)
	zw.Write([]byte(
		row("10:00:01", "/tech/post-a/", "GPTBot/1.2") +
			row("10:00:02", "/tech/post-a/", "GPTBot/1.2") +
			row("10:00:03", "/tech/post-a.md", "GPTBot/1.2") +
			row("10:00:04", "/", "PetalBot/1.0")))
	if err := zw.Close(); err != nil {
		t.Fatal(err)
	}
	if err := f.Close(); err != nil {
		t.Fatal(err)
	}
	return repo, logs
}

// TestCrawl runs the wrapper over the fixtures: the three JSON payloads must
// mirror the table columns, and a second run from the returned cursor must
// ingest nothing (zero duplicate accumulation).
func TestCrawl(t *testing.T) {
	repo, logs := writeFixtures(t)
	res, err := Crawl(CrawlRequest{RepoPath: repo, LogSource: logs, CursorsJSON: "[]"})
	if err != nil {
		t.Fatalf("Crawl: %v", err)
	}
	if res.Files != 1 || res.Hits != 3 {
		t.Errorf("Files, Hits = %d, %d, want 1, 3", res.Files, res.Hits)
	}
	var hits []cflog.HitRow
	if err := json.Unmarshal(res.HitsJSON, &hits); err != nil {
		t.Fatalf("HitsJSON: %v", err)
	}
	if len(hits) != 1 || hits[0].Bot != "GPTBot" || hits[0].Hits != 2 || hits[0].MDHits != 1 {
		t.Errorf("hits = %+v", hits)
	}
	if !strings.Contains(string(res.HitsJSON), `"hit_date":"2026-06-01"`) {
		t.Errorf("HitsJSON keys must mirror crawl_hits columns: %s", res.HitsJSON)
	}
	if !strings.Contains(string(res.UnknownJSON), `"ua":"PetalBot/1.0"`) {
		t.Errorf("UnknownJSON = %s", res.UnknownJSON)
	}
	var cursors []cflog.Cursor
	if err := json.Unmarshal(res.CursorsJSON, &cursors); err != nil {
		t.Fatalf("CursorsJSON: %v", err)
	}
	if len(cursors) != 1 || cursors[0].Source != cflog.CursorSource || cursors[0].CursorHour == "" {
		t.Errorf("cursors = %+v", cursors)
	}

	again, err := Crawl(CrawlRequest{RepoPath: repo, LogSource: logs, CursorsJSON: string(res.CursorsJSON)})
	if err != nil {
		t.Fatalf("re-Crawl: %v", err)
	}
	if again.Files != 0 || again.Hits != 0 {
		t.Errorf("re-ingest not idempotent: files=%d hits=%d", again.Files, again.Hits)
	}
}

// TestCrawlEnv pins the site-row/env contract: a missing log source or
// repo path fails, and CF_LOG_MARGIN_HOURS overrides the 2h default.
func TestCrawlEnv(t *testing.T) {
	if _, err := Crawl(CrawlRequest{CursorsJSON: "[]"}); err == nil {
		t.Error("missing cf_log_source accepted")
	}
	if _, err := Crawl(CrawlRequest{LogSource: t.TempDir(), CursorsJSON: "[]"}); err == nil {
		t.Error("missing repo_path accepted")
	}
	t.Setenv("CF_LOG_MARGIN_HOURS", "24")
	if got := marginHours(); got != 24*60*60*1e9 {
		t.Errorf("marginHours = %v, want 24h", got)
	}
	t.Setenv("CF_LOG_MARGIN_HOURS", "bogus")
	if got := marginHours(); got != 2*60*60*1e9 {
		t.Errorf("marginHours fallback = %v, want 2h", got)
	}
}
