//ff:func feature=blogyaml type=rule control=iteration dimension=1
//ff:what [llmstxt-languages] geo.llms_txt.languages가 base|all 또는 선언된 languages의 부분집합인지 검증
package blogyaml

import "fmt"

// ruleLlmsTxtLanguages requires the llms.txt language scope to be "base",
// "all", or an explicit subset of the declared languages.
func ruleLlmsTxtLanguages(filename string, b *Blog, idx lineIndex) []Diagnostic {
	langs := b.Geo.LlmsTxt.Languages
	if len(langs) == 1 && (langs[0] == "base" || langs[0] == "all") {
		return nil
	}
	declared := map[string]bool{}
	for _, l := range b.Languages {
		declared[l] = true
	}
	var diags []Diagnostic
	for i, l := range langs {
		if declared[l] {
			continue
		}
		diags = append(diags, Diagnostic{
			File: filename, Line: llmsTxtItemLine(idx, fmt.Sprintf("languages[%d]", i), "languages"),
			Rule:    "llmstxt-languages",
			Message: fmt.Sprintf("geo.llms_txt.languages[%d] %q must be base, all, or a declared language", i, l),
		})
	}
	return diags
}
