//ff:type feature=gate type=schema topic=evidence
//ff:what 수치 주장 1건 — 파일 라인, 주장 문장, 같은 문단 내 출처 링크 보유 여부 (Phase010 스캐너 공유 모델)
package gate

// Claim is one numeric-claim occurrence detected in an article body:
// a number+unit+assertion sentence. Sourced reports whether its paragraph
// carries a source link (inline http(s) link or footnote reference).
type Claim struct {
	Line    int    // 1-based file line number
	Text    string // the claim line as written (trimmed)
	Sourced bool
}
