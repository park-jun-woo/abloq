//ff:func feature=scan type=parser control=iteration dimension=1 topic=evidence
//ff:what 전 글의 인용 발생 수집 — gate.Citations 재사용, 글 내 같은 URL은 1건으로 (배치 upsert가 키를 두 번 못 만짐)
package evidence

import "github.com/park-jun-woo/abloq/pkg/gate"

// collectCitations flattens every article's citations into probe targets in
// document order. Within one article a URL appears once — citation_checks
// keys on (url, lang, section, slug) and a batch upsert must not touch the
// same key twice.
func collectCitations(arts []*gate.Article) []cite {
	cites := make([]cite, 0)
	for _, a := range arts {
		cites = append(cites, articleCites(a)...)
	}
	return cites
}
