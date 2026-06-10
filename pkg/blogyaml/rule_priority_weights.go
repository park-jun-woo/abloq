//ff:func feature=blogyaml type=rule control=iteration dimension=1
//ff:what [priority-weights-range] geo.priority_weights 범위 검증 — fetcher/train/gsc/citation 전부 0 이상 (Phase014)
package blogyaml

import "fmt"

// rulePriorityWeights validates the priority_weights coefficients: every
// weight must be >= 0 (a zero weight legitimately silences its signal).
func rulePriorityWeights(filename string, b *Blog, idx lineIndex) []Diagnostic {
	var diags []Diagnostic
	weights := []struct {
		key   string
		value int64
	}{
		{"fetcher", b.Geo.PriorityWeights.Fetcher},
		{"train", b.Geo.PriorityWeights.Train},
		{"gsc", b.Geo.PriorityWeights.GSC},
		{"citation", b.Geo.PriorityWeights.Citation},
	}
	for _, w := range weights {
		if w.value < 0 {
			diags = append(diags, Diagnostic{
				File: filename, Line: lineOf(idx, "geo.priority_weights."+w.key), Rule: "priority-weights-range",
				Message: fmt.Sprintf("geo.priority_weights.%s must be >= 0 (got %d)", w.key, w.value),
			})
		}
	}
	return diags
}
