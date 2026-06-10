//ff:func feature=blogyaml type=rule control=sequence
//ff:what 검증 룰 6종(lang-bcp47/heading-default-lang/sections-empty/threshold-range/baseurl-format/crawlers-policy)을 순서대로 실행
package blogyaml

// Validate runs all schema v1 validation rules and returns the collected diagnostics.
func Validate(filename string, b *Blog, idx lineIndex) []Diagnostic {
	var diags []Diagnostic
	diags = append(diags, ruleLangBCP47(filename, b, idx)...)
	diags = append(diags, ruleHeadingDefaultLang(filename, b, idx)...)
	diags = append(diags, ruleSectionsEmpty(filename, b, idx)...)
	diags = append(diags, ruleThresholdRange(filename, b, idx)...)
	diags = append(diags, ruleBaseURLFormat(filename, b, idx)...)
	diags = append(diags, ruleCrawlersPolicy(filename, b, idx)...)
	return diags
}
