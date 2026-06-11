//ff:type feature=blogyaml type=schema
//ff:what blog.yaml geo.llms_txt 정규화 스펙 — mode(auto/manual/off), 언어 스코프, 포지셔닝 header, pinned 엔트리, 섹션 라벨, summary 길이 상한
//ff:why 문자열 단축형(auto|manual|off)과 객체 폼의 union을 파싱 직후 이 구조체 하나로 수렴 — 이후 코드는 union을 모른다 (BUG001)
package blogyaml

// LlmsTxtSpec is the normalized geo.llms_txt declaration. blog.yaml accepts
// either a string shorthand ("auto" | "manual" | "off") or this object form;
// both decode into this one struct (JSON serialization follows it too).
// Languages holds ["base"] (default-language scope, the default), ["all"]
// (every declared language) or an explicit subset of declared languages.
// MaxSummary caps the per-entry description in runes (0 = unlimited).
type LlmsTxtSpec struct {
	Mode          string            `yaml:"mode" json:"mode"`
	Languages     []string          `yaml:"languages" json:"languages"`
	Header        string            `yaml:"header" json:"header"`
	Pinned        []LlmsPinned      `yaml:"pinned" json:"pinned"`
	SectionLabels map[string]string `yaml:"section_labels" json:"section_labels"`
	MaxSummary    int               `yaml:"max_summary" json:"max_summary"`
}
