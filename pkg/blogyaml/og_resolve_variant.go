//ff:func feature=blogyaml type=schema control=sequence
//ff:what 안 1개의 상속 병합 — 상위 기본값 위에 non-nil 필드만 오버라이드해 확정 OGVariantSpec으로 환원
package blogyaml

// ogResolveVariant merges one declared variant over the site-wide defaults:
// nil fields inherit, non-nil fields (including explicit false/"") override.
func ogResolveVariant(base OGVariantSpec, d OGVariant) OGVariantSpec {
	out := OGVariantSpec{Name: d.Name, Model: base.Model, Overlay: base.Overlay, Prompt: base.Prompt}
	if d.Model != nil {
		out.Model = *d.Model
	}
	if d.Overlay != nil {
		out.Overlay = *d.Overlay
	}
	if d.Prompt != nil {
		out.Prompt = *d.Prompt
	}
	return out
}
