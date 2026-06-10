//ff:type feature=blogyaml type=schema
//ff:what blog.yaml geo 섹션 — crawlers 정책, llms_txt, jsonld, 게이트 임계값(freshness_days/min_sources/min_internal_links)
package blogyaml

// Geo declares GEO (Generative Engine Optimization) policy and gate thresholds.
// Crawlers maps crawler category or bot name -> "allow" | "block".
type Geo struct {
	Crawlers         map[string]string `yaml:"crawlers" json:"crawlers"`
	LlmsTxt          string            `yaml:"llms_txt" json:"llms_txt"`
	JSONLD           []string          `yaml:"jsonld" json:"jsonld"`
	FreshnessDays    int               `yaml:"freshness_days" json:"freshness_days"`
	MinSources       int               `yaml:"min_sources" json:"min_sources"`
	MinInternalLinks int               `yaml:"min_internal_links" json:"min_internal_links"`
}
