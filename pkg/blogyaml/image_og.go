//ff:type feature=blogyaml type=schema
//ff:what blog.yaml image.og 블록 — 사이트 공통 provider/model/overlay/prompt 선언 + 이름 있는 안(variant) 프리셋 목록
//ff:why 비주얼 아이덴티티는 자주 안 바뀌는 사이트 단위 선언이라 SSOT가 거처 — API 키는 절대 여기 들어오지 않는다(env 전용) (BUG002)
package blogyaml

// ImageOG is the site-wide OG generation declaration. Provider "" means
// local (the deterministic text card). Model "" means the provider default.
// Prompt is a template with {title}/{summary}/{brand} placeholders ("" uses
// the built-in template, see OGPrompt). Variants are named presets that
// inherit unset fields from this block (see ResolvedVariants).
type ImageOG struct {
	Provider string      `yaml:"provider" json:"provider"`
	Model    string      `yaml:"model" json:"model"`
	Overlay  bool        `yaml:"overlay" json:"overlay"`
	Prompt   string      `yaml:"prompt" json:"prompt"`
	Variants []OGVariant `yaml:"variants" json:"variants"`
}
