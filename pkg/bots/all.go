//ff:func feature=bots type=dict control=sequence
//ff:what AI봇 UA 사전 전체를 정준 순서(분류 training→search→fetch, 분류 내 UA 알파벳순)로 반환
//ff:why robots.txt 생성의 멱등성을 위해 슬라이스 리터럴로 순서를 고정 — 맵 순회 비결정성 배제
package bots

// All returns the full AI bot dictionary in canonical order:
// training -> search -> fetch, alphabetical within each category.
func All() []Bot {
	return []Bot{
		{UserAgent: "Amazonbot", Category: "training"},
		{UserAgent: "Applebot-Extended", Category: "training"},
		{UserAgent: "Bytespider", Category: "training"},
		{UserAgent: "CCBot", Category: "training"},
		{UserAgent: "ClaudeBot", Category: "training"},
		{UserAgent: "Google-Extended", Category: "training"},
		{UserAgent: "GPTBot", Category: "training"},
		{UserAgent: "meta-externalagent", Category: "training"},
		{UserAgent: "Claude-SearchBot", Category: "search"},
		{UserAgent: "DuckAssistBot", Category: "search"},
		{UserAgent: "OAI-SearchBot", Category: "search"},
		{UserAgent: "PerplexityBot", Category: "search"},
		{UserAgent: "ChatGPT-User", Category: "fetch"},
		{UserAgent: "Claude-User", Category: "fetch"},
		{UserAgent: "Meta-ExternalFetcher", Category: "fetch"},
		{UserAgent: "MistralAI-User", Category: "fetch"},
		{UserAgent: "Perplexity-User", Category: "fetch"},
	}
}
