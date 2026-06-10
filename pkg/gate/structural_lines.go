//ff:func feature=gate type=parser control=iteration dimension=1 topic=lossless
//ff:what 구조 편집이 허용된 본문 라인 인덱스 집합 — 메인 이미지, 저작자 표기, 인식된 섹션 헤딩(레벨 불문)
package gate

// structuralLines collects the BodyLines indexes the structure rules own:
// the main image, its attribution and every recognized section heading.
func structuralLines(d *Doc) map[int]bool {
	skip := map[int]bool{}
	if d.FirstIsImage {
		skip[d.FirstContentLine] = true
	}
	if d.AttribLine >= 0 {
		skip[d.AttribLine] = true
	}
	for _, s := range d.Sections {
		skip[s.Line] = true
	}
	for _, s := range d.BadLevel {
		skip[s.Line] = true
	}
	return skip
}
