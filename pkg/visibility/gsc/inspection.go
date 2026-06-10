//ff:type feature=visibility type=schema topic=gsc
//ff:what URL Inspection 응답 요약 1건 — URL·verdict·coverage state (적재 없이 응답으로만 반환)
package gsc

// Inspection is one URL Inspection verdict summary. v1 stores nothing: the
// index-state time series is already carried by impressions/clicks — this is
// an on-demand operator readout (the API quota is small).
type Inspection struct {
	URL           string `json:"url"`
	Verdict       string `json:"verdict"`
	CoverageState string `json:"coverage_state"`
}
