//ff:func feature=gen type=generator control=iteration dimension=1
//ff:what hugo.toml 렌더가 멱등이고 baseURL/기본 언어/sitemap/언어 블록을 포함하며 [permalinks]를 내지 않는지 검증
package hugo

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func TestRender(t *testing.T) {
	b := &blogyaml.Blog{
		Site:      blogyaml.Site{BaseURL: "https://x.example.com", Title: "X", Author: "A", DefaultLangInSubdir: true},
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
		`ignoreFiles = ['insight\.yaml$']`,
		`author = "A"`,
		"[sitemap]",
		"[languages.en]",
	}
	for _, w := range wants {
		if !strings.Contains(out, w) {
			t.Errorf("want %q in hugo.toml, got:\n%s", w, out)
		}
	}
	if strings.Contains(out, "[permalinks]") {
		t.Errorf("hugo.toml must not declare [permalinks] (default path URLs = slug contract), got:\n%s", out)
	}
	b.Site.DefaultLangInSubdir = false
	if root := string(Render(b)); !strings.Contains(root, "defaultContentLanguageInSubdir = false") {
		t.Errorf("want defaultContentLanguageInSubdir = false for root-served default lang, got:\n%s", root)
	}
}
