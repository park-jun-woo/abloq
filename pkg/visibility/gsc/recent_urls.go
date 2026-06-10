//ff:func feature=visibility type=client control=iteration dimension=1 topic=gsc
//ff:what 검사 대상 URL 선별 — lastmod가 오늘(UTC) 기준 최근 N일 내인 글의 URL을 상한(max)까지
//ff:why URL Inspection 쿼터 보호의 선별 규칙: 최근 갱신 글만 검사 가치가 있다 — lastmod는 저장소 파싱 기준(단일 소스 원칙) (Phase013)
package gsc

import "time"

// RecentURLs picks the inspection candidates: pages whose lastmod date falls
// within the last recentDays (inclusive), at most max of them, input order
// preserved. Lastmod comparison is lexicographic on the date prefix —
// ISO-8601 sorts as strings.
func RecentURLs(pages []PageMod, today time.Time, recentDays, max int) []string {
	cutoff := today.UTC().AddDate(0, 0, -recentDays).Format("2006-01-02")
	var urls []string
	for _, p := range pages {
		if len(urls) >= max {
			break
		}
		if len(p.Lastmod) < 10 || p.Lastmod[:10] < cutoff {
			continue
		}
		urls = append(urls, p.URL)
	}
	return urls
}
