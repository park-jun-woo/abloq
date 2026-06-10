//ff:func feature=scan type=client control=iteration dimension=1 topic=evidence
//ff:what 전체 URL 일괄 점검 — 호스트별 그룹을 Concurrency 상한의 고루틴으로 병렬 실행, 호스트 내부는 checkHost가 순차
package evidence

// CheckAll probes every unique URL and returns url → status ("ok"|"hard"|
// "soft"). Parallelism is across hosts only (bounded by Concurrency); within
// a host checkHost enforces the per-domain rate limit.
func (c *Checker) CheckAll(urls []string) map[string]string {
	groups := groupByHost(urls)
	out := make(chan map[string]string, len(groups))
	sem := make(chan struct{}, max(c.Concurrency, 1))
	for _, group := range groups {
		sem <- struct{}{}
		go func(g []string) {
			out <- c.checkHost(g)
			<-sem
		}(group)
	}
	return mergeResults(out, len(groups))
}
