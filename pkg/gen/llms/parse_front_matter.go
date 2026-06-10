//ff:func feature=gen type=parser control=sequence
//ff:what 마크다운 바이트에서 "---" 구분 YAML front matter를 비-strict로 디코드, 없거나 깨지면 false
package llms

import (
	"bytes"

	"gopkg.in/yaml.v3"
)

// parseFrontMatter decodes the leading "---\n...\n---" YAML block.
// It returns false when the block is absent, unterminated or invalid YAML.
func parseFrontMatter(data []byte) (frontMatter, bool) {
	var fm frontMatter
	if !bytes.HasPrefix(data, []byte("---\n")) {
		return fm, false
	}
	rest := data[4:]
	end := bytes.Index(rest, []byte("\n---"))
	if end < 0 {
		return fm, false
	}
	if err := yaml.Unmarshal(rest[:end+1], &fm); err != nil {
		return fm, false
	}
	return fm, true
}
