//ff:func feature=claudemd type=generator control=sequence
//ff:what llms_txt mode manual에서 forbiddenSection 손수정 금지 목록에 llms.txt가 빠지는지 검증
package claudemd

import (
	"strings"
	"testing"
)

func TestForbiddenSectionManual(t *testing.T) {
	out := forbiddenSection(manualBlog())
	if strings.Contains(out, "llms.txt") {
		t.Errorf("manual mode must drop llms.txt from the no-hand-edit list, got:\n%s", out)
	}
	if !strings.Contains(out, "생성물(hugo.toml/robots.txt/jsonld.json)") {
		t.Errorf("want remaining generated files listed, got:\n%s", out)
	}
}
