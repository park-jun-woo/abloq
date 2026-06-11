package queueio

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	pqueueio "github.com/park-jun-woo/abloq/pkg/queueio"
)

// @func export
// @error 500
// @description Run one queue export cycle on the site's dedicated work clone (<QUEUE_EXPORT_WORKDIR>/<site>, repo URL and commit author from the site row): pull, detect agent-deleted files (consumed), write open items as quests/queue/*.yaml and push — git work, idempotent no-op commits and the deterministic serialization live in pkg/queueio, which Phase010/011 scanners reuse. An empty site repo URL keeps the established "unset = 500" per site

type ExportRequest struct {
	SiteName     string
	RepoURL      string
	Author       string
	AuthorEmail  string
	OpenJSON     string
	ExportedJSON string
}

type ExportResponse struct {
	ExportedIdsJSON []byte
	ConsumedIdsJSON []byte
	Exported        int64
	Consumed        int64
}

// Export is the thin @call wrapper around pkg/queueio.Export: JSON
// translation only. Configuration (repo URL, commit author) comes from the
// site row so no path or credential flows through the API surface; the work
// clone lives at <base>/<site> (per-site isolation — two sites must never
// race inside one checkout) and the deploy key rides on the global
// GIT_SSH_COMMAND (one key for N repos, v1).
func Export(req ExportRequest) (ExportResponse, error) {
	if req.SiteName == "" {
		return ExportResponse{}, errors.New("site name is not set")
	}
	cfg, err := pqueueio.NewConfig(req.RepoURL, filepath.Join(workdirBase(), req.SiteName),
		req.Author, req.AuthorEmail)
	if err != nil {
		return ExportResponse{}, err
	}
	open, err := pqueueio.DecodeRows([]byte(req.OpenJSON))
	if err != nil {
		return ExportResponse{}, err
	}
	exported, err := pqueueio.DecodeRows([]byte(req.ExportedJSON))
	if err != nil {
		return ExportResponse{}, err
	}
	res, err := pqueueio.Export(cfg, open, exported)
	if err != nil {
		return ExportResponse{}, err
	}
	// Marshal cannot fail: both inputs are []int64 (never nil — see rowIDs).
	exportedIDs, _ := json.Marshal(res.ExportedIDs)
	consumedIDs, _ := json.Marshal(res.ConsumedIDs)
	return ExportResponse{
		ExportedIdsJSON: exportedIDs,
		ConsumedIdsJSON: consumedIDs,
		Exported:        int64(len(res.ExportedIDs)),
		Consumed:        int64(len(res.ConsumedIDs)),
	}, nil
}

// workdirBase resolves the work-clone base directory: QUEUE_EXPORT_WORKDIR
// when set (test harnesses point it at a temp dir), otherwise the image
// default. Each site clones into its own <base>/<site> subdirectory.
func workdirBase() string {
	if v := os.Getenv("QUEUE_EXPORT_WORKDIR"); v != "" {
		return v
	}
	return "/var/lib/abloqd/queue-export"
}
