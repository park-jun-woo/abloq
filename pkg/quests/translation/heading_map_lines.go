//ff:func feature=quest type=generator control=iteration dimension=1
//ff:what blog.yaml structure.headings에서 원문→대상 언어 헤딩 맵 라인 렌더 — order 순서 유지, 헤딩 미선언 키는 생략
package translation

import (
	"fmt"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// headingMapLines renders the localized heading pairs the translator must use
// verbatim, in structure.order: `- <key>: "<origin text>" -> "<target text>"`.
// Keys without a declared heading (e.g. body, image) are layout slots, not
// headings, and are omitted.
func headingMapLines(b *blogyaml.Blog, originLang, lang string) []string {
	var lines []string
	for _, key := range b.Structure.Order {
		langs, ok := b.Structure.Headings[key]
		if !ok {
			continue
		}
		lines = append(lines, fmt.Sprintf("- %s: %q -> %q", key, langs[originLang], langs[lang]))
	}
	return lines
}
