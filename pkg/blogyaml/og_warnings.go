//ff:func feature=blogyaml type=rule control=sequence
//ff:what [og-local-variants] 경고 진단 — provider local(미선언 포함)인데 variants가 선언된 무의미 조합, 차단하지 않는다
//ff:why blogyaml.Validate는 전부 차단 진단이라 경고 채널이 없다 — 비차단 경고는 별도 함수로 분리해 validate 출력에만 얹는다 (Phase022)
package blogyaml

// OGWarnings returns non-blocking advisory diagnostics for the image.og
// block: variants without an AI provider are dead declarations. Callers print
// these without affecting the exit code.
func OGWarnings(filename string, b *Blog, idx lineIndex) []Diagnostic {
	if b == nil || len(b.Image.OG.Variants) == 0 || b.Image.OGProvider() != "local" {
		return nil
	}
	return []Diagnostic{{
		File: filename, Line: lineOf(idx, "image.og.variants"), Rule: "og-local-variants",
		Message: "image.og.variants are declared but provider is local — AI variant presets have no effect (warning)",
	}}
}
