//ff:func feature=quest type=parser control=sequence
//ff:what Seed 에러 경로 검증 — 인자 없음, 사이드카 규약 밖 경로, 루트(blog.yaml) 부재, 명세 검증 실패
package writing

import "testing"

func TestSeedErrors(t *testing.T) {
	if _, err := (Definition{}).Seed(nil); err == nil {
		t.Error("Seed(nil): want error, got nil")
	}
	root := writeInstance(t)
	bad := writeFile(t, root, "content/en/posts/notes.yaml", "x: 1\n")
	if _, err := (Definition{}).Seed([]string{bad}); err == nil {
		t.Error("non-sidecar path: want error, got nil")
	}
	orphanRoot := t.TempDir()
	_, ins := passFixtures()
	orphan := writeFile(t, orphanRoot, "content/en/posts/a.insight.yaml", ins)
	if _, err := (Definition{}).Seed([]string{orphan}); err == nil {
		t.Error("no blog.yaml root: want error, got nil")
	}
	empty := writeFile(t, root, "content/en/posts/empty.insight.yaml", "topic: t\nclaims: []\n")
	if _, err := (Definition{}).Seed([]string{empty}); err == nil {
		t.Error("invalid spec (claims empty): want error, got nil")
	}
}
