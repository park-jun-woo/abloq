//ff:func feature=gate type=frame control=sequence
//ff:what 게이트 룰 레지스트리 — 구조 7룰 + front-matter-schema/slug-consistency/honest-lastmod/hreflang-complete (11룰, 고정 순서)
package gate

// Rules returns the gate's rule catalog in execution order. The first seven
// are the structure rules ported from the parkjunwoo.com section-order quest,
// parameterized by blog.yaml structure.
func Rules() []Rule {
	return []Rule{
		{ID: "image-first", Desc: "first content line is the main image ![..](..)", Check: ruleImageFirst},
		{ID: "image-attribution", Desc: "an italic attribution line follows the main image", Check: ruleImageAttribution},
		{ID: "section-order", Desc: "recognized sections in canonical relative order", Check: ruleSectionOrder},
		{ID: "section-preserved", Desc: "no recognized section was dropped vs the baseline", Check: ruleSectionPreserved},
		{ID: "body-lossless", Desc: "every baseline body line survives (multiset subset)", Check: ruleBodyLossless},
		{ID: "front-matter-intact", Desc: "front matter unchanged vs the baseline (lastmod update allowed)", Check: ruleFrontMatterIntact},
		{ID: "heading-canonical", Desc: "recognized sections use ## level headings", Check: ruleHeadingCanonical},
		{ID: "front-matter-schema", Desc: "required front matter fields present with valid types", Check: ruleFrontMatterSchema},
		{ID: "slug-consistency", Desc: "same slug across all language versions, no missing language", Check: ruleSlugConsistency},
		{ID: "honest-lastmod", Desc: "lastmod updates require a meaningful body diff (and queue membership)", Check: ruleHonestLastmod},
		{ID: "hreflang-complete", Desc: "built pages cross-reference every language version via hreflang", Check: ruleHreflangComplete},
	}
}
