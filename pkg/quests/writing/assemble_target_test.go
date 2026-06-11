//ff:func feature=quest type=frame control=sequence
//ff:what assembleTarget 검증 — blog.yaml+글 1편으로 Target 조립(Base nil·Dir=루트), 글 부재·blog.yaml 불량은 에러
package writing

import "testing"

func TestAssembleTarget(t *testing.T) {
	root := writeInstance(t)
	art, _ := passFixtures()
	writeFile(t, root, "content/en/posts/hello.md", art)
	p := Payload{Root: root, Article: "content/en/posts/hello.md",
		Lang: "en", Section: "posts", Slug: "hello"}
	tgt, body, err := assembleTarget(p)
	if err != nil {
		t.Fatalf("assembleTarget: %v", err)
	}
	if tgt.Dir != root || len(tgt.Articles) != 1 || tgt.Articles[0].Base != nil {
		t.Errorf("target = dir %q, %d article(s)", tgt.Dir, len(tgt.Articles))
	}
	if len(body) == 0 {
		t.Error("body empty")
	}
	p.Article = "content/en/posts/absent.md"
	if _, _, err := assembleTarget(p); err == nil {
		t.Error("absent article: want error, got nil")
	}
}
