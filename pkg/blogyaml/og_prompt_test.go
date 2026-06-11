//ff:func feature=blogyaml type=generator control=iteration dimension=1
//ff:what OGPrompt의 {title}/{summary}/{brand} 치환과 빈 템플릿의 내장 기본(no text·safe margin 포함) 폴백 검증
package blogyaml

import (
	"strings"
	"testing"
)

func TestOGPrompt(t *testing.T) {
	got := OGPrompt(`Art for "{title}" about {summary}, brand {brand}.`, "My Post", "a summary", "my.com")
	if got != `Art for "My Post" about a summary, brand my.com.` {
		t.Errorf("substitution = %q", got)
	}
	def := OGPrompt("", "My Post", "", "")
	if !strings.Contains(def, "My Post") {
		t.Errorf("default template must substitute the title: %q", def)
	}
	for _, must := range []string{"No text", "no words", "safe central margin"} {
		if !strings.Contains(def, must) {
			t.Errorf("default template missing %q: %q", must, def)
		}
	}
	if blank := OGPrompt("   \n", "T", "", ""); !strings.Contains(blank, "No text") {
		t.Errorf("blank template must fall back to the default: %q", blank)
	}
}
