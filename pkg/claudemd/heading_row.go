//ff:func feature=claudemd type=generator control=iteration dimension=1
//ff:what 헤딩 키 1개의 언어별 현지화 헤딩을 마크다운 표 행으로 렌더 — 미정의 언어는 빈 칸
package claudemd

import "strings"

// headingRow renders one table row: key plus the heading text per language.
func headingRow(key string, langs []string, m map[string]string) string {
	cells := make([]string, 0, len(langs)+1)
	cells = append(cells, key)
	for _, lang := range langs {
		cells = append(cells, m[lang])
	}
	return "| " + strings.Join(cells, " | ") + " |\n"
}
