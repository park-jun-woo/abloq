//ff:func feature=quest type=parser control=sequence
//ff:what Seed 에러 경로 검증 — 인자 없음, content 밖 경로, 번역(비기본 언어) 경로 시드, 단일 언어 블로그
package translation

import "testing"

func TestSeedErrors(t *testing.T) {
	if _, err := (Definition{}).Seed(nil); err == nil {
		t.Error("no args: want error")
	}
	root := writeInstance(t)
	origin, ko := passPair()
	stray := writeFile(t, root, "static/en/posts/fixture.md", origin)
	if _, err := (Definition{}).Seed([]string{stray}); err == nil {
		t.Error("non-content path: want error")
	}
	koPath := writeFile(t, root, "content/ko/posts/fixture.md", ko)
	if _, err := (Definition{}).Seed([]string{koPath}); err == nil {
		t.Error("non-default-language origin: want error")
	}
	mono := t.TempDir()
	writeFile(t, mono, "blog.yaml", "site:\n  baseURL: https://example.com\n  title: T\n  author: A\n\nlanguages: [en]\nsections: [posts]\n")
	monoPath := writeFile(t, mono, "content/en/posts/fixture.md", origin)
	if _, err := (Definition{}).Seed([]string{monoPath}); err == nil {
		t.Error("single-language blog: want error")
	}
}
