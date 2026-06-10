//ff:func feature=postbuild type=parser control=sequence
//ff:what 선두 "---" front matter 블록을 본문과 분리 — CRLF 정규화, 블록이 없으면 전체를 본문으로 반환
package postbuild

import "strings"

// splitFrontMatter separates a leading `---\n...\n---\n` block from the body.
// Without a well-formed block, fm is empty and body is the whole content.
func splitFrontMatter(content string) (fm, body string) {
	content = strings.ReplaceAll(content, "\r\n", "\n")
	content = strings.ReplaceAll(content, "\r", "")
	if !strings.HasPrefix(content, "---\n") {
		return "", content
	}
	rest := content[len("---\n"):]
	idx := strings.Index(rest, "\n---")
	if idx < 0 {
		return "", content
	}
	fm = rest[:idx]
	after := rest[idx+len("\n---"):]
	if nl := strings.IndexByte(after, '\n'); nl >= 0 {
		after = after[nl+1:]
	} else {
		after = ""
	}
	return fm, after
}
