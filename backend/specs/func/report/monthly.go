package report

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
	pqueueio "github.com/park-jun-woo/abloq/pkg/queueio"
	"github.com/park-jun-woo/abloq/pkg/visibility/cflog"
	"github.com/park-jun-woo/abloq/pkg/visibility/priority"
	preport "github.com/park-jun-woo/abloq/pkg/visibility/report"
)

// @func monthly
// @error 500
// @description Assemble the monthly visibility report (crawl x index x citation joined per article, month-over-month trend, unknown bots, queue intake) from the ym-anchored DB aggregates riding in as JSON scalars, render markdown + JSON via pkg/visibility/report (deterministic — no clock value in the body) and publish the markdown copy as reports/<ym>.md to the blog repository through pkg/queueio (QUEUE_EXPORT_* env, dedicated -reports work clone, idempotent no-op commit); the DB row stays the lookup truth, the git commit is the publication copy

type MonthlyRequest struct {
	Ym            string
	PostsJSON     string
	BotsJSON      string
	PrevBotsJSON  string
	GscJSON       string
	PrevGscJSON   string
	CitesJSON     string
	PrevCitesJSON string
	QueueJSON     string
	UnknownJSON   string
}

type MonthlyResponse struct {
	Ym         string
	Markdown   string
	ReportJSON string
	Articles   int64
	Published  bool
}

// Monthly is the thin @call wrapper around pkg/visibility/report.Build: JSON
// translation, the default-ym resolution (the same last-closed-month (UTC)
// definition the window queries apply to ym=”) and the git publication.
// The page->article attribution uses the repository URL map and the weights
// come from blog.yaml — both single sources live outside the database.
func Monthly(req MonthlyRequest) (MonthlyResponse, error) {
	root := os.Getenv("BLOG_REPO_PATH")
	if root == "" {
		return MonthlyResponse{}, errors.New("BLOG_REPO_PATH is not set")
	}
	b, diags, err := blogyaml.Load(filepath.Join(root, "blog.yaml"))
	if err != nil {
		return MonthlyResponse{}, err
	}
	if len(diags) > 0 {
		return MonthlyResponse{}, fmt.Errorf("blog.yaml invalid: %s", diags[0].String())
	}
	ym, err := preport.ResolveYM(req.Ym, time.Now().UTC())
	if err != nil {
		return MonthlyResponse{}, err
	}
	urls, err := cflog.BuildURLMap(root, b)
	if err != nil {
		return MonthlyResponse{}, err
	}
	in := preport.Input{YM: ym, URLs: urls, Weights: priority.WeightsOf(b.Geo.PriorityWeights)}
	if err := decodeInputs(req, &in); err != nil {
		return MonthlyResponse{}, err
	}
	r := preport.Build(in)
	md := preport.Markdown(r)
	cfg, err := pqueueio.PublishConfigFromEnv()
	if err != nil {
		return MonthlyResponse{}, err
	}
	published, err := pqueueio.PublishFile(cfg, "reports/"+ym+".md", []byte(md))
	if err != nil {
		return MonthlyResponse{}, err
	}
	return MonthlyResponse{
		Ym:         ym,
		Markdown:   md,
		ReportJSON: string(preport.JSON(r)),
		Articles:   int64(len(r.Rows)),
		Published:  published,
	}, nil
}

func decodeInputs(req MonthlyRequest, in *preport.Input) error {
	fields := []struct {
		data string
		dst  any
	}{
		{req.PostsJSON, &in.Posts},
		{req.BotsJSON, &in.Bots},
		{req.PrevBotsJSON, &in.PrevBots},
		{req.GscJSON, &in.Pages},
		{req.PrevGscJSON, &in.PrevPages},
		{req.CitesJSON, &in.Cites},
		{req.PrevCitesJSON, &in.PrevCites},
		{req.QueueJSON, &in.Queue},
		{req.UnknownJSON, &in.UnknownBots},
	}
	for _, f := range fields {
		if err := json.Unmarshal([]byte(f.data), f.dst); err != nil {
			return err
		}
	}
	return nil
}
