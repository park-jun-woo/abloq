//ff:func feature=gate type=parser control=iteration dimension=1 topic=evidence
//ff:what git HEAD 원본 대비 신규 수치 주장만 선별 — 원본에 같은 정규화 문장이 있으면 제외, 미변경 글은 빈 목록
//ff:why citation-exists의 newCitations와 같은 기준선 원칙 — 게이트는 이번 작업이 추가한 주장만 판정하고, 기존 코퍼스 감사(역사적 주장)는 Phase010 스캐너 소관 (Phase006 도그푸드에서 720편 전수 판정 버그로 분류되어 도입)
package gate

// newClaims returns the numeric claims whose normalized text is absent from
// the article's git HEAD baseline. A new article (Base nil) claims everything
// anew; an unchanged article (Base == Doc) claims nothing new. The gate
// judges only these — auditing pre-existing claims is the claim scanner's job.
func newClaims(a *Article) []Claim {
	if a.Base == a.Doc {
		return nil
	}
	base := map[string]bool{}
	if a.Base != nil {
		for _, c := range DetectClaims(a.Base) {
			base[normText(c.Text)] = true
		}
	}
	var out []Claim
	for _, c := range DetectClaims(a.Doc) {
		if base[normText(c.Text)] {
			continue
		}
		out = append(out, c)
	}
	return out
}
