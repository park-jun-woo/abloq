//ff:func feature=gate type=rule control=sequence topic=baseline
//ff:what 글 1편의 honest-lastmod 판정 — lastmod 미변경/원본 없음은 통과, 토큰 diff 미달과 큐 미등재를 진단
package gate

import (
	"fmt"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// honestLastmodDiags judges one article's lastmod update. The minor-diff cheese
// is rejected: after whitespace/punctuation normalization the token diff must
// reach geo.min_meaningful_diff (~one sentence by default).
func honestLastmodDiags(t *Target, a *Article) []blogyaml.Diagnostic {
	if a.Base == nil || a.Base == a.Doc {
		return nil
	}
	oldMod := fmLineValue(a.Base.FrontMatter, "lastmod")
	newMod := fmLineValue(a.Doc.FrontMatter, "lastmod")
	if newMod == oldMod || newMod == "" {
		return nil
	}
	line := fmKeyLine(a.Doc.FrontMatter, "lastmod")
	var diags []blogyaml.Diagnostic
	diff := TokenDiff(Tokens(a.Base.Body), Tokens(a.Doc.Body))
	if diff < t.Blog.Geo.MinMeaningfulDiff {
		diags = append(diags, blogyaml.Diagnostic{File: a.Path, Line: line, Rule: "honest-lastmod",
			Message: fmt.Sprintf("lastmod updated but body token diff %d < min_meaningful_diff %d", diff, t.Blog.Geo.MinMeaningfulDiff)})
	}
	if !queueAllows(t.Dir, a.Lang+"/"+a.Section+"/"+a.Slug) {
		diags = append(diags, blogyaml.Diagnostic{File: a.Path, Line: line, Rule: "honest-lastmod",
			Message: "lastmod updated but article is not in the freshness queue (quests/queue/)"})
	}
	return diags
}
