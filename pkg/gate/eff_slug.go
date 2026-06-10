//ff:func feature=gate type=rule control=sequence
//ff:what 글의 유효 slug 결정 — front matter slug가 있으면 그 값, 없으면 파일 어간
package gate

// effSlug returns the slug Hugo uses for the article's URL.
func effSlug(a *Article) string {
	if s := fmLineValue(a.Doc.FrontMatter, "slug"); s != "" {
		return s
	}
	return a.Slug
}
