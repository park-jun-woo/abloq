//ff:func feature=blogyaml type=rule control=sequence
//ff:what 검증 룰 15종(lang-bcp47/heading-default-lang/sections-empty/threshold-range/priority-weights-range/baseurl-format/crawlers-policy/taxonomy-unique/llmstxt-mode/llmstxt-languages/llmstxt-pinned/llmstxt-labels/llmstxt-max-summary/og-provider/og-variant-name)을 순서대로 실행
package blogyaml

// Validate runs all schema v1 validation rules and returns the collected diagnostics.
func Validate(filename string, b *Blog, idx lineIndex) []Diagnostic {
	var diags []Diagnostic
	diags = append(diags, ruleLangBCP47(filename, b, idx)...)
	diags = append(diags, ruleHeadingDefaultLang(filename, b, idx)...)
	diags = append(diags, ruleSectionsEmpty(filename, b, idx)...)
	diags = append(diags, ruleThresholdRange(filename, b, idx)...)
	diags = append(diags, rulePriorityWeights(filename, b, idx)...)
	diags = append(diags, ruleBaseURLFormat(filename, b, idx)...)
	diags = append(diags, ruleCrawlersPolicy(filename, b, idx)...)
	diags = append(diags, ruleTaxonomyUnique(filename, b, idx)...)
	diags = append(diags, ruleLlmsTxtMode(filename, b, idx)...)
	diags = append(diags, ruleLlmsTxtLanguages(filename, b, idx)...)
	diags = append(diags, ruleLlmsTxtPinned(filename, b, idx)...)
	diags = append(diags, ruleLlmsTxtLabels(filename, b, idx)...)
	diags = append(diags, ruleLlmsTxtMaxSummary(filename, b, idx)...)
	diags = append(diags, ruleOGProvider(filename, b, idx)...)
	diags = append(diags, ruleOGVariantName(filename, b, idx)...)
	return diags
}
