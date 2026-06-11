//ff:func feature=quest type=parser control=iteration dimension=1 topic=lossless
//ff:what 내부 글 링크 목록 전체에 언어 프리픽스 제거를 적용 — ⑥(b) 양방향 multiset 비교의 정규화 입력
package translation

// stripAll maps stripLangPrefix over a link list (both the origin's and the
// translation's lists are normalized before the multiset comparison).
func stripAll(links []string, lang string) []string {
	var out []string
	for _, l := range links {
		out = append(out, stripLangPrefix(l, lang))
	}
	return out
}
