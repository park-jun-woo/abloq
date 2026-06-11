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
	for _, must := range []string{"No text", "no words", "safe central margin", "focal subject", "not overexposed"} {
		if !strings.Contains(def, must) {
			t.Errorf("default template missing %q: %q", must, def)
		}
	}
	// empty summary leaves no dangling label (trailing whitespace only)
	if strings.Contains(def, "{summary}") {
		t.Errorf("default template left an unsubstituted {summary} slot: %q", def)
	}
	// non-empty summary reaches the default template (the {summary} slot exists)
	withSum := OGPrompt("", "My Post", "quantum compilers explained", "")
	if !strings.Contains(withSum, "quantum compilers explained") {
		t.Errorf("default template must carry the summary into the prompt: %q", withSum)
	}
	if blank := OGPrompt("   \n", "T", "", ""); !strings.Contains(blank, "No text") {
		t.Errorf("blank template must fall back to the default: %q", blank)
	}
}
