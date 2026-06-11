//ff:func feature=quest type=parser control=sequence
//ff:what 테스트 헬퍼 — 임시 블로그 인스턴스(blog.yaml: en/ko/ja 3언어, body→sources 구조, 언어별 sources 헤딩) 생성
package translation

import "testing"

const instanceBlogYAML = `site:
  baseURL: https://example.com
  title: Translation Fixture
  author: Tester

languages: [en, ko, ja]
sections: [posts]

structure:
  order: [body, sources]
  headings:
    sources: { en: "Sources", ko: "출처", ja: "出典" }
`

func writeInstance(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	writeFile(t, root, "blog.yaml", instanceBlogYAML)
	return root
}
