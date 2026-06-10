//ff:func feature=gate type=parser control=sequence topic=evidence
//ff:what 텍스트에 출처 링크가 있는지 판정 — 인라인 http(s) 링크 또는 각주 참조([^n])
package gate

import "regexp"

var reSourceLink = regexp.MustCompile(`\]\(https?://[^)]+\)|\[\^[^\]]+\]`)

// hasSourceLink reports whether text carries a source link: a markdown inline
// link to an http(s) URL, or a footnote reference.
func hasSourceLink(text string) bool {
	return reSourceLink.MatchString(text)
}
