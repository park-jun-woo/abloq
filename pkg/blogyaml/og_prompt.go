//ff:func feature=blogyaml type=generator control=sequence
//ff:what OG 프롬프트 렌더 — {title}/{summary}/{brand} 3종 고정 치환, 빈 템플릿은 내장 기본(초점 피사체·과노출 방지·no text·safe margin 포함)
//ff:why AI가 글자를 그리는 사고를 프롬프트 레벨에서 차단 — 텍스트는 --overlay의 결정론 합성이 담당한다 (BUG002)
package blogyaml

import "strings"

// defaultOGPrompt is the built-in template when blog.yaml declares none. It
// asks for a clear focal subject with balanced exposure (guarding against the
// themeless, blown-out abstract backgrounds the old title-only prompt produced),
// forbids in-image text (the overlay owns typography) and keeps a safe central
// margin in the 1200x630 frame. {summary} trails as optional article context so
// an empty summary leaves no dangling label — just a harmless trailing space.
const defaultOGPrompt = `Minimal abstract background art for a blog article titled "{title}". ` +
	`Give it a clear focal subject with balanced, even exposure — not overexposed, avoid a blown-out center. ` +
	`No text, no words, no letters, no typography. ` +
	`1200x630 composition with a clear safe central margin for an overlaid title. {summary}`

// OGPrompt renders an OG prompt template, substituting the three fixed
// placeholders. An empty/blank template falls back to the built-in default.
func OGPrompt(tmpl, title, summary, brand string) string {
	if strings.TrimSpace(tmpl) == "" {
		tmpl = defaultOGPrompt
	}
	r := strings.NewReplacer("{title}", title, "{summary}", summary, "{brand}", brand)
	return r.Replace(tmpl)
}
