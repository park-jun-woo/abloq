//ff:func feature=scan type=client control=iteration dimension=1 topic=evidence
//ff:what 한 호스트의 URL들을 순차 점검 — 요청 사이 DomainDelay 대기(도메인별 rate limit), url→판정 맵 반환
package evidence

import "time"

// checkHost probes one host's URLs sequentially, sleeping the per-domain
// interval between consecutive probes. Sequential-per-host is the rate
// limit: the checker never holds two in-flight requests to one domain.
func (c *Checker) checkHost(urls []string) map[string]string {
	res := make(map[string]string, len(urls))
	for i, u := range urls {
		if i > 0 {
			time.Sleep(c.DomainDelay)
		}
		res[u] = c.checkURL(u)
	}
	return res
}
