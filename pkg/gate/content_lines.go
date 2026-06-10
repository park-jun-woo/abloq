//ff:func feature=gate type=parser control=iteration dimension=1 topic=lossless
//ff:what 본문 무손실 비교용 multiset 키 추출 — 메인 이미지/저작자 표기/인식 섹션 헤딩 라인을 제외한 정규화 비공백 라인
package gate

// ContentLines returns the multiset keys for the body-lossless check: every
// non-blank body line, normalized, excluding the main image line, the
// attribution line and recognized section-heading lines. Excluding those
// permits the allowed structural edits while catching deleted or altered text.
func ContentLines(d *Doc) []string {
	skip := structuralLines(d)
	var out []string
	for i, raw := range d.BodyLines {
		if skip[i] {
			continue
		}
		ln := normText(raw)
		if ln == "" {
			continue
		}
		out = append(out, ln)
	}
	return out
}
