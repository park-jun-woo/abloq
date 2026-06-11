//ff:func feature=blogyaml type=generator control=sequence
//ff:what OG 프롬프트 렌더 — {title}/{summary}/{brand} 3종 고정 치환, 빈 템플릿은 내장 기본(no text·no words·safe margin 포함)
//ff:why AI가 글자를 그리는 사고를 프롬프트 레벨에서 차단 — 텍스트는 --overlay의 결정론 합성이 담당한다 (BUG002)
package blogyaml

import "strings"

// defaultOGPrompt is the built-in template when blog.yaml declares none. It
// forbids in-image text (the overlay owns typography) and asks for a safe
// central margin in the 1200x630 frame.
const defaultOGPrompt = `Minimal abstract background art for a blog article titled "{title}". ` +
	`No text, no words, no letters, no typography. ` +
	`1200x630 composition with a clear safe central margin for an overlaid title.`

// OGPrompt renders an OG prompt template, substituting the three fixed
// placeholders. An empty/blank template falls back to the built-in default.
func OGPrompt(tmpl, title, summary, brand string) string {
	if strings.TrimSpace(tmpl) == "" {
		tmpl = defaultOGPrompt
	}
	r := strings.NewReplacer("{title}", title, "{summary}", summary, "{brand}", brand)
	return r.Replace(tmpl)
}
