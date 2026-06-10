//ff:func feature=blogyaml type=rule control=iteration dimension=1
//ff:what [lang-bcp47] languages가 비어있지 않고 각 항목이 유효한 BCP-47 언어 코드인지 검증
package blogyaml

import (
	"fmt"

	"golang.org/x/text/language"
)

// ruleLangBCP47 reports empty languages or entries that are not valid BCP-47 tags.
func ruleLangBCP47(filename string, b *Blog, idx lineIndex) []Diagnostic {
	if len(b.Languages) == 0 {
		return []Diagnostic{{
			File: filename, Line: lineOf(idx, "languages"), Rule: "lang-bcp47",
			Message: "languages must contain at least one BCP-47 language code (first entry = default language)",
		}}
	}
	var diags []Diagnostic
	for i, lang := range b.Languages {
		if _, err := language.Parse(lang); err != nil {
			diags = append(diags, Diagnostic{
				File: filename, Line: lineOf(idx, fmt.Sprintf("languages[%d]", i)), Rule: "lang-bcp47",
				Message: fmt.Sprintf("languages[%d] %q is not a valid BCP-47 language code", i, lang),
			})
		}
	}
	return diags
}
