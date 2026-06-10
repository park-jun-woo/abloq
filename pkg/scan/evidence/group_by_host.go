//ff:func feature=scan type=client control=iteration dimension=1 topic=evidence
//ff:what URL 목록을 호스트별로 묶음 — 중복 URL 제거, 호스트 단위가 도메인별 rate limit의 적용 단위
package evidence

import "net/url"

// groupByHost buckets unique URLs by host. Unparseable URLs land in the ""
// bucket and still get probed (the probe itself will fail them as soft).
func groupByHost(urls []string) map[string][]string {
	seen := map[string]bool{}
	groups := map[string][]string{}
	for _, raw := range urls {
		if seen[raw] {
			continue
		}
		seen[raw] = true
		host := ""
		if u, err := url.Parse(raw); err == nil {
			host = u.Host
		}
		groups[host] = append(groups[host], raw)
	}
	return groups
}
