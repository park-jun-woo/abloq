//ff:type feature=blogyaml type=schema
//ff:what blog.yaml geo 섹션 — crawlers 정책, llms_txt, jsonld, 게이트 임계값(freshness_days/min_sources/min_internal_links/min_meaningful_diff), 선택적 taxonomy, citation_budget(샘플링 비용 상한)
package blogyaml

// Geo declares GEO (Generative Engine Optimization) policy and gate thresholds.
// Crawlers maps crawler category or bot name -> "allow" | "block".
// MinMeaningfulDiff is the honest-lastmod token-diff threshold (normalized tokens).
// Taxonomy is the optional curated tag vocabulary — when absent, the cluster
// scanner's tag-taxonomy check is silently skipped (Phase011).
// CitationBudget caps one citation-sampling round: the maximum number of
// queries per engine (0 = sampling disabled, the opt-in default — Phase013).
type Geo struct {
	Crawlers          map[string]string `yaml:"crawlers" json:"crawlers"`
	LlmsTxt           string            `yaml:"llms_txt" json:"llms_txt"`
	JSONLD            []string          `yaml:"jsonld" json:"jsonld"`
	FreshnessDays     int               `yaml:"freshness_days" json:"freshness_days"`
	MinSources        int               `yaml:"min_sources" json:"min_sources"`
	MinInternalLinks  int               `yaml:"min_internal_links" json:"min_internal_links"`
	MinMeaningfulDiff int               `yaml:"min_meaningful_diff" json:"min_meaningful_diff"`
	CitationBudget    int               `yaml:"citation_budget" json:"citation_budget"`
	Taxonomy          []string          `yaml:"taxonomy" json:"taxonomy"`
}
