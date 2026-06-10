//ff:func feature=gate type=parser control=iteration dimension=1
//ff:what 메인 이미지 다음 첫 비공백 라인이 이탤릭 저작자 표기(*…*)인지 찾아 그 라인 인덱스를 반환 (-1 = 없음)
package gate

import (
	"regexp"
	"strings"
)

var reAttrib = regexp.MustCompile(`^\*[^*].*\*\s*$`)

// attribAfterImage returns the BodyLines index of the attribution line right
// after the main image, or -1 when the next non-blank line is not one.
func attribAfterImage(d *Doc) int {
	for i := d.FirstContentLine + 1; i < len(d.BodyLines); i++ {
		ln := strings.TrimSpace(d.BodyLines[i])
		if ln == "" {
			continue
		}
		if reAttrib.MatchString(ln) {
			return i
		}
		return -1
	}
	return -1
}
