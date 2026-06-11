//ff:func feature=cli type=command control=iteration dimension=1
//ff:what og 프롬프트용 summary 결선 — 검증된 blog의 base 언어(Languages[0]) 유효 slug == argv slug 매칭으로 1건이면 그 summary, 0/다건이면 빈값+진단 1줄
//ff:why og는 글 없이도 도는 명령 — 미발견은 에러 아닌 폴백이고, blog.yaml 부재(nil blog)면 IndexEntries(root,nil) 호출 시 nil 역참조 패닉이라 결선을 아예 건너뛴다
package main

import (
	"fmt"
	"io"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
	"github.com/park-jun-woo/abloq/pkg/content"
)

// resolveOGSummary maps the argv slug to a front matter summary using the
// canonical indexer's effective-slug computation (front matter slug overrides
// the file stem). It scopes to the base language (Languages[0]) so the same
// SSOT article yields one summary, matching the llms.txt single-source rule.
// A nil/zero-value blog (no blog.yaml) skips resolution and returns "" — it
// must NOT call IndexEntries(root, nil), which would dereference b.Languages.
func resolveOGSummary(out io.Writer, root string, b *blogyaml.Blog, slug string) string {
	if b == nil || len(b.Languages) == 0 {
		return ""
	}
	base := b.Languages[0]
	var matches []string
	for _, e := range content.IndexEntries(root, b) {
		if e.Lang == base && e.Slug == slug {
			matches = append(matches, e.Summary)
		}
	}
	if len(matches) == 1 {
		return matches[0]
	}
	if len(matches) == 0 {
		fmt.Fprintf(out, "summary 미적용: base 언어(%s)에서 slug %q 글을 찾지 못함\n", base, slug)
	} else {
		fmt.Fprintf(out, "summary 미적용: base 언어(%s)에서 slug %q가 %d건으로 모호함\n", base, slug, len(matches))
	}
	return ""
}
