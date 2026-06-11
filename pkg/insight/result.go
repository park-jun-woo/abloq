//ff:type feature=insight type=schema
//ff:what 매칭 결과 — 글 실위치 섹션, anchors 출현 claim id 목록, 미출현 claim 목록(REVIEW 보조)
package insight

// Result is the outcome of matching one insight spec against one article body.
// Missing is a REVIEW aid only — anchor occurrence does not prove
// correspondence, and absence does not prove the claim is unaddressed.
type Result struct {
	Section string   `json:"section"`
	Found   []string `json:"found"`
	Missing []Claim  `json:"missing"`
}
