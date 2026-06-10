//ff:func feature=claudemd type=generator control=sequence
//ff:what CLAUDE.md 헤더 — 블로그 정체(제목/저자/baseURL/언어/섹션)와 재생성 안내
package claudemd

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// headerSection renders the manual header and blog identity from blog.yaml.
func headerSection(b *blogyaml.Blog) string {
	langs := strings.Join(b.Languages, ", ")
	def := ""
	if len(b.Languages) > 0 {
		def = b.Languages[0]
	}
	var sb strings.Builder
	fmt.Fprintf(&sb, "# CLAUDE.md — %s 운영 매뉴얼\n\n", b.Site.Title)
	sb.WriteString("> abloq가 blog.yaml에서 생성한 에이전트 운영 매뉴얼이다. 직접 편집하지 말 것 —\n")
	sb.WriteString("> blog.yaml을 고친 뒤 `abloq claudemd .`로 재생성한다.\n\n")
	sb.WriteString("## 블로그\n\n")
	fmt.Fprintf(&sb, "- baseURL: %s\n", b.Site.BaseURL)
	fmt.Fprintf(&sb, "- 저자: %s\n", b.Site.Author)
	fmt.Fprintf(&sb, "- 언어: %s — 기본 언어는 %q (languages 첫 항목)\n", langs, def)
	fmt.Fprintf(&sb, "- 섹션: %s\n\n", strings.Join(b.Sections, ", "))
	return sb.String()
}
