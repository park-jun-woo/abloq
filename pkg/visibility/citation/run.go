//ff:func feature=visibility type=client control=iteration dimension=1 topic=citation
//ff:what 샘플링 1회전 — 전 엔진에 같은 질의 셋을 budget 상한으로 실행해 샘플을 엔진 순서대로 합산 (budget 0 = 비활성 no-op)
//ff:why 비용 상한은 blog.yaml geo.citation_budget(엔진당 최대 질의 수, 기본 0 옵트인) — 러너는 판정에 아무것도 입력하지 않는다(§6.3 비결정성 격리) (Phase013)
package citation

import "time"

// Run executes one sampling round: every engine answers the first budget
// queries (id order — the caller supplies them ordered). budget <= 0 means
// sampling is disabled and the round is a no-op.
func Run(engines []Engine, queries []Query, budget int, domain string, interval time.Duration) []Sample {
	if budget <= 0 {
		return nil
	}
	var samples []Sample
	for _, e := range engines {
		samples = append(samples, runEngine(e, queries, budget, domain, interval)...)
	}
	return samples
}
