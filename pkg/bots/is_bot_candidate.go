//ff:func feature=bots type=dict control=iteration dimension=1 topic=crawl
//ff:what 미지 봇 후보 판정 — 비어있지 않은 UA에 후보 토큰(대소문자 무시)이 하나라도 있으면 true
package bots

import "strings"

// IsBotCandidate reports whether a User-Agent looks like a bot that is not
// necessarily in the dictionary: any candidate token appears in it. Empty
// and "-" UAs are not candidates — there is nothing to add to a dictionary.
func IsBotCandidate(userAgent string) bool {
	if userAgent == "" || userAgent == "-" {
		return false
	}
	ua := strings.ToLower(userAgent)
	for _, tok := range candidateTokens() {
		if strings.Contains(ua, tok) {
			return true
		}
	}
	return false
}
