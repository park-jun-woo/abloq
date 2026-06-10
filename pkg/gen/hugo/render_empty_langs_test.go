//ff:func feature=gen type=generator control=sequence
//ff:what 언어 미선언 Blog에서도 hugo.toml 렌더가 패닉 없이 빈 기본 언어를 내는지 검증
package hugo

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func TestRenderEmptyLangs(t *testing.T) {
	out := string(Render(&blogyaml.Blog{}))
	if !strings.Contains(out, `defaultContentLanguage = ""`) {
		t.Errorf("want empty defaultContentLanguage, got:\n%s", out)
	}
}
