//ff:func feature=gen type=generator control=iteration dimension=1
//ff:what hugo.toml 렌더가 멱등이고 baseURL/기본 언어/sitemap/permalink/언어 블록을 포함하는지 검증
package hugo

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func TestRender(t *testing.T) {
	b := &blogyaml.Blog{
		Site:      blogyaml.Site{BaseURL: "https://x.example.com", Title: "X", Author: "A"},
		Languages: []string{"ko", "en"},
		Sections:  []string{"opinion"},
	}
	out := string(Render(b))
	if again := string(Render(b)); again != out {
		t.Fatalf("Render is not idempotent:\n%s\n---\n%s", out, again)
	}
	wants := []string{
		`baseURL = "https://x.example.com"`,
		`title = "X"`,
		`defaultContentLanguage = "ko"`,
		"defaultContentLanguageInSubdir = true",
		"enableRobotsTXT = false",
		`author = "A"`,
		"[sitemap]",
		`opinion = "/opinion/:slug/"`,
		"[languages.en]",
	}
	for _, w := range wants {
		if !strings.Contains(out, w) {
			t.Errorf("want %q in hugo.toml, got:\n%s", w, out)
		}
	}
}
