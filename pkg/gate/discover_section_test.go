//ff:func feature=gate type=frame control=sequence
//ff:what discoverSection이 섹션 디렉토리의 글을 수집하고 없는 디렉토리에 빈 목록을 반환하는지 검증
package gate

import (
	"path/filepath"
	"testing"
)

func TestDiscoverSection(t *testing.T) {
	hi := buildHeadingIndex(loadGateBlog(t))
	dir := filepath.Join("testdata", "repo-pass")
	arts := discoverSection(dir, hi, "ko", "tech")
	if len(arts) != 1 || arts[0].Slug != "hello" {
		t.Fatalf("want [hello], got %v", arts)
	}
	if arts[0].Doc == nil || !arts[0].Doc.HasFM {
		t.Error("want parsed Doc with front matter")
	}
	if got := discoverSection(dir, hi, "ko", "nope"); len(got) != 0 {
		t.Errorf("missing section dir: want 0, got %d", len(got))
	}
}
