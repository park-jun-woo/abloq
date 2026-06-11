//ff:func feature=blogyaml type=rule control=sequence
//ff:what [llmstxt-mode] geo.llms_txt mode가 auto|manual|off enum인지 검증 — 단축형 문자열도 같은 enum
package blogyaml

import "fmt"

// ruleLlmsTxtMode requires the effective llms.txt mode to be one of
// auto, manual or off (the string shorthand decodes into the same field).
func ruleLlmsTxtMode(filename string, b *Blog, idx lineIndex) []Diagnostic {
	mode := b.Geo.LlmsTxtMode()
	if mode == "auto" || mode == "manual" || mode == "off" {
		return nil
	}
	return []Diagnostic{{
		File: filename, Line: llmsTxtLine(idx, "mode"), Rule: "llmstxt-mode",
		Message: fmt.Sprintf("geo.llms_txt mode %q must be one of: auto, manual, off", mode),
	}}
}
