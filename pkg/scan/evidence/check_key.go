//ff:func feature=scan type=generator control=sequence topic=evidence
//ff:what citation_checks 키 (url, lang, section, slug) → 맵 키 문자열 — 개행 결합 (URL·좌표에 개행 불가)
package evidence

// checkKey joins the citation_checks unique key into one map key. Newline is
// a safe separator: none of the four parts can contain one.
func checkKey(url, lang, section, slug string) string {
	return url + "\n" + lang + "\n" + section + "\n" + slug
}
