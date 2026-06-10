//ff:func feature=gate type=output control=sequence topic=diagnostics
//ff:what 진단 메시지용 텍스트 절단 — 80바이트 초과분을 말줄임표로 대체
package gate

// trunc shortens s for diagnostic messages.
func trunc(s string) string {
	const n = 80
	if len(s) > n {
		return s[:n] + "…"
	}
	return s
}
