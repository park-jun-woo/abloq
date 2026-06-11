//ff:func feature=quest type=parser control=sequence topic=queue
//ff:what Prepare — 제출 JSON 디코드, 대상 글 일치 검증, 공통 Consumption 조립(기준선 Target+porcelain 변경 집합+허용 집합)
//ff:why 글이 HEAD에 없거나 git 저장소가 아니면 Prepare 에러(try 미소진) — Base nil 폴백은 기준선 룰 침묵 통과 치즈라 금지한다 (2차 검수 확정)
package refresh

import (
	"encoding/json"
	"fmt"

	rgate "github.com/park-jun-woo/reins/pkg/gate"
	"github.com/park-jun-woo/reins/pkg/quest"

	"github.com/park-jun-woo/abloq/pkg/quests/common"
)

// Prepare decodes one submission into a gate Context. Errors (bad JSON,
// wrong target, unreadable article, article missing from HEAD, no git
// repository) abort the submit without burning a try. No short verdict is
// ever returned (nothing is untrusted here).
func (Definition) Prepare(_ *quest.Session, it *quest.Item, raw []byte) (rgate.Context, *quest.Verdict, error) {
	var p common.QueuePayload
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
	cons, body, err := common.PrepareConsumption(p)
	if err != nil {
		return rgate.Context{}, nil, err
	}
	sub := &Submission{Consumption: cons, Article: p.Article}
	return rgate.Context{Item: it, Submission: sub, Source: string(body)}, nil, nil
}
