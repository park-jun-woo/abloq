//ff:func feature=gen type=parser control=iteration dimension=1
//ff:what postFromEntry가 글.md·번들/index.md를 Post로 만들고 _index·draft·비마크다운·빈 번들을 거르는지 검증
package llms

import (
	"path/filepath"
	"testing"
)

func TestPostFromEntry(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "plain.md"), "---\ntitle: Plain\ndate: 2026-01-01\n---\nbody\n")
	writeFile(t, filepath.Join(dir, "bundle", "index.md"), "---\ntitle: Bundle\n---\nbody\n")
	writeFile(t, filepath.Join(dir, "untitled.md"), "---\ndate: 2026-01-02\n---\nbody\n")
	writeFile(t, filepath.Join(dir, "_index.md"), "---\ntitle: Section\n---\n")
	writeFile(t, filepath.Join(dir, "draft.md"), "---\ntitle: Draft\ndraft: true\n---\n")
	writeFile(t, filepath.Join(dir, "image.png"), "not markdown")
	writeFile(t, filepath.Join(dir, "empty-bundle", "notes.txt"), "no index.md")
	cases := []struct {
		entryName string
		wantOK    bool
		wantSlug  string
		wantTitle string
	}{
		{"plain.md", true, "plain", "Plain"},
		{"bundle", true, "bundle", "Bundle"},
		{"untitled.md", true, "untitled", "untitled"},
		{"_index.md", false, "", ""},
		{"draft.md", false, "", ""},
		{"image.png", false, "", ""},
		{"empty-bundle", false, "", ""},
	}
	for _, tc := range cases {
		t.Run(tc.entryName, func(t *testing.T) { checkPostFromEntry(t, dir, tc.entryName, tc.wantOK, tc.wantSlug, tc.wantTitle) })
	}
}
