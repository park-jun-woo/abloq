//ff:type feature=visibility type=schema topic=citation
//ff:what 인용 샘플 1행 — 질의 id·엔진·인용 여부·근거 JSON·추출기 버전, citation_samples 1행과 1:1 (JSON 키 = 컬럼명)
//ff:why 추출 로직이 바뀌면 ExtractorVersion으로 시계열 단절을 표시한다 — 정밀도보다 추세 일관성 (Phase013)
package citation

// ExtractorVersion marks which citation-extraction heuristic produced a
// sample. Bump it whenever the matching logic changes so the time series
// shows the break instead of a fake trend shift.
const ExtractorVersion = "v1"

// Sample is one engine×query sampling record. JSON keys mirror the abloqd
// citation_samples table columns; Evidence is a JSON string (matched URLs or
// the engine error).
type Sample struct {
	CitationQueriesID int64  `json:"citation_queries_id"`
	Engine            string `json:"engine"`
	Cited             bool   `json:"cited"`
	Evidence          string `json:"evidence"`
	ExtractorVersion  string `json:"extractor_version"`
}
