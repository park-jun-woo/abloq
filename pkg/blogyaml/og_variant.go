//ff:type feature=blogyaml type=schema
//ff:what image.og.variants 항목 1개 — name 필수, model/overlay/prompt는 optional 포인터(미지정=상위 상속, 명시=오버라이드)
//ff:why "미지정 vs false/빈 문자열 명시"는 zero-value로 구분 불가 — 포인터 필드로 표현하고 파싱 직후 ResolvedVariants가 병합 완료 구조체로 수렴, 이후 코드는 optional을 모른다 (Phase021 선례의 포인터 안)
package blogyaml

// OGVariant is one named OG preset as declared in blog.yaml. Nil fields
// inherit the site-wide ImageOG value; non-nil fields override it — an
// explicit `overlay: false` or `model: ""` is a real override, not absence.
// Name goes into draft filenames, so it must be URL-safe, unique and not the
// reserved "default".
type OGVariant struct {
	Name    string  `yaml:"name" json:"name"`
	Model   *string `yaml:"model" json:"model,omitempty"`
	Overlay *bool   `yaml:"overlay" json:"overlay,omitempty"`
	Prompt  *string `yaml:"prompt" json:"prompt,omitempty"`
}
