//ff:func feature=scan type=rule control=selection topic=evidence
//ff:what HTTP 상태 코드 → 점검 판정 — 2xx/3xx는 ok, 404/410은 hard(경성), 그 외(5xx 포함)는 soft(연성)
//ff:why rot 확정은 분류와 무관하게 연속 실패 횟수로만 한다 — 분류는 진단 표시용이며, 판정을 가르는 것은 실패의 지속성이다
package evidence

// classify maps one probe's HTTP status code to the check status. Network
// errors never reach here (the probe classifies them as soft directly).
func classify(code int) string {
	switch {
	case code < 400:
		return "ok"
	case code == 404 || code == 410:
		return "hard"
	default:
		return "soft"
	}
}
