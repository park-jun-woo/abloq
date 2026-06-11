//ff:func feature=blogyaml type=rule control=sequence
//ff:what [llmstxt-max-summary] geo.llms_txt.max_summary가 0 이상(0 = 무제한)인지 검증
package blogyaml

import "fmt"

// ruleLlmsTxtMaxSummary requires max_summary to be >= 0 (0 = unlimited).
func ruleLlmsTxtMaxSummary(filename string, b *Blog, idx lineIndex) []Diagnostic {
	if b.Geo.LlmsTxt.MaxSummary >= 0 {
		return nil
	}
	return []Diagnostic{{
		File: filename, Line: llmsTxtLine(idx, "max_summary"), Rule: "llmstxt-max-summary",
		Message: fmt.Sprintf("geo.llms_txt.max_summary must be >= 0 (got %d)", b.Geo.LlmsTxt.MaxSummary),
	}}
}
