package visibility

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// writeGSCRepo builds a blog repository with one recently-modified article
// (lastmod = today, so the inspection pass always selects it).
func writeGSCRepo(t *testing.T) string {
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
	today := time.Now().UTC().Format("2006-01-02")
	post := "---\ntitle: A\ndate: 2026-05-01\nlastmod: " + today + "\n---\nbody\n"
	if err := os.WriteFile(filepath.Join(postDir, "post-a.md"), []byte(post), 0o644); err != nil {
		t.Fatal(err)
	}
	return repo
}

// gscStub serves the token, Search Analytics and URL Inspection endpoints.
func gscStub(t *testing.T) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/token":
			w.Write([]byte(`{"access_token":"stub-token"}`))
		case strings.Contains(r.URL.EscapedPath(), "searchAnalytics/query"):
			w.Write([]byte(`{"rows":[{"keys":["https://t.example.com/tech/post-a/"],"clicks":3,"impressions":120,"position":4.2}]}`))
		case strings.HasSuffix(r.URL.Path, "index:inspect"):
			w.Write([]byte(`{"inspectionResult":{"indexStatusResult":{"verdict":"PASS","coverageState":"Submitted and indexed"}}}`))
		default:
			t.Errorf("unexpected path %q", r.URL.Path)
			http.NotFound(w, r)
		}
	}))
}

// TestIngestGSC runs the wrapper against the stub: the first run covers the
// lookback days and returns one row per day; the second run (cursor at the
// last closed day) is a no-op. The opt-in inspect pass returns the verdict
// summary of the recently-modified article.
func TestIngestGSC(t *testing.T) {
	srv := gscStub(t)
	defer srv.Close()
	repo := writeGSCRepo(t)
	t.Setenv("GSC_SEARCH_API_BASE", srv.URL)
	t.Setenv("GSC_TOKEN_URL", srv.URL+"/token")
	t.Setenv("GSC_SA_JSON", testSAJSON(t))
	t.Setenv("GSC_SA_JSON_PATH", "")
	t.Setenv("GSC_LOOKBACK_DAYS", "3")

	res, err := IngestGsc(IngestGscRequest{RepoPath: repo, Cursor: "", Inspect: false})
	if err != nil {
		t.Fatalf("IngestGSC: %v", err)
	}
	if res.Days != 3 || res.Rows != 3 {
		t.Errorf("days=%d rows=%d, want 3/3 (lookback)", res.Days, res.Rows)
	}
	if res.Inspected != 0 || res.InspectionsJSON != "[]" {
		t.Errorf("inspect off but inspected=%d inspections=%s", res.Inspected, res.InspectionsJSON)
	}
	var rows []map[string]any
	if err := json.Unmarshal(res.RowsJSON, &rows); err != nil {
		t.Fatalf("rows json: %v", err)
	}
	if len(rows) != 3 || rows[0]["page"] != "https://t.example.com/tech/post-a/" {
		t.Errorf("rows = %v", rows)
	}
	for _, key := range []string{"snap_date", "page", "impressions", "clicks", "avg_position"} {
		if _, ok := rows[0][key]; !ok {
			t.Errorf("row JSON lacks column key %q", key)
		}
	}

	// Cursor at the last closed day — nothing to collect.
	cursor := time.Now().UTC().AddDate(0, 0, -2).Format("2006-01-02")
	res, err = IngestGsc(IngestGscRequest{RepoPath: repo, Cursor: cursor, Inspect: false})
	if err != nil {
		t.Fatalf("up-to-date cursor: %v", err)
	}
	if res.Days != 0 || res.Rows != 0 || string(res.RowsJSON) != "[]" {
		t.Errorf("no-op run = %+v", res)
	}

	// Opt-in inspection: the today-lastmod article is selected and PASSes.
	res, err = IngestGsc(IngestGscRequest{RepoPath: repo, Cursor: cursor, Inspect: true})
	if err != nil {
		t.Fatalf("inspect run: %v", err)
	}
	if res.Inspected != 1 || !strings.Contains(res.InspectionsJSON, `"verdict":"PASS"`) {
		t.Errorf("inspected=%d inspections=%s", res.Inspected, res.InspectionsJSON)
	}

	// The site row's gsc_site_url wins over the baseURL derivation; envOr
	// returns the set value and the default.
	if site, _, err := siteProperty(repo, "sc-domain:t.example.com"); err != nil || site != "sc-domain:t.example.com" {
		t.Errorf("siteProperty with a row gsc_site_url = %q, %v", site, err)
	}
	if got := envOr("GSC_SITE_URL_UNSET_PROBE", "fallback"); got != "fallback" {
		t.Errorf("envOr empty = %q, want fallback", got)
	}

	// Search Analytics failure aborts the collection.
	t.Setenv("GSC_SEARCH_API_BASE", "http://127.0.0.1:1")
	if _, err := IngestGsc(IngestGscRequest{RepoPath: repo, Cursor: "", Inspect: false}); err == nil {
		t.Error("collect failure accepted")
	}
	t.Setenv("GSC_SEARCH_API_BASE", srv.URL)

	// Token failure aborts before any collection.
	t.Setenv("GSC_SA_JSON", `{"client_email":"x@test","private_key":"not-pem"}`)
	if _, err := IngestGsc(IngestGscRequest{RepoPath: repo}); err == nil {
		t.Error("bad SA key accepted")
	}
	t.Setenv("GSC_SA_JSON", testSAJSON(t))

	// Invalid blog.yaml (validation diagnostics) aborts.
	badRepo := t.TempDir()
	badBlog := "site:\n  baseURL: not-a-url\n  title: T\n  author: A\nlanguages: [ko]\nsections: [tech]\n"
	if err := os.WriteFile(filepath.Join(badRepo, "blog.yaml"), []byte(badBlog), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, _, err := siteProperty(badRepo, ""); err == nil {
		t.Error("invalid blog.yaml accepted")
	}
	if _, _, err := siteProperty(filepath.Join(badRepo, "missing"), ""); err == nil {
		t.Error("missing blog.yaml accepted")
	}

	if _, err := IngestGsc(IngestGscRequest{}); err == nil {
		t.Error("missing repo_path accepted")
	}
}

// TestInspectRecent covers the inspection error paths: an unreadable
// repository and a failing inspection endpoint.
func TestInspectRecent(t *testing.T) {
	repo := writeGSCRepo(t)
	t.Setenv("GSC_INSPECT_RECENT_DAYS", "36500")
	b, _, err := blogyaml.Load(filepath.Join(repo, "blog.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	now := time.Now().UTC()

	if _, err := inspectRecent("http://127.0.0.1:1", "tok", "https://t.example.com/", repo, b, now); err == nil {
		t.Error("failing inspection endpoint accepted")
	}

	srv := gscStub(t)
	defer srv.Close()
	ins, err := inspectRecent(srv.URL, "tok", "https://t.example.com/", repo, b, now)
	if err != nil || len(ins) != 1 || ins[0].Verdict != "PASS" {
		t.Errorf("inspectRecent = %+v, %v", ins, err)
	}

	// No recent articles — an empty (non-nil) summary, zero quota burnt.
	t.Setenv("GSC_INSPECT_RECENT_DAYS", "0")
	t.Setenv("GSC_INSPECT_MAX", "10")
	ins, err = inspectRecent(srv.URL, "tok", "https://t.example.com/", repo, b, now.AddDate(0, 0, 30))
	if err != nil || ins == nil || len(ins) != 0 {
		t.Errorf("no candidates = %+v, %v, want empty non-nil", ins, err)
	}

	if _, err := inspectRecent(srv.URL, "tok", "https://t.example.com/", filepath.Join(repo, "missing"), b, now); err == nil {
		t.Error("unreadable repository accepted")
	}
}
