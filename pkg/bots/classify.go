//ff:func feature=bots type=dict control=iteration dimension=1 topic=crawl
//ff:what 실 UA 문자열 전체에서 사전 토큰 부분일치(대소문자 무시)로 봇 이름을 분류 — 최장 토큰 우선, 미매칭 false
//ff:why CategoryOf는 토큰 정확일치라 실 크롤러 UA(Mozilla/5.0 … ChatGPT-User/1.0 …)에 매칭 불가 — 로그 분류는 contains가 필요하고, 토큰이 다른 토큰을 포함할 수 있어(Applebot ⊂ Applebot-Extended 류) 최장 일치로 결정성을 고정한다 (Phase012)
package bots

import "strings"

// Classify matches a full User-Agent string against the bot dictionary by
// case-insensitive substring and returns the canonical dictionary name. When
// several tokens match, the longest one wins. The second return is false
// when no dictionary token appears in the UA.
func Classify(userAgent string) (string, bool) {
	ua := strings.ToLower(userAgent)
	best := ""
	for _, b := range All() {
		if len(b.UserAgent) > len(best) && strings.Contains(ua, strings.ToLower(b.UserAgent)) {
			best = b.UserAgent
		}
	}
	return best, best != ""
}
