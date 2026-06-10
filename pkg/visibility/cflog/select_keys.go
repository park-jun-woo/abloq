//ff:func feature=visibility type=parser control=iteration dimension=1 topic=crawl
//ff:what 커서 시간대 초과·마지막 닫힌 시간대 이하의 로그 키만 선별 — 시간 프리픽스 없는 키는 제외
//ff:why "cursorHour < hour <= lastClosed"가 수집 구간의 전부다 — 키 자체가 아니라 키의 시간 프리픽스로 거른다(랜덤 접미사 start-after 금지). YYYY-MM-DD-HH는 사전순 = 시간순 (Phase012)
package cflog

// selectKeys keeps the log keys whose hour prefix lies in the half-open
// ingest window (cursorHour, lastClosed]. Keys without an hour prefix are
// not log objects.
func selectKeys(keys []string, cursorHour, lastClosed string) []string {
	var out []string
	for _, key := range keys {
		hour, ok := hourOfKey(key)
		if ok && hour > cursorHour && hour <= lastClosed {
			out = append(out, key)
		}
	}
	return out
}
