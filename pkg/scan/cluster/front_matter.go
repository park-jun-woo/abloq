//ff:type feature=scan type=schema topic=cluster
//ff:what 클러스터 스캐너가 읽는 front matter 부분집합 — date/slug/draft/tags만 디코드, 나머지 키는 무시
package cluster

// frontMatter is the cluster scanner's front matter subset; unknown keys are
// ignored. Slug overrides the file stem (the same contract as the posts
// index and the gate's effSlug).
type frontMatter struct {
	Date  string   `yaml:"date"`
	Slug  string   `yaml:"slug"`
	Draft bool     `yaml:"draft"`
	Tags  []string `yaml:"tags"`
}
