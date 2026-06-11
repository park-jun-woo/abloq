//ff:func feature=blogyaml type=rule control=iteration dimension=1
//ff:what [llmstxt-labels] geo.llms_txt.section_labels의 각 키가 선언된 sections에 존재하는지 검증
package blogyaml

import "fmt"

// ruleLlmsTxtLabels requires every section_labels key to name a declared section.
func ruleLlmsTxtLabels(filename string, b *Blog, idx lineIndex) []Diagnostic {
	declared := map[string]bool{}
	for _, s := range b.Sections {
		declared[s] = true
	}
	var diags []Diagnostic
	for _, key := range sortedKeys(b.Geo.LlmsTxt.SectionLabels) {
		if declared[key] {
			continue
		}
		diags = append(diags, Diagnostic{
			File: filename, Line: llmsTxtItemLine(idx, "section_labels."+key, "section_labels"),
			Rule:    "llmstxt-labels",
			Message: fmt.Sprintf("geo.llms_txt.section_labels.%s is not a declared section", key),
		})
	}
	return diags
}
