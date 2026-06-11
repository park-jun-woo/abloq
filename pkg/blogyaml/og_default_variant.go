//ff:func feature=blogyaml type=schema control=sequence
//ff:what 기본 설정 안 생성 — 사이트 공통 model/overlay/prompt 그대로, 이름은 예약명 "default" (variants 미선언·미지정 경로의 안)
package blogyaml

// DefaultVariant builds the default-settings candidate from the site-wide
// block. Its name is the reserved "default" — which is why no declared
// variant may use that name (og-variant-name rule).
func (o ImageOG) DefaultVariant() OGVariantSpec {
	return OGVariantSpec{Name: "default", Model: o.Model, Overlay: o.Overlay, Prompt: o.Prompt}
}
