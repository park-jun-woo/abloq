//ff:type feature=gate type=schema topic=evidence
//ff:what 수치 주장 검출용 문단 1개 — 검출 대상 본문 라인 인덱스와 원문 텍스트의 병렬 목록
package gate

// claimPara is one paragraph eligible for claim detection: consecutive
// non-blank prose lines, with code, quote and heading lines already excluded.
// lines[i] is the BodyLines index of texts[i].
type claimPara struct {
	lines []int
	texts []string
}
