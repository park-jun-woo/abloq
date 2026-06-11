//ff:func feature=quest type=frame control=sequence
//ff:what blog.yaml 로드 + 대상 글 읽기·파싱 → 단일 글 게이트 Target 조립 — 공용 추출본(quests/common.AssembleTarget) 위임, Base nil 규약
//ff:why Phase017에서 번역 퀘스트와 공유하도록 구현을 pkg/quests/common으로 추출 — 복제 금지, writing은 추출본을 쓴다
package writing

import (
	agate "github.com/park-jun-woo/abloq/pkg/gate"
	"github.com/park-jun-woo/abloq/pkg/quests/common"
)

// assembleTarget builds the single-article gate target for one submission
// (Base nil — the quest convention). The implementation lives in
// pkg/quests/common (Phase017 extraction).
func assembleTarget(p Payload) (*agate.Target, []byte, error) {
	return common.AssembleTarget(p.Root, p.Article, p.Lang, p.Section, p.Slug)
}
