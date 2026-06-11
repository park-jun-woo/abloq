//ff:func feature=quest type=frame control=sequence topic=queue
//ff:what AssembleTarget + git HEAD 원본 파싱 부착 — 큐 소비 퀘스트의 기준선 Target 조립, HEAD 부재·git 불가는 에러 (퀘스트 공용)
//ff:why writing/translation의 Base nil 규약(전량 신규 판정)과 다르다 — 큐 소비는 기존 글 수정이 본질이라 기준선 룰(honest-lastmod·claim 룰·front-matter-intact)이 HEAD 대비로 실동작해야 한다 (Phase018 계획)
package common

import agate "github.com/park-jun-woo/abloq/pkg/gate"

// AssembleBaseTarget builds the single-article gate target for one queue
// consumption: AssembleTarget plus the parsed git HEAD snapshot as the
// article's Base. Errors (unreadable article, article missing from HEAD,
// no git repository) abort the submit without burning a try.
func AssembleBaseTarget(root, article, lang, section, slug string) (*agate.Target, []byte, error) {
	tgt, body, err := AssembleTarget(root, article, lang, section, slug)
	if err != nil {
		return nil, nil, err
	}
	raw, err := gitHeadRaw(root, article)
	if err != nil {
		return nil, nil, err
	}
	tgt.Articles[0].Base = agate.ParseArticle(tgt.Blog, lang, string(raw))
	return tgt, body, nil
}
