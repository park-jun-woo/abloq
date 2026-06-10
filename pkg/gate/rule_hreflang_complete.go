//ff:func feature=gate type=rule control=iteration dimension=1
//ff:what [hreflang-complete] Hugo 빌드 산출물(public/) head의 hreflang 상호 참조가 전 언어 버전을 덮는지 검증 — 미빌드 시 스킵
package gate

import (
	"os"
	"path/filepath"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// ruleHreflangComplete checks the built pages: every article page must carry
// an hreflang alternate link for each existing language version (self
// included). Skipped when dir/public does not exist (site not built).
func ruleHreflangComplete(t *Target) []blogyaml.Diagnostic {
	if _, err := os.Stat(filepath.Join(t.Dir, "public")); err != nil {
		return nil
	}
	sibs := siblingLangs(t.Articles)
	var diags []blogyaml.Diagnostic
	for _, a := range t.Articles {
		diags = append(diags, hreflangDiags(t.Dir, a, sibs[a.Section+"/"+a.Slug])...)
	}
	return diags
}
