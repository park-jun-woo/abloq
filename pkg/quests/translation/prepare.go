//ff:func feature=quest type=parser control=sequence
//ff:what Prepare — 제출 JSON 디코드, 대상 글 일치 검증, hugo 존재 사전 점검, 번역 Target 조립 + 원문 적재, 원문 date·lastmod 부재는 에러(try 미소진)
//ff:why 원문 lastmod 부재는 front-matter-schema와 fm-mirror(⑦)의 충족 불가 쌍 — 원문 자체가 게이트 위반이라 번역이 진행 불가하므로 FAIL(try 소진)이 아니라 Prepare 에러로 중단한다 (Phase017 계획 ⑦)
package translation

import (
	"encoding/json"
	"fmt"

	rgate "github.com/park-jun-woo/reins/pkg/gate"
	"github.com/park-jun-woo/reins/pkg/quest"

	"github.com/park-jun-woo/abloq/pkg/quests/common"
)

// Prepare decodes one submission into a gate Context. Errors (bad JSON, wrong
// target, hugo missing, unreadable article/origin, origin date/lastmod
// missing) abort the submit without burning a try. No short verdict is ever
// returned (nothing is untrusted here).
func (Definition) Prepare(_ *quest.Session, it *quest.Item, raw []byte) (rgate.Context, *quest.Verdict, error) {
	var p Payload
	if err := it.DecodePayload(&p); err != nil {
		return rgate.Context{}, nil, err
	}
	var sf SubmitFile
	if err := json.Unmarshal(raw, &sf); err != nil {
		return rgate.Context{}, nil, fmt.Errorf("submission JSON: %w", err)
	}
	if sf.Article != p.Article {
		return rgate.Context{}, nil, fmt.Errorf("submission article %q does not match the seeded target %q", sf.Article, p.Article)
	}
	if _, err := hugoPath(); err != nil {
		return rgate.Context{}, nil, fmt.Errorf("hugo binary not found in PATH (the gate builds the whole instance): %w", err)
	}
	tgt, body, err := common.AssembleTarget(p.Root, p.Article, p.Lang, p.Section, p.Slug)
	if err != nil {
		return rgate.Context{}, nil, err
	}
	origin, err := loadOrigin(tgt.Blog, p)
	if err != nil {
		return rgate.Context{}, nil, err
	}
	if _, ok := fmTime(origin.Doc, "date"); !ok {
		return rgate.Context{}, nil, fmt.Errorf("origin %s: front matter date missing or unparseable — the origin itself violates front-matter-schema; fix the origin first", p.Origin)
	}
	if _, ok := fmTime(origin.Doc, "lastmod"); !ok {
		return rgate.Context{}, nil, fmt.Errorf("origin %s: front matter lastmod missing or unparseable — the origin itself violates front-matter-schema; fix the origin first", p.Origin)
	}
	sub := &Submission{Target: tgt, Origin: origin, Article: p.Article,
		Root: p.Root, Lang: p.Lang, OriginLang: p.OriginLang}
	return rgate.Context{Item: it, Submission: sub, Source: string(body)}, nil, nil
}
