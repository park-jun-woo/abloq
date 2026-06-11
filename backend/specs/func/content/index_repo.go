package content

import (
	"encoding/json"
	"errors"

	pcontent "github.com/park-jun-woo/abloq/pkg/content"
)

// @func indexRepo
// @error 500
// @description Parse every article front matter under the site's repo_path (blog.yaml SSOT drives languages, sections and the URL contract) and return the posts index as JSON

type IndexRepoRequest struct {
	RepoPath string
}

type IndexRepoResponse struct {
	EntriesJSON []byte
	Count       int64
}

// IndexRepo is the thin @call wrapper around pkg/content.IndexRepo: the blog
// repository root rides in from the site row (multisite — the handler
// injects sites.repo_path), so no path credential ever flows through the
// API surface and no instance-global env couples the sites together.
func IndexRepo(req IndexRepoRequest) (IndexRepoResponse, error) {
	if req.RepoPath == "" {
		return IndexRepoResponse{}, errors.New("site repo_path is not set")
	}
	entries, err := pcontent.IndexRepo(req.RepoPath)
	if err != nil {
		return IndexRepoResponse{}, err
	}
	data, err := json.Marshal(entries)
	if err != nil {
		return IndexRepoResponse{}, err
	}
	return IndexRepoResponse{EntriesJSON: data, Count: int64(len(entries))}, nil
}
