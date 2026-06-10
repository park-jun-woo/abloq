//ff:func feature=gate type=parser control=iteration dimension=1 topic=evidence
//ff:what 문단 1개에서 수치 주장 라인을 수집 — 문단 전체에 출처 링크가 있으면 전 주장을 Sourced로 표시
package gate

import "strings"

// paraClaims detects the numeric claims in one paragraph. Sourcing is judged
// at paragraph granularity: a source link anywhere in the paragraph covers
// every claim it contains.
func paraClaims(d *Doc, p claimPara) []Claim {
	sourced := hasSourceLink(strings.Join(p.texts, "\n"))
	var out []Claim
	for i, txt := range p.texts {
		if !isClaimLine(txt) {
			continue
		}
		out = append(out, Claim{Line: bodyLine(d, p.lines[i]), Text: strings.TrimSpace(txt), Sourced: sourced})
	}
	return out
}
