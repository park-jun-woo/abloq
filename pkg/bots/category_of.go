//ff:func feature=bots type=dict control=iteration dimension=1
//ff:what User-Agent 토큰(대소문자 무시)으로 사전을 조회해 분류(training/search/fetch)를 반환
package bots

import "strings"

// CategoryOf looks up a bot category by User-Agent token, case-insensitively.
// The second return is false when the token is not in the dictionary.
func CategoryOf(userAgent string) (string, bool) {
	for _, b := range All() {
		if strings.EqualFold(b.UserAgent, userAgent) {
			return b.Category, true
		}
	}
	return "", false
}
