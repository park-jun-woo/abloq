//ff:func feature=quest type=parser control=sequence
//ff:what seedOrigin 검증 — 루트 탐색·Key 부품·언어 목록·원문 lastmod 적재, lastmod 부재는 hasLastmod=false
package translation

import (
	"path/filepath"
	"testing"
	"time"
)

func TestSeedOrigin(t *testing.T) {
	root := writeInstance(t)
	origin, _ := passPair()
	path := writeFile(t, root, "content/en/posts/fixture.md", origin)
	src, err := seedOrigin(path)
	if err != nil {
		t.Fatalf("seedOrigin: %v", err)
	}
	if src.root != root || src.origin != "content/en/posts/fixture.md" ||
		src.originLang != "en" || src.section != "posts" || src.slug != "fixture" {
		t.Errorf("src = %+v", src)
	}
	want := time.Date(2026, 6, 3, 0, 0, 0, 0, time.UTC)
	if !src.hasLastmod || !src.lastmod.Equal(want) {
		t.Errorf("lastmod = %v (has=%v), want %v", src.lastmod, src.hasLastmod, want)
	}
	bare := writeFile(t, root, "content/en/posts/bare.md", "no front matter\n")
	src, err = seedOrigin(bare)
	if err != nil {
		t.Fatalf("seedOrigin bare: %v", err)
	}
	if src.hasLastmod {
		t.Error("origin without lastmod: hasLastmod = true, want false")
	}
	if _, err := seedOrigin(filepath.Join(root, "content/en/posts/ghost.md")); err == nil {
		t.Error("origin file absent: want error")
	}
	rootless := filepath.Join(t.TempDir(), "content/en/posts/fixture.md")
	if _, err := seedOrigin(rootless); err == nil {
		t.Error("no blog.yaml ancestor: want error")
	}
	bad := t.TempDir()
	writeFile(t, bad, "blog.yaml", "site:\n  baseURL: not-a-url\n  title: T\n  author: A\n\nlanguages: [en, ko]\nsections: [posts]\n")
	badPath := writeFile(t, bad, "content/en/posts/fixture.md", origin)
	if _, err := seedOrigin(badPath); err == nil {
		t.Error("blog.yaml diagnostics: want error")
	}
	broken := t.TempDir()
	writeFile(t, broken, "blog.yaml", "site: [")
	brokenPath := writeFile(t, broken, "content/en/posts/fixture.md", origin)
	if _, err := seedOrigin(brokenPath); err == nil {
		t.Error("blog.yaml unparseable: want error")
	}
}
