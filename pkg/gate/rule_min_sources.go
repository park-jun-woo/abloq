//ff:func feature=gate type=rule control=iteration dimension=1 topic=evidence
//ff:what [min-sources] 글당 sources 섹션의 출처 항목 수 ≥ geo.min_sources 검증 — 섹션 누락도 진단
package gate

import (
	"fmt"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// ruleMinSources requires each article's sources section to list at least
// geo.min_sources entries. Skipped when the threshold is zero or the blog's
// structure declares no sources section.
func ruleMinSources(t *Target) []blogyaml.Diagnostic {
	min := t.Blog.Geo.MinSources
	if min <= 0 || !orderHas(t.Blog, "sources") {
		return nil
	}
	var diags []blogyaml.Diagnostic
	for _, a := range t.Articles {
		n, head, found := sourceCount(a.Doc)
		if found && n >= min {
			continue
		}
		line, msg := 1, fmt.Sprintf("sources section missing — geo.min_sources requires >= %d source(s)", min)
		if found {
			line = bodyLine(a.Doc, head)
			msg = fmt.Sprintf("sources section lists %d source(s), geo.min_sources requires >= %d", n, min)
		}
		diags = append(diags, blogyaml.Diagnostic{File: a.Path, Line: line, Rule: "min-sources", Message: msg})
	}
	return diags
}
