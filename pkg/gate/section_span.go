//ff:func feature=gate type=parser control=iteration dimension=1 topic=evidence
//ff:what 인식 섹션 1개의 본문 라인 범위(헤딩 라인, 다음 섹션 직전) 탐색 — 없으면 false
package gate

// sectionSpan locates a recognized section's body-line range: start is the
// heading's BodyLines index, end is the next recognized section's heading
// index (or the body end). ok is false when the section is absent.
func sectionSpan(d *Doc, key string) (start, end int, ok bool) {
	for i, s := range d.Sections {
		if s.Key != key {
			continue
		}
		end = len(d.BodyLines)
		if i+1 < len(d.Sections) {
			end = d.Sections[i+1].Line
		}
		return s.Line, end, true
	}
	return 0, 0, false
}
