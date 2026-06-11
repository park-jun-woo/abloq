//ff:func feature=sitesyaml type=rule control=sequence
//ff:what 검증 룰 5종(sites-empty/name-format/name-unique/repo-path/gsc-site-url)을 순서대로 실행
package sitesyaml

// Validate runs all schema v1 validation rules and returns the collected diagnostics.
func Validate(filename string, s *Sites, idx lineIndex) []Diagnostic {
	var diags []Diagnostic
	diags = append(diags, ruleSitesEmpty(filename, s, idx)...)
	diags = append(diags, ruleNameFormat(filename, s, idx)...)
	diags = append(diags, ruleNameUnique(filename, s, idx)...)
	diags = append(diags, ruleRepoPath(filename, s, idx)...)
	diags = append(diags, ruleGSCSiteURL(filename, s, idx)...)
	return diags
}
