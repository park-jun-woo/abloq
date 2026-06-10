//ff:func feature=gen type=parser control=sequence
//ff:what collectLang이 blog.yaml 선언 섹션 순서대로 한 언어의 발행 글을 모으는지 검증
package llms

import (
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func TestCollectLang(t *testing.T) {
	root := t.TempDir()
	writeFile(t, filepath.Join(root, "content", "ko", "tech", "t.md"), "---\ntitle: T\n---\n")
	writeFile(t, filepath.Join(root, "content", "ko", "opinion", "o.md"), "---\ntitle: O\n---\n")
	b := &blogyaml.Blog{Sections: []string{"opinion", "tech"}}
	posts := collectLang(root, b, "ko")
	if len(posts) != 2 || posts[0].Section != "opinion" || posts[1].Section != "tech" {
		t.Errorf("want [opinion tech] posts in section order, got %+v", posts)
	}
}
