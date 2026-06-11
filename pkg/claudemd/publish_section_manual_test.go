//ff:func feature=claudemd type=generator control=sequence
//ff:what llms_txt mode manual에서 publishSection 재생성 단계가 llms.txt를 언급하지 않는지 검증
package claudemd

import (
	"strings"
	"testing"
)

func TestPublishSectionManual(t *testing.T) {
	out := publishSection(manualBlog())
	if strings.Contains(out, "llms.txt") {
		t.Errorf("manual mode must not mention llms.txt in the regenerate step, got:\n%s", out)
	}
	if !strings.Contains(out, "파생물 재생성 (필수)") {
		t.Errorf("want regenerate step kept, got:\n%s", out)
	}
}
