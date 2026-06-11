//ff:func feature=blogyaml type=parser control=sequence topic=diagnostics
//ff:what geo.llms_txt 하위 항목(시퀀스 인덱스·중첩 키)의 진단 라인 조회 — 항목 경로가 없으면 상위 키 라인으로 폴백
package blogyaml

// llmsTxtItemLine resolves the diagnostic line for a nested geo.llms_txt
// path (e.g. "languages[1]", "pinned[0].url"), falling back to the parent
// key when the exact item is absent from the source (scalar shorthand etc.).
func llmsTxtItemLine(idx lineIndex, item, parent string) int {
	if line, ok := idx["geo.llms_txt."+item]; ok {
		return line
	}
	return llmsTxtLine(idx, parent)
}
