//ff:func feature=blogyaml type=schema control=sequence
//ff:what 빌드 URL·public 경로의 언어 세그먼트 결정 — 기본 언어가 루트 서빙이면 빈 문자열, 아니면 언어 코드
//ff:why hugo defaultContentLanguageInSubdir=false 사이트(기본 언어 = 루트)에서 llms.txt URL·hreflang 페이지 경로·.md 병행 서빙 경로가 같은 규칙을 공유해야 한다 (Phase006 parkjunwoo.com 역이식에서 발견)
package blogyaml

// URLLang returns the language path segment for built URLs and public/ paths:
// empty when lang is the default language served at the site root
// (site.default_lang_in_subdir: false), the language code otherwise.
func (b *Blog) URLLang(lang string) string {
	if !b.Site.DefaultLangInSubdir && len(b.Languages) > 0 && lang == b.Languages[0] {
		return ""
	}
	return lang
}
