//ff:func feature=quest type=parser control=iteration dimension=1
//ff:what 원문 1편 × (선언 언어 − 기본 언어) Item 생성 — 기존 번역 lastmod ≥ 원문이면 미생성(갱신 필요 언어만), Key=lang/section/slug
//ff:why 부분 실패는 reins 래칫의 자연 동작 — 아이템 독립이라 일부 언어 FAIL이 타 언어 PASS에 무영향, 전체 롤백 없음 (Phase017 계획)
package translation

import "github.com/park-jun-woo/reins/pkg/quest"

// seedItems expands one origin into per-language TODO items. A language whose
// translation file already exists with lastmod >= the origin's is up to date
// and gets no item; when the origin lastmod is unparseable every translation
// is treated as stale (Prepare rejects such an origin before any try burns).
func seedItems(src seedSrc) ([]*quest.Item, error) {
	var items []*quest.Item
	for _, lang := range src.blog.Languages[1:] {
		article := transPath(src, lang)
		if transFresh(src, lang, article) {
			continue
		}
		it := &quest.Item{Key: lang + "/" + src.section + "/" + src.slug, State: quest.TODO}
		p := Payload{Root: src.root, Origin: src.origin, Article: article,
			OriginLang: src.originLang, Lang: lang, Section: src.section, Slug: src.slug}
		if err := it.SetPayload(p); err != nil {
			return nil, err
		}
		items = append(items, it)
	}
	return items, nil
}
