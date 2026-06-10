//ff:func feature=gate type=rule control=sequence
//ff:what 글 1편의 front matter 스키마 진단 — 블록 존재, title 비공백 문자열, date/lastmod 파싱과 순서, tags 비공백 목록
package gate

import "github.com/park-jun-woo/abloq/pkg/blogyaml"

// fmSchemaDiags validates one article's front matter against the required
// schema: title, date, lastmod and tags must exist with valid types, and
// lastmod must not precede date.
func fmSchemaDiags(a *Article) []blogyaml.Diagnostic {
	diag := func(line int, msg string) blogyaml.Diagnostic {
		return blogyaml.Diagnostic{File: a.Path, Line: line, Rule: "front-matter-schema", Message: msg}
	}
	if !a.Doc.HasFM {
		return []blogyaml.Diagnostic{diag(1, "front matter block missing or malformed")}
	}
	m, ok := fmMap(a.Doc.FrontMatter)
	if !ok {
		return []blogyaml.Diagnostic{diag(1, "front matter is not valid YAML")}
	}
	var diags []blogyaml.Diagnostic
	if s, _ := m["title"].(string); s == "" {
		diags = append(diags, diag(fmKeyLine(a.Doc.FrontMatter, "title"), "title must be a non-empty string"))
	}
	date, dateOK := parseFMTime(m["date"])
	if !dateOK {
		diags = append(diags, diag(fmKeyLine(a.Doc.FrontMatter, "date"), "date must be a RFC3339 or YYYY-MM-DD value"))
	}
	lastmod, lastmodOK := parseFMTime(m["lastmod"])
	if !lastmodOK {
		diags = append(diags, diag(fmKeyLine(a.Doc.FrontMatter, "lastmod"), "lastmod must be a RFC3339 or YYYY-MM-DD value"))
	}
	if dateOK && lastmodOK && lastmod.Before(date) {
		diags = append(diags, diag(fmKeyLine(a.Doc.FrontMatter, "lastmod"), "lastmod must not precede date"))
	}
	if tags, _ := m["tags"].([]any); len(tags) == 0 {
		diags = append(diags, diag(fmKeyLine(a.Doc.FrontMatter, "tags"), "tags must be a non-empty list"))
	}
	return diags
}
