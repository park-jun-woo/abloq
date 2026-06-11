//ff:func feature=quest type=frame control=sequence topic=queue
//ff:what Consumption의 ConsCarrier 자기 구현 — 임베드한 퀘스트 Submission이 메서드 승격으로 공통 룰 계약을 충족
package common

// Cons exposes the shared consumption context to the queue rules. A quest
// submission embedding *Consumption satisfies ConsCarrier by promotion.
func (c *Consumption) Cons() *Consumption { return c }
