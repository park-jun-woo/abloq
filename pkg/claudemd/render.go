//ff:func feature=claudemd type=generator control=sequence
//ff:what blog.yaml에서 CLAUDE.md 운영 매뉴얼 바이트를 렌더 — 6개 섹션 고정 순서, 같은 입력이면 바이트 동일
package claudemd

import "github.com/park-jun-woo/abloq/pkg/blogyaml"

// Render produces CLAUDE.md deterministically from blog.yaml: the agent
// operations manual (publish procedure, gates, layout, forbidden moves).
func Render(b *blogyaml.Blog) []byte {
	s := headerSection(b) +
		dirsSection(b) +
		structureSection(b) +
		publishSection(b) +
		commandsSection() +
		forbiddenSection(b)
	return []byte(s)
}
