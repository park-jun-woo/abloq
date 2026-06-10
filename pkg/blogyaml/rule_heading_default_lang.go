//ff:func feature=blogyaml type=rule control=iteration dimension=1
//ff:what [heading-default-lang] structure.headings의 각 헤딩 키에 기본 언어(languages[0]) 항목이 존재하는지 검증
package blogyaml

import "fmt"

// ruleHeadingDefaultLang requires every headings entry to cover the default language.
func ruleHeadingDefaultLang(filename string, b *Blog, idx lineIndex) []Diagnostic {
	if len(b.Languages) == 0 {
		return nil // reported by lang-bcp47
	}
	def := b.Languages[0]
	var diags []Diagnostic
	for _, key := range sortedKeys(b.Structure.Headings) {
		if _, ok := b.Structure.Headings[key][def]; !ok {
			diags = append(diags, Diagnostic{
				File: filename, Line: lineOf(idx, "structure.headings."+key), Rule: "heading-default-lang",
				Message: fmt.Sprintf("structure.headings.%s is missing the default language %q", key, def),
			})
		}
	}
	return diags
}
