//ff:func feature=image type=generator control=sequence
//ff:what 잡 1건의 출력 경로 결정 — 다중 안은 DraftDir/{slug}/{variant}-{n}.webp, 단일 직행은 OutDir/{slug}.webp
package img

import (
	"fmt"
	"path/filepath"
)

// ogJobOut resolves where one generation lands: the reviewed draft directory
// for multi-candidate runs, the adopted path for the single direct call.
func ogJobOut(spec OGAISpec, variant string, n int) string {
	if spec.Multi {
		return filepath.Join(spec.DraftDir, spec.Slug, fmt.Sprintf("%s-%d.webp", variant, n))
	}
	return filepath.Join(spec.OutDir, spec.Slug+".webp")
}
