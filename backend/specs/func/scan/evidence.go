package scan

import (
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
	pqueueio "github.com/park-jun-woo/abloq/pkg/queueio"
	pevidence "github.com/park-jun-woo/abloq/pkg/scan/evidence"
)

// @func evidence
// @error 500
// @description Run one evidence scan over the site's blog repository at repo_path: detect unsourced numeric claims (gate detector, claims_ignore honored) and probe every citation URL (HEAD/GET, concurrency + per-domain rate limit inside pkg/scan/evidence, LINKCHECK_HOST_OVERRIDE for stubs) — rot confirms only at 3 consecutive failed scans, computed from the previous citation_checks state riding in as JSON; returns the updated check state, the kind=evidence queue candidates and their count. The abloq CLI shares the same pkg statelessly (`abloq scan evidence`)

type EvidenceRequest struct {
	RepoPath   string
	ChecksJSON string
}

type EvidenceResponse struct {
	ChecksJSON []byte
	ItemsJSON  []byte
	Detected   int64
}

// Evidence is the thin @call wrapper around pkg/scan/evidence.Scan: JSON
// translation plus the repository root from the site row (multisite — the
// handler injects sites.repo_path). The check JSON field names mirror the
// citation_checks columns, so the backend's jsonb_agg supply and the pkg's
// output feed the same upsert.
func Evidence(req EvidenceRequest) (EvidenceResponse, error) {
	if req.RepoPath == "" {
		return EvidenceResponse{}, errors.New("site repo_path is not set")
	}
	b, diags, err := blogyaml.Load(filepath.Join(req.RepoPath, "blog.yaml"))
	if err != nil {
		return EvidenceResponse{}, err
	}
	if len(diags) > 0 {
		return EvidenceResponse{}, fmt.Errorf("blog.yaml invalid: %s", diags[0].String())
	}
	var prev []pevidence.Check
	if err := json.Unmarshal([]byte(req.ChecksJSON), &prev); err != nil {
		return EvidenceResponse{}, err
	}
	res := pevidence.Scan(req.RepoPath, b, prev, pevidence.NewChecker())
	// Marshal cannot fail: Check is plain string/int64 fields, never nil slice.
	checksJSON, _ := json.Marshal(res.Checks)
	return EvidenceResponse{
		ChecksJSON: checksJSON,
		ItemsJSON:  pqueueio.EncodeRows(res.Items),
		Detected:   int64(len(res.Items)),
	}, nil
}
