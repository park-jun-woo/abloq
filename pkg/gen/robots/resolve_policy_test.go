//ff:func feature=gen type=generator control=iteration dimension=1
//ff:what resolvePolicy가 봇 이름 키 > 분류 키 > 기본 allow 우선순위를 지키는지 검증
package robots

import (
	"testing"

	"github.com/park-jun-woo/abloq/pkg/bots"
)

func TestResolvePolicy(t *testing.T) {
	gptbot := bots.Bot{UserAgent: "GPTBot", Category: "training"}
	cases := []struct {
		name     string
		crawlers map[string]string
		bot      bots.Bot
		want     string
	}{
		{"empty defaults to allow", nil, gptbot, "allow"},
		{"category key", map[string]string{"training": "block"}, gptbot, "block"},
		{"bot name overrides category", map[string]string{"training": "block", "gptbot": "allow"}, gptbot, "allow"},
		{"bot name key is lowercase", map[string]string{"bytespider": "block"}, bots.Bot{UserAgent: "Bytespider", Category: "training"}, "block"},
		{"unrelated keys ignored", map[string]string{"search": "block"}, gptbot, "allow"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) { checkResolvePolicy(t, tc.crawlers, tc.bot, tc.want) })
	}
}
