//ff:func feature=content type=parser control=sequence
//ff:what 마크다운 바이트에서 "---" 구분 YAML front matter와 본문을 분리 디코드 — 블록이 없거나 깨지면 false
package content

import (
	"bytes"

	"gopkg.in/yaml.v3"
)

// parseFrontMatter decodes the leading "---\n...\n---" YAML block and returns
// the remaining body. It returns ok=false when the block is absent,
// unterminated or invalid YAML.
func parseFrontMatter(data []byte) (fm frontMatter, body string, ok bool) {
	if !bytes.HasPrefix(data, []byte("---\n")) {
		return fm, "", false
	}
	rest := data[4:]
	end := bytes.Index(rest, []byte("\n---"))
	if end < 0 {
		return fm, "", false
	}
	if err := yaml.Unmarshal(rest[:end+1], &fm); err != nil {
		return fm, "", false
	}
	after := rest[end+len("\n---"):]
	if nl := bytes.IndexByte(after, '\n'); nl >= 0 {
		after = after[nl+1:]
	} else {
		after = nil
	}
	return fm, string(after), true
}
