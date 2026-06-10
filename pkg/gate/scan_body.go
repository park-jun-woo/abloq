//ff:func feature=gate type=parser control=iteration dimension=1
//ff:what 본문 라인을 스캔해 첫 콘텐츠 라인/메인 이미지 여부와 인식된 섹션 헤딩(정상/비정상 레벨)을 수집
package gate

import (
	"regexp"
	"strings"
)

var reImage = regexp.MustCompile(`^!\[[^\]]*\]\([^)]*\)`)

// scanBody fills the structural features of d from its body lines.
func scanBody(hi headingIndex, lang string, d *Doc) {
	for i, raw := range d.BodyLines {
		ln := strings.TrimSpace(raw)
		if ln == "" {
			continue
		}
		if d.FirstContentLine == -1 {
			d.FirstContentLine = i
			d.FirstIsImage = reImage.MatchString(ln)
		}
		hit, ok := headingHit(hi, lang, ln, i)
		if !ok {
			continue
		}
		if hit.Level == 2 {
			d.Sections = append(d.Sections, hit)
		} else {
			d.BadLevel = append(d.BadLevel, hit)
		}
	}
}
