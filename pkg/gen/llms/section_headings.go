//ff:func feature=gen type=generator control=iteration dimension=1
//ff:what 정렬된 발행 글에서 실제 출현할 섹션 그룹 헤딩 텍스트 집합을 산출 — pinned group의 섹션 합류 판정 기준
package llms

import "github.com/park-jun-woo/abloq/pkg/blogyaml"

// sectionHeadings collects the heading texts the post groups will emit, so
// pinned entries can be merged into a section group only when that heading
// actually appears (pinned for empty sections render their own group instead).
func sectionHeadings(b *blogyaml.Blog, sorted []Post, multi bool) map[string]bool {
	heads := map[string]bool{}
	for _, p := range sorted {
		heads[headingText(b, p.Lang, p.Section, multi)] = true
	}
	return heads
}
