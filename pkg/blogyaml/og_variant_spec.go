//ff:type feature=blogyaml type=schema
//ff:what 상속 병합이 끝난 OG 안 1개 — 이름/모델/오버레이/프롬프트 전 필드 확정값, optional 표현이 사라진 정규화 산물
package blogyaml

// OGVariantSpec is one fully merged OG candidate: every field holds its
// effective value, no pointers survive past normalization. Downstream code
// (cmd, pkg/img) only ever sees this form.
type OGVariantSpec struct {
	Name    string
	Model   string
	Overlay bool
	Prompt  string
}
