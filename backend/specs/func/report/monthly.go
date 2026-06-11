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
// @description Assemble the monthly visibility report for one site (crawl x index x citation joined per article, month-over-month trend, unknown bots, queue intake) from the ym-anchored site-scoped DB aggregates riding in as JSON scalars, render markdown + JSON via pkg/visibility/report (deterministic — no clock value in the body) and publish the markdown copy as reports/<ym>.md to the site's blog repository through pkg/queueio (repo URL and author from the site row, dedicated per-site -reports work clone, idempotent no-op commit); the DB row stays the lookup truth, the git commit is the publication copy

type MonthlyRequest struct {
	RepoPath      string
	SiteName      string
	RepoURL       string
	Author        string
	AuthorEmail   string
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
// The repository path and the publication repo URL/author ride in from the
// site row (multisite); the page->article attribution uses the repository
// URL map and the weights come from blog.yaml — both single sources live
// outside the database.
func Monthly(req MonthlyRequest) (MonthlyResponse, error) {
	if req.RepoPath == "" {
		return MonthlyResponse{}, errors.New("site repo_path is not set")
	}
	if req.SiteName == "" {
		return MonthlyResponse{}, errors.New("site name is not set")
	}
	b, diags, err := blogyaml.Load(filepath.Join(req.RepoPath, "blog.yaml"))
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
	urls, err := cflog.BuildURLMap(req.RepoPath, b)
	if err != nil {
		return MonthlyResponse{}, err
	}
	in := preport.Input{YM: ym, URLs: urls, Weights: priority.WeightsOf(b.Geo.PriorityWeights)}
	if err := decodeInputs(req, &in); err != nil {
		return MonthlyResponse{}, err
	}
	r := preport.Build(in)
	md := preport.Markdown(r)
	cfg, err := pqueueio.NewPublishConfig(req.RepoURL, filepath.Join(workdirBase(), req.SiteName),
		req.Author, req.AuthorEmail)
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

// workdirBase resolves the work-clone base directory: QUEUE_EXPORT_WORKDIR
// when set (test harnesses point it at a temp dir), otherwise the image
// default — the same convention as the queue exporter, NewPublishConfig
// then separates the publisher clone with the "-reports" suffix.
func workdirBase() string {
	if v := os.Getenv("QUEUE_EXPORT_WORKDIR"); v != "" {
		return v
	}
	return "/var/lib/abloqd/queue-export"
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
