//ff:func feature=gen type=parser control=sequence
//ff:what collectSection이 발행 글만 수집하고 없는 섹션 디렉토리에는 빈 목록을 내는지 검증
package llms

import (
	"path/filepath"
	"testing"
)

func TestCollectSection(t *testing.T) {
	root := t.TempDir()
	section := filepath.Join(root, "content", "ko", "opinion")
	writeFile(t, filepath.Join(section, "a.md"), "---\ntitle: A\n---\nbody\n")
	writeFile(t, filepath.Join(section, "d.md"), "---\ntitle: D\ndraft: true\n---\n")
	posts := collectSection(root, "ko", "opinion")
	if len(posts) != 1 || posts[0].Slug != "a" {
		t.Errorf("want 1 published post (slug a), got %+v", posts)
	}
	if missing := collectSection(root, "ko", "missing"); missing != nil {
		t.Errorf("want nil for missing section dir, got %+v", missing)
	}
}
