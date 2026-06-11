//ff:type feature=quest type=schema
//ff:what 퀘스트 제출물 공통 계약 — 게이트 Target을 내놓는 인터페이스, AdaptRule이 퀘스트별 Submission 타입에 비의존하게 만든다
//ff:why Phase017에서 writing 프라이빗 어댑터를 공용 추출하며 신설 — 두 번째 소비자(번역 퀘스트)가 검증한 추상화 (No N=1 abstraction)
package common

import agate "github.com/park-jun-woo/abloq/pkg/gate"

// TargetCarrier is what every quest submission must expose for the shared
// rule adapter: the assembled single-article gate target under review.
type TargetCarrier interface {
	GateTarget() *agate.Target
}
