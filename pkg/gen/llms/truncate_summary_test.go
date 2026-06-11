//ff:func feature=gen type=generator control=sequence
//ff:what truncateSummary가 0=무제한, 상한 이하 원문 유지, 초과 시 rune 단위 절단+"…" 1자(다국어 안전)인지 검증
package llms

import "testing"

func TestTruncateSummary(t *testing.T) {
	if got := truncateSummary("abcdef", 0); got != "abcdef" {
		t.Errorf("max 0 must be unlimited, got %q", got)
	}
	if got := truncateSummary("abc", 5); got != "abc" {
		t.Errorf("short text must pass through, got %q", got)
	}
	if got := truncateSummary("abcdef", 4); got != "abcd…" {
		t.Errorf("want rune cut + ellipsis, got %q", got)
	}
	if got := truncateSummary("가나다라마", 3); got != "가나다…" {
		t.Errorf("multibyte text must cut on runes, got %q", got)
	}
}
