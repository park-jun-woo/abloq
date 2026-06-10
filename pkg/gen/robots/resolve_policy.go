//ff:func feature=gen type=generator control=sequence
//ff:what geo.crawlers에서 봇 하나의 정책을 해석 — 봇 이름 키(소문자) > 분류 키 > 기본 allow 순으로 우선
package robots

import (
	"strings"

	"github.com/park-jun-woo/abloq/pkg/bots"
)

// resolvePolicy returns "allow" or "block" for one bot.
// A lowercase bot-name key overrides its category key; absent keys default to allow.
func resolvePolicy(crawlers map[string]string, bot bots.Bot) string {
	if p, ok := crawlers[strings.ToLower(bot.UserAgent)]; ok {
		return p
	}
	if p, ok := crawlers[bot.Category]; ok {
		return p
	}
	return "allow"
}
