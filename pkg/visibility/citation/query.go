//ff:type feature=visibility type=schema topic=citation
//ff:what 샘플링 질의 1건 — citation_queries id와 질의문 (JSON 키 = 컬럼명, 백엔드 jsonb_agg와 CLI 파일이 같은 형식)
package citation

// Query is one standard sampling query. The backend supplies it from the
// citation_queries table (jsonb_agg), the CLI from a --queries file — both
// feed the same runner.
type Query struct {
	ID   int64  `json:"id" yaml:"id"`
	Text string `json:"query_text" yaml:"query_text"`
}
