//ff:func feature=gen type=generator control=iteration dimension=1
//ff:what robots.txt 렌더가 멱등이고 분류 정책·봇 이름 오버라이드·기본 그룹·Sitemap을 정확히 전개하는지 검증
package robots

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func TestRender(t *testing.T) {
	b := &blogyaml.Blog{
		Site: blogyaml.Site{BaseURL: "https://x.example.com"},
		Geo: blogyaml.Geo{Crawlers: map[string]string{
			"training": "block", "gptbot": "allow", "bytespider": "block",
		}},
	}
	out := string(Render(b))
	if again := string(Render(b)); again != out {
		t.Fatalf("Render is not idempotent:\n%s\n---\n%s", out, again)
	}
	wants := []string{
		"User-agent: GPTBot\nDisallow:\n",        // bot-name override beats category
		"User-agent: ClaudeBot\nDisallow: /\n",   // category training = block
		"User-agent: Bytespider\nDisallow: /\n",  // bot-name block
		"User-agent: OAI-SearchBot\nDisallow:\n", // unspecified search defaults to allow
		"User-agent: *\nAllow: /\n",              // default group
		"Sitemap: https://x.example.com/sitemap.xml\n",
	}
	for _, w := range wants {
		if !strings.Contains(out, w) {
			t.Errorf("want %q in robots.txt, got:\n%s", w, out)
		}
	}
}
