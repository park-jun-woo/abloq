package visibility

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/park-jun-woo/abloq/pkg/archive"
	"github.com/park-jun-woo/abloq/pkg/blogyaml"
	pcontent "github.com/park-jun-woo/abloq/pkg/content"
	"github.com/park-jun-woo/abloq/pkg/visibility/gsc"
)

// @func ingestGsc
// @error 500
// @description Run one incremental GSC collection for one site: exchange the site's service-account assertion (gsc_sa_json_path; empty falls back to the global GSC_SA_JSON/GSC_SA_JSON_PATH env) for a webmasters.readonly token (scope is a parameter of pkg/archive — never the archiver's indexing scope), list the closed days after the MAX(snap_date) cursor (today UTC minus GSC_DELAY_MARGIN_DAYS, default 2; first run covers GSC_LOOKBACK_DAYS, default 28) and fetch the per-page Search Analytics rows for each. When the request opts in, additionally inspect the URLs of articles whose repository-parsed lastmod falls within GSC_INSPECT_RECENT_DAYS (default 7), at most GSC_INSPECT_MAX (default 10) of them, and return the verdict summaries without storing them. The site property is the row's gsc_site_url or the site blog.yaml baseURL. The abloq CLI shares the same pkg statelessly (`abloq ingest gsc`)

type IngestGscRequest struct {
	RepoPath   string
	SiteURL    string
	SAJSONPath string
	Cursor     string
	Inspect    bool
}

type IngestGscResponse struct {
	RowsJSON        []byte
	Days            int64
	Rows            int64
	Inspected       int64
	InspectionsJSON string
}

// IngestGsc is the thin @call wrapper around pkg/visibility/gsc: the cursor
// rides in as a scalar and the site identity (repo path, GSC property, SA
// JSON path) rides in from the site row — a site's collection must never
// run under another site's credentials (multisite isolation); only the
// tuning margins stay env. The row JSON field names mirror the
// gsc_snapshots columns, so the backend's upsert and the pkg's output feed
// the same batch.
func IngestGsc(req IngestGscRequest) (IngestGscResponse, error) {
	site, b, err := siteProperty(req.RepoPath, req.SiteURL)
	if err != nil {
		return IngestGscResponse{}, err
	}
	token, err := archive.GSCTokenWith(archive.Keys{GSCSAJSONPath: req.SAJSONPath},
		archive.ScopeWebmastersReadonly)
	if err != nil {
		return IngestGscResponse{}, err
	}
	base := envOr("GSC_SEARCH_API_BASE", "https://searchconsole.googleapis.com")
	now := time.Now().UTC()
	dates := gsc.Dates(req.Cursor, now, intEnv("GSC_DELAY_MARGIN_DAYS", 2), intEnv("GSC_LOOKBACK_DAYS", 28))
	res, err := gsc.Collect(base, token, site, dates)
	if err != nil {
		return IngestGscResponse{}, err
	}
	if res.Rows == nil {
		res.Rows = []gsc.Snapshot{}
	}
	// Marshal cannot fail: the row types are plain string/number fields.
	rowsJSON, _ := json.Marshal(res.Rows)
	out := IngestGscResponse{
		RowsJSON:        rowsJSON,
		Days:            res.Days,
		Rows:            int64(len(res.Rows)),
		InspectionsJSON: "[]",
	}
	if !req.Inspect {
		return out, nil
	}
	inspections, err := inspectRecent(base, token, site, req.RepoPath, b, now)
	if err != nil {
		return IngestGscResponse{}, err
	}
	insJSON, _ := json.Marshal(inspections)
	out.Inspected = int64(len(inspections))
	out.InspectionsJSON = string(insJSON)
	return out, nil
}

// siteProperty resolves the Search Console property: the site row's
// gsc_site_url wins, otherwise the site blog.yaml baseURL as a URL-prefix
// property. The loaded Blog rides back for the inspection candidate walk.
func siteProperty(repoPath, siteURL string) (string, *blogyaml.Blog, error) {
	if repoPath == "" {
		return "", nil, errors.New("site repo_path is not set")
	}
	b, diags, err := blogyaml.Load(filepath.Join(repoPath, "blog.yaml"))
	if err != nil {
		return "", nil, err
	}
	if len(diags) > 0 {
		return "", nil, fmt.Errorf("blog.yaml invalid: %s", diags[0].String())
	}
	if siteURL != "" {
		return siteURL, b, nil
	}
	site := b.Site.BaseURL
	if !strings.HasSuffix(site, "/") {
		site += "/"
	}
	return site, b, nil
}

// inspectRecent runs the opt-in URL Inspection pass: articles whose
// repository-parsed lastmod is within GSC_INSPECT_RECENT_DAYS, capped at
// GSC_INSPECT_MAX per round (quota guard).
func inspectRecent(base, token, site, repoPath string, b *blogyaml.Blog, now time.Time) ([]gsc.Inspection, error) {
	entries, err := pcontent.IndexRepo(repoPath)
	if err != nil {
		return nil, err
	}
	pages := make([]gsc.PageMod, 0, len(entries))
	for _, e := range entries {
		pages = append(pages, gsc.PageMod{URL: e.URL, Lastmod: e.Lastmod})
	}
	urls := gsc.RecentURLs(pages, now, intEnv("GSC_INSPECT_RECENT_DAYS", 7), intEnv("GSC_INSPECT_MAX", 10))
	inspections, err := gsc.Inspect(base, token, site, urls)
	if err != nil {
		return nil, err
	}
	if inspections == nil {
		inspections = []gsc.Inspection{}
	}
	return inspections, nil
}

// envOr returns the environment value of key, or def when unset/empty.
func envOr(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

// intEnv reads a non-negative integer env with a default.
func intEnv(key string, def int) int {
	if v := os.Getenv(key); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n >= 0 {
			return n
		}
	}
	return def
}
