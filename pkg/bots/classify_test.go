//ff:func feature=bots type=dict control=iteration dimension=1 topic=crawl
//ff:what Classify가 실 UA 문자열을 부분일치로 분류하는지 검증 — 장식된 실 UA, 최장 토큰 우선, 대소문자 무시, 미매칭 false
package bots

import "testing"

func TestClassify(t *testing.T) {
	cases := []struct {
		ua   string
		want string
		ok   bool
	}{
		{"Mozilla/5.0 AppleWebKit/537.36 (KHTML, like Gecko); compatible; ChatGPT-User/1.0; +https://openai.com/bot", "ChatGPT-User", true},
		{"Mozilla/5.0 AppleWebKit/537.36 (KHTML, like Gecko; compatible; ClaudeBot/1.0; +claudebot@anthropic.com)", "ClaudeBot", true},
		{"Mozilla/5.0 AppleWebKit/537.36 (KHTML, like Gecko); compatible; GPTBot/1.2; +https://openai.com/gptbot", "GPTBot", true},
		{"mozilla/5.0 (compatible; perplexitybot/1.0; +https://perplexity.ai/perplexitybot)", "PerplexityBot", true},
		{"Mozilla/5.0 (compatible; Applebot-Extended/0.1)", "Applebot-Extended", true},
		{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36", "", false},
		{"", "", false},
	}
	for _, c := range cases {
		got, ok := Classify(c.ua)
		if got != c.want || ok != c.ok {
			t.Errorf("Classify(%q) = (%q, %v), want (%q, %v)", c.ua, got, ok, c.want, c.ok)
		}
	}
}
