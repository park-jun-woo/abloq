//ff:func feature=gen type=generator control=sequence
//ff:what 섹션 그룹 헤딩 텍스트 조립 — section_labels 라벨 우선, 다언어 스코프일 때만 "{lang}/" 접두
package llms

import "github.com/park-jun-woo/abloq/pkg/blogyaml"

// headingText builds one group heading: the section_labels label (or the raw
// section name), prefixed with "{lang}/" only in a multi-language scope.
func headingText(b *blogyaml.Blog, lang, section string, multi bool) string {
	label := section
	if l, ok := b.Geo.LlmsTxt.SectionLabels[section]; ok {
		label = l
	}
	if multi {
		return lang + "/" + label
	}
	return label
}
