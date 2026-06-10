//ff:func feature=gen type=generator control=sequence
//ff:what llms.txt 목록 항목 1줄 렌더 — "- [제목](URL)" + 설명이 있으면 ": 설명"
package llms

import "fmt"

// postLine renders one llms.txt list entry.
func postLine(baseURL string, p Post) string {
	line := fmt.Sprintf("- [%s](%s)", p.Title, postURL(baseURL, p))
	if p.Description != "" {
		line += ": " + p.Description
	}
	return line
}
