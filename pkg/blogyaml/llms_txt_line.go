//ff:func feature=blogyaml type=parser control=sequence topic=diagnostics
//ff:what geo.llms_txt 하위 키의 진단 라인 조회 — 객체 폼 키가 없으면(단축형 등) geo.llms_txt 키 라인으로 폴백
package blogyaml

// llmsTxtLine resolves the diagnostic line for a geo.llms_txt sub-key,
// falling back to the geo.llms_txt key itself for the string shorthand.
func llmsTxtLine(idx lineIndex, key string) int {
	if line, ok := idx["geo.llms_txt."+key]; ok {
		return line
	}
	return lineOf(idx, "geo.llms_txt")
}
