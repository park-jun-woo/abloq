//ff:func feature=insight type=rule control=iteration dimension=1
//ff:what claim 1건의 출현 판정 — anchors(동의어 목록) 중 하나라도 폴딩된 본문의 부분문자열이면 true
package insight

import "strings"

// claimFound reports whether any anchor of c occurs in the folded body.
// foldedBody must already be normalized by fold; anchorless claims are never
// found (the validator warns about them).
func claimFound(foldedBody string, c Claim) bool {
	for _, a := range c.Anchors {
		if strings.Contains(foldedBody, fold(a)) {
			return true
		}
	}
	return false
}
