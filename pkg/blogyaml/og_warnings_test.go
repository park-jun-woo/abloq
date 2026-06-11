//ff:func feature=blogyaml type=rule control=sequence
//ff:what OGWarnings가 local+variants 조합에만 경고를 내고 gemini·무variants·nil Blog에는 침묵하는지 검증
package blogyaml

import "testing"

func TestOGWarnings(t *testing.T) {
	localVariants := &Blog{Image: Image{OG: ImageOG{Variants: []OGVariant{{Name: "minimal"}}}}}
	warns := OGWarnings("blog.yaml", localVariants, lineIndex{"image.og.variants": 7})
	if len(warns) != 1 || warns[0].Rule != "og-local-variants" || warns[0].Line != 7 {
		t.Fatalf("local+variants: %v", warns)
	}
	gemini := &Blog{Image: Image{OG: ImageOG{Provider: "gemini", Variants: []OGVariant{{Name: "minimal"}}}}}
	if w := OGWarnings("blog.yaml", gemini, lineIndex{}); len(w) != 0 {
		t.Errorf("gemini+variants: %v, want none", w)
	}
	if w := OGWarnings("blog.yaml", &Blog{}, lineIndex{}); len(w) != 0 {
		t.Errorf("no variants: %v, want none", w)
	}
	if w := OGWarnings("blog.yaml", nil, lineIndex{}); len(w) != 0 {
		t.Errorf("nil blog: %v, want none", w)
	}
}
