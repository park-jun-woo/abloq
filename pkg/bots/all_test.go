//ff:func feature=bots type=dict control=iteration dimension=1
//ff:what All이 정준 분류 순서(training→search→fetch)·UA 중복 없음·대표 봇 포함을 지키는지 검증
package bots

import "testing"

func TestAll(t *testing.T) {
	rank := map[string]int{"training": 0, "search": 1, "fetch": 2}
	seen := map[string]bool{}
	prev := 0
	for _, b := range All() {
		r, ok := rank[b.Category]
		if !ok {
			t.Errorf("unknown category %q for %s", b.Category, b.UserAgent)
		}
		if r < prev {
			t.Errorf("category order broken at %s (%s)", b.UserAgent, b.Category)
		}
		if seen[b.UserAgent] {
			t.Errorf("duplicate user agent %s", b.UserAgent)
		}
		prev = r
		seen[b.UserAgent] = true
	}
	if !seen["GPTBot"] || !seen["Bytespider"] || !seen["OAI-SearchBot"] || !seen["Claude-User"] {
		t.Errorf("dictionary missing a representative bot: %v", seen)
	}
}
