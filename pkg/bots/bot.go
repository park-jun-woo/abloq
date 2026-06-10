//ff:type feature=bots type=schema
//ff:what AI봇 UA 사전 항목 — User-Agent 토큰과 분류(training/search/fetch)
package bots

// Bot is one AI crawler dictionary entry.
// Category is one of "training" (model training crawlers),
// "search" (AI search indexers) or "fetch" (user-triggered fetchers).
type Bot struct {
	UserAgent string
	Category  string
}
