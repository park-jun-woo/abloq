//ff:func feature=quest type=parser control=sequence
//ff:what 테스트 헬퍼 — 임시 블로그 인스턴스(blog.yaml: en/posts, 구조 5키, min_sources 기본 1) 생성
package writing

import "testing"

const instanceBlogYAML = `site:
  baseURL: https://example.com
  title: Quest Fixture
  author: Tester

languages: [en]
sections: [posts]

structure:
  order: [image, attribution, body, sources, changelog]
  headings:
    sources: { en: "Sources" }
    changelog: { en: "Changelog" }
`

func writeInstance(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	writeFile(t, root, "blog.yaml", instanceBlogYAML)
	return root
}
