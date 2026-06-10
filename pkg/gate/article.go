//ff:type feature=gate type=schema topic=baseline
//ff:what 대상 글 1편 — 언어/섹션/slug/경로와 현재 파싱본(Doc), git HEAD 원본 파싱본(Base)
package gate

// Article is one content file under gate review. Base is the parsed git HEAD
// snapshot: nil means no baseline exists (new article, or no git repository),
// and baseline-comparison rules skip the article.
type Article struct {
	Lang    string
	Section string
	Slug    string // file stem (front matter slug may override URLs — see effSlug)
	Path    string // path relative to Target.Dir (e.g. content/ko/tech/foo.md)
	Doc     *Doc
	Base    *Doc
}
