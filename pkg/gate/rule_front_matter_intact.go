//ff:func feature=gate type=rule control=iteration dimension=1 topic=baseline
//ff:what [front-matter-intact] 원본(git HEAD) 대비 front matter 불변(lastmod 갱신만 허용) 검증
package gate

import "github.com/park-jun-woo/abloq/pkg/blogyaml"

// ruleFrontMatterIntact requires the front matter to stay byte-identical to
// the baseline except for an allowed lastmod update.
func ruleFrontMatterIntact(t *Target) []blogyaml.Diagnostic {
	var diags []blogyaml.Diagnostic
	for _, a := range t.Articles {
		if a.Base == nil || a.Base == a.Doc {
			continue
		}
		if !a.Doc.HasFM {
			diags = append(diags, blogyaml.Diagnostic{File: a.Path, Line: 1, Rule: "front-matter-intact",
				Message: "front matter block missing or malformed"})
			continue
		}
		if diff, ok := fmIntactDiff(a.Base.FrontMatter, a.Doc.FrontMatter); !ok {
			diags = append(diags, blogyaml.Diagnostic{File: a.Path, Line: 1, Rule: "front-matter-intact",
				Message: "front matter changed (only lastmod may change): " + trunc(diff)})
		}
	}
	return diags
}
