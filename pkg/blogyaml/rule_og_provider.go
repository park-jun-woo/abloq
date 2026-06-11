//ff:func feature=blogyaml type=rule control=sequence
//ff:what [og-provider] image.og provider가 local|gemini enum인지 검증 — 빈 값(미선언)은 local로 통과
package blogyaml

import "fmt"

// ruleOGProvider requires the effective OG provider to be a known one.
// Adding a provider implementation extends this enum in the same change.
func ruleOGProvider(filename string, b *Blog, idx lineIndex) []Diagnostic {
	p := b.Image.OGProvider()
	if p == "local" || p == "gemini" {
		return nil
	}
	return []Diagnostic{{
		File: filename, Line: lineOf(idx, "image.og.provider"), Rule: "og-provider",
		Message: fmt.Sprintf("image.og provider %q must be one of: local, gemini", p),
	}}
}
