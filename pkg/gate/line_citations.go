//ff:func feature=gate type=parser control=iteration dimension=1 topic=evidence
//ff:what 라인 1개에서 인라인 http(s) 링크를 인용으로 추출 — 이미지(![..](..))는 제외
package gate

import "regexp"

var reCitation = regexp.MustCompile(`\[([^\]]*)\]\((https?://[^)\s]+)\)`)

// lineCitations extracts the external citations written on one line.
// Image embeds are not citations.
func lineCitations(ln string, fileLine int) []Citation {
	var out []Citation
	for _, m := range reCitation.FindAllStringSubmatchIndex(ln, -1) {
		if m[0] > 0 && ln[m[0]-1] == '!' {
			continue
		}
		out = append(out, Citation{Label: ln[m[2]:m[3]], URL: ln[m[4]:m[5]], Line: fileLine})
	}
	return out
}
