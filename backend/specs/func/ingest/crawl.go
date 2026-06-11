package ingest

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
	"github.com/park-jun-woo/abloq/pkg/visibility/cflog"
)

// @func crawl
// @error 500
// @description Run one incremental CloudFront crawl ingest for one site: open the site's cf_log_source (local directory or s3://bucket/prefix with stdlib sigv4 over the AWS_* env credentials), build the URI reverse map from the site's repository at repo_path (Blog.URLLang page paths plus the .md parallel-served twins), and aggregate only the closed hours past the cursor — the cursor is a time boundary ("YYYY-MM-DD-HH" UTC), never a file key, and hours inside the safety margin (CF_LOG_MARGIN_HOURS, default 2) stay untouched so late-delivered files are not lost. Returns the crawl_hits/ingest_cursors/unknown_bots batch payloads plus the processed file count and the html+md hit total. The abloq CLI shares the same pkg statelessly (`abloq ingest crawl`)

type CrawlRequest struct {
	RepoPath    string
	LogSource   string
	CursorsJSON string
}

type CrawlResponse struct {
	HitsJSON    []byte
	CursorsJSON []byte
	UnknownJSON []byte
	Files       int64
	Hits        int64
}

// Crawl is the thin @call wrapper around pkg/visibility/cflog.Collect: JSON
// translation plus the site-row-configured source and repository (multisite
// — the handler injects sites.cf_log_source and sites.repo_path; only the
// safety margin stays a global env). The row JSON field names mirror the
// crawl_hits/ingest_cursors/unknown_bots columns, so the backend's
// jsonb_agg supply and the pkg's output feed the same upserts.
func Crawl(req CrawlRequest) (CrawlResponse, error) {
	if req.LogSource == "" {
		return CrawlResponse{}, errors.New("site cf_log_source is not set")
	}
	if req.RepoPath == "" {
		return CrawlResponse{}, errors.New("site repo_path is not set")
	}
	b, diags, err := blogyaml.Load(filepath.Join(req.RepoPath, "blog.yaml"))
	if err != nil {
		return CrawlResponse{}, err
	}
	if len(diags) > 0 {
		return CrawlResponse{}, fmt.Errorf("blog.yaml invalid: %s", diags[0].String())
	}
	urls, err := cflog.BuildURLMap(req.RepoPath, b)
	if err != nil {
		return CrawlResponse{}, err
	}
	src, err := cflog.OpenSource(req.LogSource)
	if err != nil {
		return CrawlResponse{}, err
	}
	var cursors []cflog.Cursor
	if err := json.Unmarshal([]byte(req.CursorsJSON), &cursors); err != nil {
		return CrawlResponse{}, err
	}
	res, err := cflog.Collect(src, urls, cursors, time.Now().UTC(), marginHours())
	if err != nil {
		return CrawlResponse{}, err
	}
	// Marshal cannot fail: the row types are plain string/int64 fields.
	hitsJSON, _ := json.Marshal(res.Hits)
	cursorsJSON, _ := json.Marshal(res.Cursors)
	unknownJSON, _ := json.Marshal(res.Unknown)
	return CrawlResponse{
		HitsJSON:    hitsJSON,
		CursorsJSON: cursorsJSON,
		UnknownJSON: unknownJSON,
		Files:       res.Files,
		Hits:        res.Total,
	}, nil
}

// marginHours reads the late-delivery safety margin from CF_LOG_MARGIN_HOURS
// (default 2 — CF delivery is rarely up to 24h late; raise the env when that
// bites).
func marginHours() time.Duration {
	if v := os.Getenv("CF_LOG_MARGIN_HOURS"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n >= 0 {
			return time.Duration(n) * time.Hour
		}
	}
	return 2 * time.Hour
}
