//ff:func feature=init type=generator control=iteration dimension=1
//ff:what 언어 목록의 sources 헤딩 인라인 맵("ko: "출처", en: "Sources"")을 렌더 — 선언 순서 유지
package main

import (
	"fmt"
	"strings"
)

// headingLine renders the inline YAML map for structure.headings.sources.
func headingLine(langs []string) string {
	pairs := make([]string, 0, len(langs))
	for _, lang := range langs {
		pairs = append(pairs, fmt.Sprintf("%s: %q", lang, sourcesHeading(lang)))
	}
	return strings.Join(pairs, ", ")
}
