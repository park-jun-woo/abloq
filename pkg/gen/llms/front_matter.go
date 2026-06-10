//ff:type feature=gen type=schema
//ff:what front matter 최소 스키마 — llms.txt 생성에 필요한 title/date/draft/description만 디코드
//ff:why 전수 콘텐츠 인덱서는 Phase007 소관 — 여기서는 비-strict 최소 파싱으로 멱등 목록 생성만 담당
package llms

// frontMatter is the minimal front matter subset; unknown keys are ignored.
type frontMatter struct {
	Title       string `yaml:"title"`
	Date        string `yaml:"date"`
	Draft       bool   `yaml:"draft"`
	Description string `yaml:"description"`
}
