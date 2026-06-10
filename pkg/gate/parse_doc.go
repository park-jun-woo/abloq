//ff:func feature=gate type=parser control=sequence
//ff:what 글 원문을 구조 파싱본(Doc)으로 변환 — front matter 분리, 본문 스캔, 저작자 표기 라인 탐지
package gate

import "strings"

// parseDoc parses one article's full content for the given language.
// hi drives section-heading recognition (built from blog.yaml structure).
func parseDoc(hi headingIndex, lang, content string) *Doc {
	d := &Doc{FirstContentLine: -1, AttribLine: -1, BodyStart: 1}
	d.FrontMatter, d.Body, d.HasFM = splitFrontMatter(content)
	if d.HasFM {
		// line 1 fence + front matter lines + closing fence, body starts after
		d.BodyStart = strings.Count(d.FrontMatter, "\n") + 4
	}
	d.BodyLines = strings.Split(d.Body, "\n")
	scanBody(hi, lang, d)
	if d.FirstIsImage {
		d.AttribLine = attribAfterImage(d)
	}
	return d
}
