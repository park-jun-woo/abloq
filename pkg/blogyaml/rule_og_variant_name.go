//ff:func feature=blogyaml type=rule control=iteration dimension=1
//ff:what [og-variant-name] 각 variant name이 URL-safe·유니크·예약어 default 아님을 검증 — 드래프트 파일명 {name}-{n}.webp의 전제
package blogyaml

import "fmt"

// ruleOGVariantName validates every declared variant name: it lands in draft
// filenames, so it must be URL-safe, unique, and must not shadow the reserved
// "default" (the implicit default-settings candidate).
func ruleOGVariantName(filename string, b *Blog, idx lineIndex) []Diagnostic {
	var diags []Diagnostic
	seen := map[string]bool{}
	for i, v := range b.Image.OG.Variants {
		line := lineOf(idx, fmt.Sprintf("image.og.variants[%d].name", i))
		switch {
		case !ogNameSafe(v.Name):
			diags = append(diags, Diagnostic{File: filename, Line: line, Rule: "og-variant-name",
				Message: fmt.Sprintf("image.og.variants[%d] name %q must be URL-safe (lowercase letters, digits, '-', '_')", i, v.Name)})
		case v.Name == "default":
			diags = append(diags, Diagnostic{File: filename, Line: line, Rule: "og-variant-name",
				Message: fmt.Sprintf("image.og.variants[%d] name \"default\" is reserved for the default-settings candidate", i)})
		case seen[v.Name]:
			diags = append(diags, Diagnostic{File: filename, Line: line, Rule: "og-variant-name",
				Message: fmt.Sprintf("image.og.variants[%d] name %q duplicates an earlier variant", i, v.Name)})
		}
		seen[v.Name] = true
	}
	return diags
}
