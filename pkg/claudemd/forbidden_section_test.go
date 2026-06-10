//ff:func feature=claudemd type=generator control=iteration dimension=1
//ff:what forbiddenSection이 치즈 방어 금지사항(생성물 수정/lastmod 위조/출처 날조/임계값 완화)을 포함하는지 검증
package claudemd

import (
	"strings"
	"testing"
)

func TestForbiddenSection(t *testing.T) {
	out := forbiddenSection()
	wants := []string{
		"치즈 방어",
		"손으로 고치지 않는다",
		"honest-lastmod",
		"citation-exists",
		"quests/queue/",
		"임계값을 완화하지 않는다",
	}
	for _, w := range wants {
		if !strings.Contains(out, w) {
			t.Errorf("want %q in forbidden section, got:\n%s", w, out)
		}
	}
}
