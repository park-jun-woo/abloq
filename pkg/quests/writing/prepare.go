//ff:func feature=quest type=parser control=sequence
//ff:what Prepare — 제출 JSON 디코드, 대상 글 일치 검증(치즈 방어), 단일 글 Target 조립, match 미출현 산출, REVIEW·로그 적재
package writing

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	rgate "github.com/park-jun-woo/reins/pkg/gate"
	"github.com/park-jun-woo/reins/pkg/quest"

	"github.com/park-jun-woo/abloq/pkg/insight"
)

// Prepare decodes one submission into a gate Context. Errors (bad JSON, wrong
// target, unreadable article/spec) abort the submit without burning a try;
// absent review/worklog files are NOT errors — the review-record rule judges
// them. No short verdict is ever returned (nothing is untrusted here).
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
	tgt, body, err := assembleTarget(p)
	if err != nil {
		return rgate.Context{}, nil, err
	}
	ins, errs, _, err := insight.Load(filepath.Join(p.Root, filepath.FromSlash(p.Insight)))
	if err != nil {
		return rgate.Context{}, nil, err
	}
	if len(errs) > 0 {
		return rgate.Context{}, nil, fmt.Errorf("%s", errs[0].String())
	}
	res := insight.Match(ins, p.Article, body)
	sub := &Submission{
		Target: tgt, Article: p.Article, Missing: res.Missing,
		Review: readOptional(p.Root, sf.Review), ReviewPath: sf.Review,
		Worklog: readOptional(p.Root, sf.Worklog), WorklogPath: sf.Worklog,
	}
	return rgate.Context{Item: it, Submission: sub, Source: string(body)}, nil, nil
}
