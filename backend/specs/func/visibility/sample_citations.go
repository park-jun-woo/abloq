package visibility

import (
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
	"github.com/park-jun-woo/abloq/pkg/visibility/citation"
)

// @func sampleCitations
// @error 500
// @description Run one citation-sampling round for one site: the active query set rides in as JSON (id order), every engine whose API key env is set (PERPLEXITY_API_KEY / OPENAI_API_KEY / ANTHROPIC_API_KEY, base URLs env-overridable) answers the first geo.citation_budget queries (blog.yaml under the site's repo_path — 0 disables sampling), and the own-domain citations (baseURL substring match, extractor v1) come back as the citation_samples batch payload. Engine calls are throttled by CITATION_INTERVAL_MS (default 1000). The abloq CLI shares the same pkg statelessly (`abloq sample citations --queries <file>`)

type SampleCitationsRequest struct {
	RepoPath    string
	QueriesJSON string
}

type SampleCitationsResponse struct {
	SamplesJSON []byte
	Engines     int64
	Sampled     int64
}

// SampleCitations is the thin @call wrapper around pkg/visibility/citation.Run:
// JSON translation plus the env-configured engines and the site's blog.yaml
// budget and domain (the engine keys stay instance-global by design — v1
// shares one account across sites; the budget caps each site separately).
func SampleCitations(req SampleCitationsRequest) (SampleCitationsResponse, error) {
	if req.RepoPath == "" {
		return SampleCitationsResponse{}, errors.New("site repo_path is not set")
	}
	b, diags, err := blogyaml.Load(filepath.Join(req.RepoPath, "blog.yaml"))
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
