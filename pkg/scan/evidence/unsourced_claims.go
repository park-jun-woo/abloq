//ff:func feature=scan type=parser control=iteration dimension=1 topic=evidence
//ff:what 글 1편의 무출처 수치 주장 수집 — gate.DetectClaims 재사용, claims_ignore 예외 동일 적용, 위치는 저장소 상대 path:line
package evidence

import (
	"strconv"

	"github.com/park-jun-woo/abloq/pkg/gate"
)

// unsourcedClaims maps one article's unsourced numeric claims to their queue
// payload form. Detection and the claims_ignore exemption are the gate's own
// (no reimplementation); Loc keeps the repository-relative article path so
// CLI and endpoint emit identical payloads.
func unsourcedClaims(a *gate.Article) []ClaimRef {
	if gate.ClaimsExempt(a) {
		return nil
	}
	var refs []ClaimRef
	for _, c := range gate.DetectClaims(a.Doc) {
		if c.Sourced {
			continue
		}
		refs = append(refs, ClaimRef{Hash: gate.HashText(c.Text), Loc: a.Path + ":" + strconv.Itoa(c.Line), Text: c.Text})
	}
	return refs
}
