//ff:func feature=blogyaml type=rule control=sequence
//ff:what [sections-empty] sections가 비어있지 않은지 검증
package blogyaml

// ruleSectionsEmpty requires at least one content section.
func ruleSectionsEmpty(filename string, b *Blog, idx lineIndex) []Diagnostic {
	if len(b.Sections) > 0 {
		return nil
	}
	return []Diagnostic{{
		File: filename, Line: lineOf(idx, "sections"), Rule: "sections-empty",
		Message: "sections must contain at least one section",
	}}
}
