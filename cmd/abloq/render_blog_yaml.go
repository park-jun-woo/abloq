//ff:func feature=init type=generator control=sequence
//ff:what init 답변에서 blog.yaml 바이트를 렌더 — 표준 crawlers 정책과 body→sources 구조 계약 포함, 결정적
package main

import (
	"fmt"
	"strings"
)

// renderBlogYAML writes the initial SSOT. The schema is blog.yaml v1; values
// not asked at init keep their documented defaults (geo thresholds, deploy).
func renderBlogYAML(o initOpts) []byte {
	var sb strings.Builder
	sb.WriteString("# blog.yaml — 이 블로그의 SSOT (abloq init 생성)\n")
	sb.WriteString("# 모든 파생물과 게이트 파라미터가 여기서 나온다. 스키마: docs/blog-yaml.md\n")
	sb.WriteString("site:\n")
	fmt.Fprintf(&sb, "  baseURL: %s\n", o.BaseURL)
	fmt.Fprintf(&sb, "  title: %q\n", o.Title)
	fmt.Fprintf(&sb, "  author: %q\n", o.Author)
	fmt.Fprintf(&sb, "\nlanguages: [%s]\n", strings.Join(o.Languages, ", "))
	fmt.Fprintf(&sb, "sections: [%s]\n", strings.Join(o.Sections, ", "))
	sb.WriteString("\nstructure:\n")
	sb.WriteString("  order: [body, sources]\n")
	sb.WriteString("  headings:\n")
	fmt.Fprintf(&sb, "    sources: { %s }\n", headingLine(o.Languages))
	sb.WriteString("\ngeo:\n")
	sb.WriteString("  crawlers:\n")
	sb.WriteString("    training: allow\n")
	sb.WriteString("    search: allow\n")
	sb.WriteString("    fetch: allow\n")
	sb.WriteString("\ndeploy:\n")
	sb.WriteString("  provider: s3-cloudfront\n")
	return []byte(sb.String())
}
