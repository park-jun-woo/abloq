//ff:func feature=image type=generator control=iteration dimension=1
//ff:what 텍스트를 폰트 측정 기반 그리디 줄바꿈 — 픽셀 폭 maxW를 넘지 않는 줄 목록 반환
package img

import (
	"strings"

	"golang.org/x/image/font"
)

// WrapText splits text into lines no wider than maxW pixels under face.
func WrapText(face font.Face, text string, maxW int) []string {
	var lines []string
	line := ""
	for _, word := range strings.Fields(text) {
		cand := word
		if line != "" {
			cand = line + " " + word
		}
		if line != "" && font.MeasureString(face, cand).Ceil() > maxW {
			lines = append(lines, line)
			line = word
			continue
		}
		line = cand
	}
	if line != "" {
		lines = append(lines, line)
	}
	return lines
}
