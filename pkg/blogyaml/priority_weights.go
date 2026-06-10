//ff:type feature=blogyaml type=schema
//ff:what blog.yaml geo.priority_weights — 측정 신호 가중 계수(fetcher/train/gsc/citation), 기본값은 fetcher 최고 가중 (Phase014)
package blogyaml

// PriorityWeights are the measurement-signal coefficients of the Phase014
// priority scorer. Defaults keep fetcher the highest weight: a
// user-triggered fetch is the strongest consumption evidence (§6.1).
// Citation feeds the priority score only — §6.3's single allowed use.
type PriorityWeights struct {
	Fetcher  int64 `yaml:"fetcher" json:"fetcher"`
	Train    int64 `yaml:"train" json:"train"`
	GSC      int64 `yaml:"gsc" json:"gsc"`
	Citation int64 `yaml:"citation" json:"citation"`
}
