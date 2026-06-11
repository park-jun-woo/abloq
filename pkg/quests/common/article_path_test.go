//ff:func feature=quest type=parser control=sequence topic=queue
//ff:what ArticlePath가 플랫·번들 글을 해석하고 둘 다 없으면 에러인지 검증
package common

import "testing"

func TestArticlePath(t *testing.T) {
	root, _ := writeFixture(t, "content/en/posts/flat.md", fixtureArticleMD)
	got, err := ArticlePath(root, "en", "posts", "flat")
	if err != nil || got != "content/en/posts/flat.md" {
		t.Errorf("flat = %q (%v)", got, err)
	}
	root2, _ := writeFixture(t, "content/en/posts/bundle/index.md", fixtureArticleMD)
	got, err = ArticlePath(root2, "en", "posts", "bundle")
	if err != nil || got != "content/en/posts/bundle/index.md" {
		t.Errorf("bundle = %q (%v)", got, err)
	}
	if _, err := ArticlePath(root, "en", "posts", "nope"); err == nil {
		t.Error("missing article: want error")
	}
}
