//ff:func feature=visibility type=parser control=iteration dimension=1 topic=crawl
//ff:what 임시 블로그 저장소 픽스처 — ko 루트 서빙(기본 언어)+en 서브디렉토리, tech 글 3편(ko post-a·post-b, en post-a), URL 역매핑 골든 테스트 공용
package cflog

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

const testBlogYAML = "site:\n  baseURL: https://blog.test\n  title: T\n  author: A\n" +
	"  default_lang_in_subdir: false\n" +
	"languages: [ko, en]\nsections: [tech]\n" +
	"structure:\n  order: [body, sources]\n  headings:\n    sources: { ko: \"출처\", en: \"Sources\" }\n"

// writeRepoFixture builds the blog repository whose served URLs the
// committed testdata/logs fixtures hit: /tech/post-a/, /tech/post-b/
// (root-served ko) and /en/tech/post-a/, each with its .md twin.
func writeRepoFixture(t *testing.T) (string, *blogyaml.Blog) {
	t.Helper()
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "blog.yaml"), []byte(testBlogYAML), 0o644); err != nil {
		t.Fatalf("write blog.yaml: %v", err)
	}
	posts := map[string]string{
		"content/ko/tech/post-a.md": "---\ntitle: A\ndate: 2026-05-01\n---\n\n본문 A.\n",
		"content/ko/tech/post-b.md": "---\ntitle: B\ndate: 2026-05-02\n---\n\n본문 B.\n",
		"content/en/tech/post-a.md": "---\ntitle: A EN\ndate: 2026-05-01\n---\n\nBody A.\n",
	}
	for name, body := range posts {
		p := filepath.Join(dir, filepath.FromSlash(name))
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			t.Fatalf("MkdirAll: %v", err)
		}
		if err := os.WriteFile(p, []byte(body), 0o644); err != nil {
			t.Fatalf("write %s: %v", name, err)
		}
	}
	b, diags, err := blogyaml.Load(filepath.Join(dir, "blog.yaml"))
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if len(diags) != 0 {
		t.Fatalf("fixture blog.yaml has diagnostics: %v", diags)
	}
	return dir, b
}
