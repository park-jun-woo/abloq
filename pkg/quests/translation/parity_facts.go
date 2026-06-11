//ff:func feature=quest type=rule control=sequence topic=lossless
//ff:what translation-parity ①~⑦ 전 검사 실행 — 헤딩/문단/이미지/코드블록/외부 링크/내부 링크/fm-mirror Fact 수집
package translation

import "github.com/park-jun-woo/reins/pkg/quest"

// parityFacts runs every structural parity check between the origin and the
// translation, in the plan's ①~⑦ order, collecting one Fact per violation.
func parityFacts(sub *Submission) []quest.Fact {
	o, t := sub.Origin.Doc, sub.Target.Articles[0].Doc
	where := sub.Article
	var facts []quest.Fact
	facts = append(facts, checkHeadings(where, o, t)...)
	facts = append(facts, checkParas(where, o, t)...)
	facts = append(facts, checkMultiset(where+"#images", "image path multiset", imagePaths(o), imagePaths(t))...)
	facts = append(facts, checkMultiset(where+"#code", "code block multiset (translation-forbidden)", codeBlocks(o), codeBlocks(t))...)
	facts = append(facts, checkMultiset(where+"#external-links", "external link URL multiset", externalLinks(o), externalLinks(t))...)
	facts = append(facts, checkInternalLinks(sub)...)
	facts = append(facts, checkFMMirror(sub)...)
	return facts
}
