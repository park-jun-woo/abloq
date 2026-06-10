package visibility

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
	"github.com/park-jun-woo/abloq/pkg/visibility/citation"
)

// @func sampleCitations
// @error 500
// @description Run one citation-sampling round: the active query set rides in as JSON (id order), every engine whose API key env is set (PERPLEXITY_API_KEY / OPENAI_API_KEY / ANTHROPIC_API_KEY, base URLs env-overridable) answers the first geo.citation_budget queries (blog.yaml under BLOG_REPO_PATH — 0 disables sampling), and the own-domain citations (baseURL substring match, extractor v1) come back as the citation_samples batch payload. Engine calls are throttled by CITATION_INTERVAL_MS (default 1000). The abloq CLI shares the same pkg statelessly (`abloq sample citations --queries <file>`)

type SampleCitationsRequest struct {
	QueriesJSON string
}

type SampleCitationsResponse struct {
	SamplesJSON []byte
	Engines     int64
	Sampled     int64
}

// SampleCitations is the thin @call wrapper around pkg/visibility/citation.Run:
// JSON translation plus the env-configured engines and the blog.yaml budget
// and domain. The sample JSON field names mirror the citation_samples
// columns, so the backend's batch insert and the pkg's output agree.
func SampleCitations(req SampleCitationsRequest) (SampleCitationsResponse, error) {
	root := os.Getenv("BLOG_REPO_PATH")
	if root == "" {
		return SampleCitationsResponse{}, errors.New("BLOG_REPO_PATH is not set")
	}
	b, diags, err := blogyaml.Load(filepath.Join(root, "blog.yaml"))
	if err != nil {
		return SampleCitationsResponse{}, err
	}
	if len(diags) > 0 {
		return SampleCitationsResponse{}, fmt.Errorf("blog.yaml invalid: %s", diags[0].String())
	}
	var queries []citation.Query
	if err := json.Unmarshal([]byte(req.QueriesJSON), &queries); err != nil {
		return SampleCitationsResponse{}, err
	}
	engines := citation.EnginesFromEnv()
	samples := citation.Run(engines, queries, b.Geo.CitationBudget, b.Site.BaseURL, citation.IntervalFromEnv())
	if samples == nil {
		samples = []citation.Sample{}
	}
	// Marshal cannot fail: the sample types are plain string/bool/int fields.
	samplesJSON, _ := json.Marshal(samples)
	return SampleCitationsResponse{
		SamplesJSON: samplesJSON,
		Engines:     int64(len(engines)),
		Sampled:     int64(len(samples)),
	}, nil
}
