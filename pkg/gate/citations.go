//ff:func feature=gate type=parser control=iteration dimension=1 topic=evidence
//ff:what 공개 Citations API — 글 1편의 외부 인용 링크 전부 추출, 코드 펜스 내부 제외 (Phase010 스캐너 재사용)
package gate

import "strings"

// Citations extracts every external citation from the article body, in
// document order. Links inside fenced code blocks are not citations.
// The citation-exists gate rule and the link-rot scanner share this extractor.
func Citations(d *Doc) []Citation {
	inFence := false
	var out []Citation
	for i, raw := range d.BodyLines {
		ln := strings.TrimSpace(raw)
		if strings.HasPrefix(ln, "```") || strings.HasPrefix(ln, "~~~") {
			inFence = !inFence
			continue
		}
		if inFence {
			continue
		}
		out = append(out, lineCitations(raw, bodyLine(d, i))...)
	}
	return out
}
