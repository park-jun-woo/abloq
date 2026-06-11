//ff:type feature=quest type=schema topic=queue
//ff:what 큐 소비 퀘스트 공통 제출 컨텍스트 — git 기준선 Target, Prepare 시점 변경 파일 집합(porcelain), queue-scope 허용 집합, 큐 payload 주장 해시
//ff:why Phase018 소비 퀘스트 3종(refresh/evidence/cluster)이 같은 룰(queue-scope·claim-scope·claim-preserved)을 공유한다 — 퀘스트별 Submission이 이 구조체를 임베드해 메서드 승격으로 공통 계약을 충족한다
package common

import agate "github.com/park-jun-woo/abloq/pkg/gate"

// Consumption is the shared submission context of every queue-consuming
// quest: the single-article gate target with the git HEAD baseline attached,
// the working-tree change set captured at Prepare time (git status
// --porcelain, untracked included), the queue-scope allowed path set, and
// the claim hashes the queue payload authorizes to change (evidence kind;
// empty otherwise). Quest submissions embed *Consumption so the shared rules
// reach it through the ConsCarrier contract.
type Consumption struct {
	Target       *agate.Target
	Changed      []string
	Allowed      map[string]bool
	QueuedClaims map[string]bool
}
