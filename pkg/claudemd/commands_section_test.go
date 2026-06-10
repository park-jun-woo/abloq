//ff:func feature=claudemd type=generator control=iteration dimension=1
//ff:what commandsSection이 abloq 서브커맨드 전부(validate/generate/check/gate/postbuild/image/claudemd)를 안내하는지 검증
package claudemd

import (
	"strings"
	"testing"
)

func TestCommandsSection(t *testing.T) {
	out := commandsSection()
	wants := []string{
		"abloq validate .",
		"abloq generate .",
		"abloq check .",
		"abloq gate .",
		"abloq gate --offline .",
		"abloq postbuild md .",
		"abloq image convert",
		"abloq image og",
		"abloq claudemd .",
	}
	for _, w := range wants {
		if !strings.Contains(out, w) {
			t.Errorf("want %q in commands section, got:\n%s", w, out)
		}
	}
}
