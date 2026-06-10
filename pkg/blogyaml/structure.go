//ff:type feature=blogyaml type=schema
//ff:what blog.yaml structure 섹션 — 글의 정규 섹션 순서(order)와 언어별 헤딩 맵(headings), 구조 게이트 룰의 입력
package blogyaml

// Structure declares the canonical article layout.
// Headings maps heading key -> language code -> localized heading text.
type Structure struct {
	Order    []string                     `yaml:"order" json:"order"`
	Headings map[string]map[string]string `yaml:"headings" json:"headings"`
}
