package scan

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
	pcontent "github.com/park-jun-woo/abloq/pkg/content"
	pqueueio "github.com/park-jun-woo/abloq/pkg/queueio"
	pfreshness "github.com/park-jun-woo/abloq/pkg/scan/freshness"
	"github.com/park-jun-woo/abloq/pkg/visibility/cflog"
	"github.com/park-jun-woo/abloq/pkg/visibility/priority"
	"github.com/park-jun-woo/abloq/pkg/visibility/report"
)

// @func freshness
// @error 500
// @description Detect posts whose lastmod exceeded geo.freshness_days (blog.yaml under BLOG_REPO_PATH) and return prioritized refresh queue candidates as batch-insert JSON — detection lives in pkg/scan/freshness, the Composite scorer (measured 30d signals with cold-start fallback, weights from geo.priority_weights) in pkg/visibility/priority, the signal assembly (bot classification via pkg/bots, GSC page attribution via the repository URL map) in pkg/visibility/report; the abloq CLI shares the same detection with an empty signals map (`abloq scan freshness`)

type FreshnessRequest struct {
	PostsJSON string
	HitsJSON  string
	BotsJSON  string
	GscJSON   string
	CitesJSON string
}

type FreshnessResponse struct {
	ItemsJSON []byte
	Detected  int64
}

// Freshness is the thin @call wrapper around pkg/scan/freshness.Scan: JSON
// translation, the measured-signal assembly and the Composite scorer
// injection. The posts JSON field names mirror pkg/content.Entry tags, so
// the backend's jsonb_agg supply and the CLI's direct repository parse feed
// the same logic; bot classification (pkg/bots) and GSC page attribution
// (repository URL map) happen here because neither lives in the database.
func Freshness(req FreshnessRequest) (FreshnessResponse, error) {
	root := os.Getenv("BLOG_REPO_PATH")
	if root == "" {
		return FreshnessResponse{}, errors.New("BLOG_REPO_PATH is not set")
	}
	b, diags, err := blogyaml.Load(filepath.Join(root, "blog.yaml"))
	if err != nil {
		return FreshnessResponse{}, err
	}
	if len(diags) > 0 {
		return FreshnessResponse{}, fmt.Errorf("blog.yaml invalid: %s", diags[0].String())
	}
	var entries []pcontent.Entry
	if err := json.Unmarshal([]byte(req.PostsJSON), &entries); err != nil {
		return FreshnessResponse{}, err
	}
	var sums []pfreshness.HitSum
	if err := json.Unmarshal([]byte(req.HitsJSON), &sums); err != nil {
		return FreshnessResponse{}, err
	}
	var bots []report.BotSum
	if err := json.Unmarshal([]byte(req.BotsJSON), &bots); err != nil {
		return FreshnessResponse{}, err
	}
	var pages []report.PageSum
	if err := json.Unmarshal([]byte(req.GscJSON), &pages); err != nil {
		return FreshnessResponse{}, err
	}
	var cites []report.CiteSum
	if err := json.Unmarshal([]byte(req.CitesJSON), &cites); err != nil {
		return FreshnessResponse{}, err
	}
	urls, err := cflog.BuildURLMap(root, b)
	if err != nil {
		return FreshnessResponse{}, err
	}
	signals := report.MergeSignals(pfreshness.SignalsMap(sums),
		report.BotTotals(bots), report.PageTotals(pages, urls), report.CiteHits(cites))
	scorer := priority.Composite{W: priority.WeightsOf(b.Geo.PriorityWeights)}
	items := pfreshness.Scan(entries, signals, b.Languages, b.Geo.FreshnessDays, time.Now().UTC(), scorer)
	return FreshnessResponse{ItemsJSON: pqueueio.EncodeRows(items), Detected: int64(len(items))}, nil
}
