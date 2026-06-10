//ff:func feature=gate type=rule control=sequence
//ff:what 본문 라인 인덱스를 파일 라인 번호로 변환 — 음수(본문 없음)는 1
package gate

// bodyLine converts a BodyLines index into a 1-based file line number.
func bodyLine(d *Doc, idx int) int {
	if idx < 0 {
		return 1
	}
	return d.BodyStart + idx
}
