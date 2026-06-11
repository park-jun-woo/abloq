//ff:func feature=image type=generator control=sequence
//ff:what ogJobOut 검증 — Multi는 DraftDir/{slug}/{variant}-{n}.webp, 단일 직행은 OutDir/{slug}.webp
package img

import (
	"path/filepath"
	"testing"
)

func TestOGJobOut(t *testing.T) {
	spec := OGAISpec{Slug: "post", OutDir: "static/images", DraftDir: "og"}

	// single direct call: adopted path, variant/n ignored
	if got, want := ogJobOut(spec, "minimal", 2), filepath.Join("static/images", "post.webp"); got != want {
		t.Errorf("direct = %q, want %q", got, want)
	}

	// multi: draft path keyed by variant and sample number
	spec.Multi = true
	if got, want := ogJobOut(spec, "minimal", 2), filepath.Join("og", "post", "minimal-2.webp"); got != want {
		t.Errorf("draft = %q, want %q", got, want)
	}
}
