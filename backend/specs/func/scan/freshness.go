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
	"github.com/park-jun-woo/abloq/pkg/visibility/priority"
)

// @func freshness
// @error 500
// @description Detect posts whose lastmod exceeded geo.freshness_days (blog.yaml under BLOG_REPO_PATH) and return prioritized refresh queue candidates as batch-insert JSON — detection, cold-start priority and serialization live in pkg/scan/freshness + pkg/queueio, which the abloq CLI shares (`abloq scan freshness`)

type FreshnessRequest struct {
	PostsJSON string
	HitsJSON  string
}

type FreshnessResponse struct {
	ItemsJSON []byte
	Detected  int64
}

// Freshness is the thin @call wrapper around pkg/scan/freshness.Scan: JSON
// translation plus the freshness threshold from the mounted blog.yaml. The
// posts JSON field names mirror pkg/content.Entry tags, so the backend's
// jsonb_agg supply and the CLI's direct repository parse feed the same logic.
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
	items := pfreshness.Scan(entries, pfreshness.HitsMap(sums), b.Geo.FreshnessDays, time.Now().UTC(), priority.ColdStart{})
	return FreshnessResponse{ItemsJSON: pqueueio.EncodeRows(items), Detected: int64(len(items))}, nil
}
