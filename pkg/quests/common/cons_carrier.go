//ff:type feature=quest type=schema topic=queue
//ff:what 큐 소비 제출물 공통 계약 — 공유 룰(queue-scope·claim 룰)이 퀘스트별 Submission 타입에 비의존하게 만드는 인터페이스
package common

// ConsCarrier is what every queue-consuming quest submission must expose for
// the shared queue rules: the consumption context under review.
type ConsCarrier interface {
	Cons() *Consumption
}
