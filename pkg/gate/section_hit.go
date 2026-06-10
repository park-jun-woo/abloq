//ff:type feature=gate type=schema
//ff:what 인식된 섹션 헤딩 1회 출현 — structure.order 헤딩 키, 헤딩 레벨, 원문 텍스트, 본문 라인 인덱스
package gate

// SectionHit is one recognized section heading occurrence in an article body.
type SectionHit struct {
	Key   string // heading key from structure.order (e.g. "sources")
	Level int    // heading level (number of leading #)
	Text  string // heading text (after the #'s)
	Line  int    // index in Doc.BodyLines
}
