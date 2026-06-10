//ff:func feature=gate type=parser control=iteration dimension=1 topic=evidence
//ff:what 본문을 주장 검출 대상 문단으로 분할 — 코드 펜스/들여쓴 코드/인용/헤딩/이미지/구조 라인 제외, 빈 줄로 구분
package gate

import "strings"

// claimParas splits the body into claim-eligible paragraphs. Fenced and
// indented code, blockquotes, headings, image lines and the structural lines
// (main image, attribution, section headings) never carry gateable claims.
func claimParas(d *Doc) []claimPara {
	skip := structuralLines(d)
	var paras []claimPara
	var cur claimPara
	flush := func() {
		if len(cur.lines) > 0 {
			paras = append(paras, cur)
		}
		cur = claimPara{}
	}
	inFence := false
	for i, raw := range d.BodyLines {
		ln := strings.TrimSpace(raw)
		if strings.HasPrefix(ln, "```") || strings.HasPrefix(ln, "~~~") {
			inFence = !inFence
			continue
		}
		if ln == "" {
			flush()
			continue
		}
		if inFence || skip[i] || strings.HasPrefix(ln, ">") || strings.HasPrefix(ln, "#") ||
			strings.HasPrefix(ln, "![") || strings.HasPrefix(raw, "    ") || strings.HasPrefix(raw, "\t") {
			continue
		}
		cur.lines = append(cur.lines, i)
		cur.texts = append(cur.texts, raw)
	}
	flush()
	return paras
}
