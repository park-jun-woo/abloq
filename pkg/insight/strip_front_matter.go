//ff:func feature=insight type=parser control=sequence
//ff:what 글 바이트에서 선두 "---" front matter 블록을 떼고 본문만 반환 — 블록이 없거나 미종결이면 전체가 본문
package insight

import "bytes"

// stripFrontMatter returns the article body without the leading YAML front
// matter block. Anchors must match the body only (design: front matter
// excluded from match semantics).
func stripFrontMatter(data []byte) string {
	if !bytes.HasPrefix(data, []byte("---\n")) {
		return string(data)
	}
	rest := data[4:]
	end := bytes.Index(rest, []byte("\n---"))
	if end < 0 {
		return string(data)
	}
	after := rest[end+len("\n---"):]
	if nl := bytes.IndexByte(after, '\n'); nl >= 0 {
		return string(after[nl+1:])
	}
	return ""
}
