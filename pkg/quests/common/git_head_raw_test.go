//ff:func feature=quest type=parser control=sequence topic=queue
//ff:what gitHeadRaw가 HEAD 스냅샷 바이트를 읽고 HEAD 부재 경로·비git 디렉토리는 에러인지 검증
package common

import (
	"strings"
	"testing"
)

func TestGitHeadRaw(t *testing.T) {
	root, _ := writeFixture(t, "content/en/posts/a.md", fixtureArticleMD)
	gitFixture(t, root)
	raw, err := gitHeadRaw(root, "content/en/posts/a.md")
	if err != nil || !strings.Contains(string(raw), "Body text.") {
		t.Errorf("gitHeadRaw: %v / %q", err, raw)
	}
	if _, err := gitHeadRaw(root, "content/en/posts/missing.md"); err == nil {
		t.Error("path absent from HEAD: want error")
	}
	if _, err := gitHeadRaw(t.TempDir(), "x.md"); err == nil {
		t.Error("non-git dir: want error")
	}
}
