//ff:func feature=scan type=parser control=iteration dimension=1 topic=evidence
//ff:what 글 1편의 인용 URL을 중복 없이 cite로 변환 — 본문 등장 순서 유지
package evidence

import "github.com/park-jun-woo/abloq/pkg/gate"

// articleCites lists one article's unique citation URLs in document order.
func articleCites(a *gate.Article) []cite {
	seen := map[string]bool{}
	var cites []cite
	for _, c := range gate.Citations(a.Doc) {
		if seen[c.URL] {
			continue
		}
		seen[c.URL] = true
		cites = append(cites, cite{Lang: a.Lang, Section: a.Section, Slug: a.Slug, URL: c.URL})
	}
	return cites
}
