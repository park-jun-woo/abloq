//ff:type feature=content type=schema
//ff:what 인덱서가 읽는 front matter 부분집합 — title/date/lastmod/slug/draft/tags/summary/description만 디코드, 나머지 키는 무시
package content

// frontMatter is the indexer's front matter subset; unknown keys are ignored.
// Slug overrides the file stem in URLs (same contract as the gate's effSlug).
// Summary is the abloq article schema's one-line abstract; Description is the
// generic Hugo alias accepted as an override (same priority as pkg/gen/llms).
type frontMatter struct {
	Title       string   `yaml:"title"`
	Date        string   `yaml:"date"`
	Lastmod     string   `yaml:"lastmod"`
	Slug        string   `yaml:"slug"`
	Draft       bool     `yaml:"draft"`
	Tags        []string `yaml:"tags"`
	Summary     string   `yaml:"summary"`
	Description string   `yaml:"description"`
}
