//ff:func feature=claudemd type=generator control=iteration dimension=1
//ff:what dirsSection이 SSOT/콘텐츠/생성물 수정 금지/큐 디렉토리 규약을 포함하는지 검증
package claudemd

import (
	"strings"
	"testing"
)

func TestDirsSection(t *testing.T) {
	out := dirsSection()
	wants := []string{
		"`blog.yaml` — SSOT",
		"content/{lang}/{section}/{slug}.md",
		"생성물. 직접 수정 금지",
		"quests/queue/",
		"deploy/terraform/",
	}
	for _, w := range wants {
		if !strings.Contains(out, w) {
			t.Errorf("want %q in dirs section, got:\n%s", w, out)
		}
	}
}
