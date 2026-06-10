//ff:func feature=postbuild type=generator control=sequence
//ff:what 글 원문에서 노이즈 제로 .md를 렌더 — front matter 제거, "# title" 헤더 부착, AI 컨텍스트 포맷
package postbuild

import (
	"strings"

	"gopkg.in/yaml.v3"
)

// RenderMD strips the front matter and prepends "# {title}" so the served
// .md is a clean, self-titled document for AI agents.
func RenderMD(data []byte) []byte {
	fm, body := splitFrontMatter(string(data))
	var meta struct {
		Title string `yaml:"title"`
	}
	_ = yaml.Unmarshal([]byte(fm), &meta)
	body = strings.TrimLeft(body, "\n")
	if meta.Title == "" {
		return []byte(body)
	}
	return []byte("# " + meta.Title + "\n\n" + body)
}
