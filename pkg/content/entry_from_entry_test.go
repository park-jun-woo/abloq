//ff:func feature=content type=parser control=iteration dimension=1
//ff:what entryFromEntry가 번들/플랫 글을 인덱스 항목으로 변환하고 _접두·비마크다운·draft·front matter 없음·index.md 없는 번들을 거르는지 검증
package content

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func TestEntryFromEntry(t *testing.T) {
	root := filepath.Join("testdata", "fixture")
	b, diags, err := blogyaml.Load(filepath.Join(root, "blog.yaml"))
	if err != nil || len(diags) > 0 {
		t.Fatalf("blog.yaml load: err=%v diags=%v", err, diags)
	}
	sectionDir := filepath.Join(root, "content", "ko", "tech")
	dirEntries, err := os.ReadDir(sectionDir)
	if err != nil {
		t.Fatalf("ReadDir: %v", err)
	}
	want := map[string]bool{
		"_index.md":  false, // 언더스코어 접두 — 스킵
		"notes.txt":  false, // 비마크다운 — 스킵
		"draft-x.md": false, // draft: true — 스킵
		"no-fm.md":   false, // front matter 없음 — 스킵
		"empty-dir":  false, // index.md 없는 번들 — 스킵
		"post-a":     true,  // 번들(post-a/index.md)
		"post-b.md":  true,  // 플랫 파일
	}
	for _, de := range dirEntries {
		e, ok := entryFromEntry(sectionDir, b, "ko", "tech", de)
		expect, known := want[de.Name()]
		if !known {
			t.Errorf("unexpected fixture entry %q", de.Name())
			continue
		}
		if ok != expect {
			t.Errorf("entryFromEntry(%q) ok = %v, want %v", de.Name(), ok, expect)
			continue
		}
		if de.Name() == "post-b.md" && (e.Slug != "post-b" || e.Title != "post-b" || e.Lastmod != e.Date) {
			t.Errorf("post-b fallbacks (title→slug, lastmod→date) = %+v", e)
		}
	}
	enDir := filepath.Join(root, "content", "en", "tech")
	enEntries, err := os.ReadDir(enDir)
	if err != nil || len(enEntries) != 1 {
		t.Fatalf("en/tech fixture: err=%v entries=%v", err, enEntries)
	}
	if e, ok := entryFromEntry(enDir, b, "en", "tech", enEntries[0]); !ok || e.Slug != "custom-en" {
		t.Errorf("front matter slug must override the file stem custom.md: ok=%v entry=%+v", ok, e)
	}
}
