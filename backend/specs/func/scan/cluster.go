package scan

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
	pqueueio "github.com/park-jun-woo/abloq/pkg/queueio"
	pcluster "github.com/park-jun-woo/abloq/pkg/scan/cluster"
)

// @func cluster
// @error 500
// @description Run one cluster scan over the site's blog repository at repo_path: build the default-language tag/internal-link graph (bodies parsed directly — posts only stores link counts), detect the four cluster violations (tag-taxonomy when geo.taxonomy is declared, no-orphan-tag, min-internal-links from geo.min_internal_links, no-isolated-post) and return one kind=cluster queue candidate per violating article with deterministically ranked link suggestions (shared tags, same section, date proximity, section/slug tie break). The abloq CLI shares the same pkg statelessly (`abloq scan cluster`)

type ClusterRequest struct {
	RepoPath string
}

type ClusterResponse struct {
	ItemsJSON []byte
	Detected  int64
}

// Cluster is the thin @call wrapper around pkg/scan/cluster.Scan: the
// repository root rides in from the site row (multisite — the handler
// injects sites.repo_path) and the detection thresholds from its blog.yaml
// — the graph, the violations and the candidate ranking all live in the
// pkg, so the CLI and this endpoint can never diverge.
func Cluster(req ClusterRequest) (ClusterResponse, error) {
	if req.RepoPath == "" {
		return ClusterResponse{}, errors.New("site repo_path is not set")
	}
	b, diags, err := blogyaml.Load(filepath.Join(req.RepoPath, "blog.yaml"))
	if err != nil {
		return ClusterResponse{}, err
	}
	if len(diags) > 0 {
		return ClusterResponse{}, fmt.Errorf("blog.yaml invalid: %s", diags[0].String())
	}
	items := pcluster.Scan(req.RepoPath, b)
	return ClusterResponse{ItemsJSON: pqueueio.EncodeRows(items), Detected: int64(len(items))}, nil
}
