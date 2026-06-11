//ff:func feature=blogyaml type=schema control=iteration dimension=1
//ff:what ResolvedVariants/Variant/DefaultVariant의 상속 병합 검증 — nil 상속, 명시 false/"" 오버라이드, 선언 순서, 미선언 이름 false
package blogyaml

import "testing"

func TestOGResolvedVariants(t *testing.T) {
	model, empty, off := "imagen-4", "", false
	prompt := "Photo prompt."
	og := ImageOG{
		Provider: "gemini", Model: "base-model", Overlay: true, Prompt: "Site prompt.",
		Variants: []OGVariant{
			{Name: "minimal"},
			{Name: "photo", Model: &model, Overlay: &off, Prompt: &prompt},
			{Name: "reset", Model: &empty},
		},
	}
	got := og.ResolvedVariants()
	if len(got) != 3 {
		t.Fatalf("resolved = %d, want 3", len(got))
	}
	want := []OGVariantSpec{
		{Name: "minimal", Model: "base-model", Overlay: true, Prompt: "Site prompt."},
		{Name: "photo", Model: "imagen-4", Overlay: false, Prompt: "Photo prompt."},
		{Name: "reset", Model: "", Overlay: true, Prompt: "Site prompt."},
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("variant[%d] = %+v, want %+v", i, got[i], want[i])
		}
	}

	if v, ok := og.Variant("photo"); !ok || v != want[1] {
		t.Errorf("Variant(photo) = %+v %v, want %+v true", v, ok, want[1])
	}
	if _, ok := og.Variant("nope"); ok {
		t.Error("Variant(nope) must report false")
	}

	def := og.DefaultVariant()
	if def != (OGVariantSpec{Name: "default", Model: "base-model", Overlay: true, Prompt: "Site prompt."}) {
		t.Errorf("DefaultVariant = %+v", def)
	}

	if vs := (ImageOG{}).ResolvedVariants(); vs != nil {
		t.Errorf("empty declaration: %v, want nil", vs)
	}
}
