//ff:func feature=sitesyaml type=rule control=sequence
//ff:what [sites-empty] sites 리스트가 1개 이상인지 검증 — 키 부재·빈 리스트 모두 거부 (사이트 없는 인스턴스는 선언 오류)
package sitesyaml

// ruleSitesEmpty rejects an empty or absent sites list: an instance that
// serves no site is a declaration mistake, not a valid configuration.
func ruleSitesEmpty(filename string, s *Sites, idx lineIndex) []Diagnostic {
	if len(s.Sites) > 0 {
		return nil
	}
	return []Diagnostic{{
		File: filename, Line: lineOf(idx, "sites"), Rule: "sites-empty",
		Message: "sites must declare at least one site",
	}}
}
