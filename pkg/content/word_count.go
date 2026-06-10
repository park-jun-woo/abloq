//ff:func feature=content type=parser control=sequence
//ff:what 본문의 공백 구분 토큰 수 집계 — 인덱스의 word_count 지표 (언어 불문 단일 규칙)
package content

import "strings"

// wordCount counts whitespace-separated tokens in the body. One rule for all
// languages: the metric feeds relative comparisons (freshness/cluster queues),
// not typography.
func wordCount(body string) int64 {
	return int64(len(strings.Fields(body)))
}
