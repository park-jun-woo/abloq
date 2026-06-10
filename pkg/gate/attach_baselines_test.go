//ff:func feature=gate type=frame control=sequence topic=baseline
//ff:what attachBaselines가 미변경·미추적=현재본 공유, 변경=HEAD 파싱본, 비 git=nil로 Base를 채우는지 검증
package gate

import (
	"path/filepath"
	"testing"
)

func TestAttachBaselines(t *testing.T) {
	b := loadGateBlog(t)
	hi := buildHeadingIndex(b)
	dir := t.TempDir()
	rel := filepath.Join("content", "en", "tech", "a.md")
	keptRel := filepath.Join("content", "en", "tech", "c.md")
	writeRepoFile(t, dir, rel, "---\ntitle: A\nlastmod: 2026-01-01\n---\nold body\n")
	writeRepoFile(t, dir, keptRel, "---\ntitle: C\n---\nkept body\n")
	initGitRepo(t, dir)
	writeRepoFile(t, dir, rel, "---\ntitle: A\nlastmod: 2026-02-01\n---\nnew body\n")
	newRel := filepath.Join("content", "en", "tech", "b.md")
	writeRepoFile(t, dir, newRel, "---\ntitle: B\n---\nbrand new\n")

	changed := &Article{Lang: "en", Path: rel, Doc: parseDoc(hi, "en", "x")}
	fresh := &Article{Lang: "en", Path: newRel, Doc: parseDoc(hi, "en", "y")}
	kept := &Article{Lang: "en", Path: keptRel, Doc: parseDoc(hi, "en", "k")}
	attachBaselines(dir, hi, []*Article{changed, fresh, kept})
	if kept.Base != kept.Doc {
		t.Error("unchanged file: want Base shared with Doc")
	}
	if changed.Base == nil || changed.Base == changed.Doc {
		t.Fatal("changed file: want distinct HEAD baseline")
	}
	if got := fmLineValue(changed.Base.FrontMatter, "lastmod"); got != "2026-01-01" {
		t.Errorf("baseline lastmod = %q, want HEAD value", got)
	}
	if fresh.Base != fresh.Doc {
		t.Error("untracked file: want Base shared with Doc (no baseline obligations)")
	}
	unchanged := &Article{Lang: "en", Path: rel, Doc: parseDoc(hi, "en", "z")}
	attachBaselines(t.TempDir(), hi, []*Article{unchanged})
	if unchanged.Base != nil {
		t.Error("non-git dir: want nil baseline")
	}
}
