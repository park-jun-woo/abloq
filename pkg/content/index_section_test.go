//ff:func feature=content type=parser control=sequence
//ff:what indexSection이 content/{lang}/{section}/ 발행 글만 수집하고 디렉토리가 없으면 빈 목록을 돌려주는지 검증
package content

import (
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func TestIndexSection(t *testing.T) {
	root := filepath.Join("testdata", "fixture")
	b, diags, err := blogyaml.Load(filepath.Join(root, "blog.yaml"))
	if err != nil || len(diags) > 0 {
		t.Fatalf("blog.yaml load: err=%v diags=%v", err, diags)
	}
	entries := indexSection(root, b, "ko", "tech")
	if len(entries) != 2 {
		t.Fatalf("ko/tech entries = %d, want 2 (draft·_index·비마크다운·index.md 없는 번들 제외): %+v", len(entries), entries)
	}
	if entries[0].Slug != "post-a" || entries[1].Slug != "post-b" {
		t.Errorf("ko/tech order = [%s %s], want directory-name order", entries[0].Slug, entries[1].Slug)
	}
	if missing := indexSection(root, b, "ko", "opinion"); missing != nil {
		t.Errorf("missing section dir = %+v, want nil", missing)
	}
}
