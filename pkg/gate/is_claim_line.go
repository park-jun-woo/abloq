//ff:func feature=gate type=parser control=sequence topic=evidence
//ff:what 라인 1개의 수치 주장 여부 판정 — 인라인 코드 제거 후 수치+단위 패턴과 단정 어휘가 함께 있으면 true
//ff:why 오탐 0을 미탐보다 우선 — 단위 없는 수(연도/버전/장 번호)와 단정 어휘 없는 수치는 주장으로 보지 않는다
package gate

import "regexp"

var (
	reInlineCode = regexp.MustCompile("`[^`]*`")
	// a number immediately followed by a measurement unit
	reClaimNum = regexp.MustCompile(`\d[\d,]*(?:\.\d+)?\s*(?:%|％|퍼센트|배|ms|밀리초|µs|ns|초|시간|명|건|회|위|원|달러|GB|MB|KB|TB|GiB|MiB|fps|qps|rps|°C|℃|kg|km|pp\b|(?i:x\b|times\b|percent\b|points?\b))`)
	// an assertive change/achievement verb in the same line (≈ sentence)
	reClaimAssert = regexp.MustCompile(`증가|감소|상승|하락|향상|개선|단축|절감|절약|줄었|줄어|늘었|늘어|빨라|느려|높아|높았|낮아|낮았|달성|기록|차지|도달|돌파|급증|급감|(?i:increase|decrease|improv|reduc|faster|slower|grew|growth|dropped|fell\b|rose\b|gained|achiev|reached|accounts? for|boost|doubled|tripled|halved|cut\b)`)
)

// isClaimLine reports whether ln states a numeric claim: a number+unit token
// and an assertive verb in the same line. Inline code spans are stripped first.
func isClaimLine(ln string) bool {
	ln = reInlineCode.ReplaceAllString(ln, "")
	return reClaimNum.MatchString(ln) && reClaimAssert.MatchString(ln)
}
