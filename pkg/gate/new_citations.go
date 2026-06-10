//ff:func feature=gate type=parser control=iteration dimension=1 topic=evidence
//ff:what git HEAD 원본 대비 신규 인용만 선별 — 원본에 없는 URL의 인용, 미변경 글(Base==Doc)은 빈 목록
package gate

// newCitations returns the citations whose URL is absent from the article's
// git HEAD baseline. A new article (Base nil) cites everything anew; an
// unchanged article (Base == Doc) cites nothing new. The gate verifies only
// these — re-checking existing citations is the Phase010 scanner's job.
func newCitations(a *Article) []Citation {
	if a.Base == a.Doc {
		return nil
	}
	base := map[string]bool{}
	if a.Base != nil {
		for _, c := range Citations(a.Base) {
			base[c.URL] = true
		}
	}
	var out []Citation
	for _, c := range Citations(a.Doc) {
		if base[c.URL] {
			continue
		}
		out = append(out, c)
	}
	return out
}
