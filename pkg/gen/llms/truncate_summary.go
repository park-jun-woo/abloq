//ff:func feature=gen type=generator control=sequence
//ff:what 설명문을 max_summary rune 수로 절단하고 "…" 1자를 덧붙임 — 0이면 무제한, 바이트 절단 없음(다국어 안전)
package llms

// truncateSummary caps a description at max runes plus a single "…".
// max <= 0 means unlimited; runes (never bytes) keep multibyte text intact.
func truncateSummary(s string, max int) string {
	if max <= 0 {
		return s
	}
	r := []rune(s)
	if len(r) <= max {
		return s
	}
	return string(r[:max]) + "…"
}
