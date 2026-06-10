//ff:func feature=gate type=parser control=iteration dimension=2 topic=evidence
//ff:what 인식된 sources 섹션 내부의 본문 라인 인덱스 집합 — 출처 명기 목록행은 수치 주장 검출 면제
//ff:why Phase010 parkjunwoo.com 사본 1회전에서 출처 섹션의 인용 목록행 자체가 무출처 수치 주장으로 검출되는 오탐 발견 — blog.yaml 헤딩 맵으로 인식된 섹션만 면제(헤딩 미선언 언어는 면제 없음, 보수적)
package gate

// sourcesLines collects the BodyLines indexes inside every recognized sources
// section (heading line through the next recognized section). Lines listed
// there ARE the citations, so the claim detector exempts them. Recognition
// comes from blog.yaml's per-language heading map: when the language declares
// no sources heading, nothing is recognized and nothing is exempt.
func sourcesLines(d *Doc) map[int]bool {
	in := map[int]bool{}
	for i, s := range d.Sections {
		if s.Key != "sources" {
			continue
		}
		end := len(d.BodyLines)
		if i+1 < len(d.Sections) {
			end = d.Sections[i+1].Line
		}
		for ln := s.Line; ln < end; ln++ {
			in[ln] = true
		}
	}
	return in
}
