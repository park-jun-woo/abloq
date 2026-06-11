//ff:func feature=claudemd type=generator control=sequence
//ff:what llms_txt mode manual에서 dirsSection 생성물 목록에 static/llms.txt가 빠지는지 검증
package claudemd

import (
	"strings"
	"testing"
)

func TestDirsSectionManual(t *testing.T) {
	out := dirsSection(manualBlog())
	if strings.Contains(out, "llms.txt") {
		t.Errorf("manual mode must drop static/llms.txt from the generated list, got:\n%s", out)
	}
	if !strings.Contains(out, "`hugo.toml` · `static/robots.txt` · `data/jsonld.json`") {
		t.Errorf("want remaining generated files listed, got:\n%s", out)
	}
}
