//ff:func feature=quest type=parser control=sequence topic=lossless
//ff:what 내부 글 링크에서 언어 프리픽스(/{lang}/)를 제거해 언어 중립 경로로 정규화 — 프리픽스가 없으면 그대로 (패리티 ⑥(b))
//ff:why 제거는 원문 포함 양쪽에 적용한다 — 기본 언어가 루트 서빙이면 원문 링크에 프리픽스가 없지만, default_lang_in_subdir 사이트는 원문에도 프리픽스가 붙으므로 한쪽만 벗기면 비교가 어긋난다 (Phase017 검수 확정)
package translation

import "strings"

// stripLangPrefix normalizes one internal article link to its
// language-neutral path: "/<lang>/rest" becomes "/rest"; links without the
// prefix (default language served at the site root) pass through unchanged.
func stripLangPrefix(link, lang string) string {
	rest, ok := strings.CutPrefix(link, "/"+lang+"/")
	if !ok {
		return link
	}
	return "/" + rest
}
