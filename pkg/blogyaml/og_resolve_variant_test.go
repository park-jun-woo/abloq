//ff:func feature=blogyaml type=schema control=sequence
//ff:what ogResolveVariant 검증 — nil 필드는 base 상속, non-nil은 명시 false/"" 포함 오버라이드, 이름은 선언 측 채택
package blogyaml

import "testing"

func TestOGResolveVariant(t *testing.T) {
	base := OGVariantSpec{Name: "default", Model: "base-model", Overlay: true, Prompt: "Site prompt."}

	// all nil: every field inherits, name comes from the declaration
	got := ogResolveVariant(base, OGVariant{Name: "minimal"})
	if got != (OGVariantSpec{Name: "minimal", Model: "base-model", Overlay: true, Prompt: "Site prompt."}) {
		t.Errorf("all-nil merge = %+v", got)
	}

	// non-nil overrides, including explicit false and empty string
	empty, off, prompt := "", false, "Photo prompt."
	got = ogResolveVariant(base, OGVariant{Name: "photo", Model: &empty, Overlay: &off, Prompt: &prompt})
	if got != (OGVariantSpec{Name: "photo", Model: "", Overlay: false, Prompt: "Photo prompt."}) {
		t.Errorf("override merge = %+v", got)
	}
}
