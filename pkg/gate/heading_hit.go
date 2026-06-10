//ff:func feature=gate type=parser control=sequence
//ff:what 본문 라인 1개를 헤딩으로 해석해 blog.yaml 헤딩 맵에 있으면 SectionHit으로 분류 (정확 일치, 어간 매칭 금지)
package gate

import "regexp"

var reHeading = regexp.MustCompile(`^(#{1,6})\s+(.*\S)\s*$`)

// headingHit classifies one trimmed body line as a recognized section heading.
func headingHit(hi headingIndex, lang, ln string, line int) (SectionHit, bool) {
	m := reHeading.FindStringSubmatch(ln)
	if m == nil {
		return SectionHit{}, false
	}
	key, ok := hi.byLang[lang][normText(m[2])]
	if !ok {
		return SectionHit{}, false
	}
	return SectionHit{Key: key, Level: len(m[1]), Text: m[2], Line: line}, true
}
