//ff:func feature=visibility type=parser control=iteration dimension=1 topic=crawl
//ff:what 키 목록에서 시간 프리픽스가 있는 로그 키만 남김 — CLI 전량 수집의 키 선별
package cflog

// LogKeys keeps only the keys that carry a CloudFront hour prefix — the
// CLI's whole-source selection (no cursor, no margin).
func LogKeys(keys []string) []string {
	var out []string
	for _, key := range keys {
		if _, ok := hourOfKey(key); ok {
			out = append(out, key)
		}
	}
	return out
}
