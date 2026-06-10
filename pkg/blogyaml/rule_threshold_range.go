//ff:func feature=blogyaml type=rule control=sequence
//ff:what [threshold-range] 게이트 임계값 범위 검증 — freshness_days ≥ 1, min_sources ≥ 0, min_internal_links ≥ 0, min_meaningful_diff ≥ 1
package blogyaml

import "fmt"

// ruleThresholdRange validates the numeric gate thresholds in geo.
func ruleThresholdRange(filename string, b *Blog, idx lineIndex) []Diagnostic {
	var diags []Diagnostic
	if b.Geo.FreshnessDays < 1 {
		diags = append(diags, Diagnostic{
			File: filename, Line: lineOf(idx, "geo.freshness_days"), Rule: "threshold-range",
			Message: fmt.Sprintf("geo.freshness_days must be >= 1 (got %d)", b.Geo.FreshnessDays),
		})
	}
	if b.Geo.MinSources < 0 {
		diags = append(diags, Diagnostic{
			File: filename, Line: lineOf(idx, "geo.min_sources"), Rule: "threshold-range",
			Message: fmt.Sprintf("geo.min_sources must be >= 0 (got %d)", b.Geo.MinSources),
		})
	}
	if b.Geo.MinInternalLinks < 0 {
		diags = append(diags, Diagnostic{
			File: filename, Line: lineOf(idx, "geo.min_internal_links"), Rule: "threshold-range",
			Message: fmt.Sprintf("geo.min_internal_links must be >= 0 (got %d)", b.Geo.MinInternalLinks),
		})
	}
	if b.Geo.MinMeaningfulDiff < 1 {
		diags = append(diags, Diagnostic{
			File: filename, Line: lineOf(idx, "geo.min_meaningful_diff"), Rule: "threshold-range",
			Message: fmt.Sprintf("geo.min_meaningful_diff must be >= 1 (got %d)", b.Geo.MinMeaningfulDiff),
		})
	}
	return diags
}
