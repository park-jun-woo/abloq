//ff:func feature=init type=dict control=sequence
//ff:what 언어 코드별 sources 섹션 현지화 헤딩 사전 — 미등록 언어는 "Sources" 폴백
package main

// sourcesHeadings localizes the canonical sources heading per language.
var sourcesHeadings = map[string]string{
	"ko": "출처", "en": "Sources", "ja": "出典", "zh": "来源",
	"es": "Fuentes", "fr": "Sources", "de": "Quellen", "pt": "Fontes",
	"ru": "Источники", "id": "Sumber", "ar": "المصادر", "he": "מקורות",
}

// sourcesHeading returns the localized sources heading for lang.
func sourcesHeading(lang string) string {
	if h, ok := sourcesHeadings[lang]; ok {
		return h
	}
	return "Sources"
}
