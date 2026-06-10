//ff:func feature=gate type=parser control=iteration dimension=1 topic=evidence
//ff:what front matter claims_ignore 해석 — 비공백 사유 문자열 목록이면 글 단위 예외(exempt), 사유 누락이면 bad
package gate

import "strings"

// claimsIgnore reads the article-level claim exemption. claims_ignore must be
// a non-empty list of non-empty reason strings; exempt is true only then.
// bad is true when the key exists but states no usable reason.
func claimsIgnore(a *Article) (exempt, bad bool) {
	m, ok := fmMap(a.Doc.FrontMatter)
	if !ok {
		return false, false
	}
	raw, present := m["claims_ignore"]
	if !present {
		return false, false
	}
	list, ok := raw.([]any)
	if !ok || len(list) == 0 {
		return false, true
	}
	for _, e := range list {
		s, _ := e.(string)
		if strings.TrimSpace(s) == "" {
			return false, true
		}
	}
	return true, false
}
