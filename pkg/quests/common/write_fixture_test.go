//ff:func feature=quest type=parser control=sequence
//ff:what 테스트 헬퍼 — 임시 블로그 인스턴스(blog.yaml: en/posts, body→sources, min_sources 기본 1)와 루트 기준 파일 기록
package common

import (
	"os"
	"path/filepath"
	"testing"
)

const fixtureBlogYAML = `site:
  baseURL: https://example.com
  title: Common Fixture
  author: Tester

languages: [en]
sections: [posts]

structure:
  order: [body, sources]
  headings:
    sources: { en: "Sources" }
`

func writeFixture(t *testing.T, rel, content string) (root, abs string) {
	t.Helper()
	root = t.TempDir()
	write := func(r, c string) string {
		p := filepath.Join(root, filepath.FromSlash(r))
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			t.Fatalf("mkdir: %v", err)
		}
		if err := os.WriteFile(p, []byte(c), 0o644); err != nil {
			t.Fatalf("write %s: %v", r, err)
		}
		return p
	}
	write("blog.yaml", fixtureBlogYAML)
	if rel != "" {
		abs = write(rel, content)
	}
	return root, abs
}
