//ff:func feature=gen type=generator control=sequence
//ff:what resolvePolicy 케이스 하나를 실행해 정책 해석 결과를 검증
package robots

import (
	"testing"

	"github.com/park-jun-woo/abloq/pkg/bots"
)

func checkResolvePolicy(t *testing.T, crawlers map[string]string, bot bots.Bot, want string) {
	t.Helper()
	if got := resolvePolicy(crawlers, bot); got != want {
		t.Errorf("resolvePolicy(%v, %s) = %q, want %q", crawlers, bot.UserAgent, got, want)
	}
}
