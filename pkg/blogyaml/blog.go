//ff:type feature=blogyaml type=schema
//ff:what blog.yaml 스키마 v1 최상위 — 블로그 하나의 전 선언 (site/languages/sections/structure/geo/image/deploy)
//ff:why 후속 Phase 키(min_meaningful_diff, taxonomy 등)는 섹션 구조체에 필드만 추가하면 되도록 평탄한 섹션 분할 유지
package blogyaml

// Blog is the root of blog.yaml schema v1. Unknown keys are rejected (strict parsing).
type Blog struct {
	Site      Site      `yaml:"site" json:"site"`
	Languages []string  `yaml:"languages" json:"languages"` // first entry = default language
	Sections  []string  `yaml:"sections" json:"sections"`
	Structure Structure `yaml:"structure" json:"structure"`
	Geo       Geo       `yaml:"geo" json:"geo"`
	Image     Image     `yaml:"image" json:"image"`
	Deploy    Deploy    `yaml:"deploy" json:"deploy"`
}
