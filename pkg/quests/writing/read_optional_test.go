//ff:func feature=quest type=parser control=sequence
//ff:what readOptional 검증 — 존재 파일은 본문, 빈 경로·부재 파일은 빈 문자열
package writing

import "testing"

func TestReadOptional(t *testing.T) {
	root := t.TempDir()
	writeFile(t, root, "quests/writing/logs/a.md", "log\n")
	if got := readOptional(root, "quests/writing/logs/a.md"); got != "log\n" {
		t.Errorf("got %q", got)
	}
	if got := readOptional(root, ""); got != "" {
		t.Errorf("empty path: got %q", got)
	}
	if got := readOptional(root, "nope.md"); got != "" {
		t.Errorf("absent file: got %q", got)
	}
}
