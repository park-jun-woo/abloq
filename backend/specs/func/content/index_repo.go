package content

import (
	"encoding/json"
	"errors"
	"os"

	pcontent "github.com/park-jun-woo/abloq/pkg/content"
)

// @func indexRepo
// @error 500
// @description Parse every article front matter under BLOG_REPO_PATH (blog.yaml SSOT drives languages, sections and the URL contract) and return the posts index as JSON

type IndexRepoRequest struct{}

type IndexRepoResponse struct {
	EntriesJSON []byte
	Count       int64
}

// IndexRepo is the thin @call wrapper around pkg/content.IndexRepo: the blog
// repository root comes from the BLOG_REPO_PATH environment variable so that
// no path credential ever flows through the API surface.
func IndexRepo(req IndexRepoRequest) (IndexRepoResponse, error) {
	root := os.Getenv("BLOG_REPO_PATH")
	if root == "" {
		return IndexRepoResponse{}, errors.New("BLOG_REPO_PATH is not set")
	}
	entries, err := pcontent.IndexRepo(root)
	if err != nil {
		return IndexRepoResponse{}, err
	}
	data, err := json.Marshal(entries)
	if err != nil {
		return IndexRepoResponse{}, err
	}
	return IndexRepoResponse{EntriesJSON: data, Count: int64(len(entries))}, nil
}
