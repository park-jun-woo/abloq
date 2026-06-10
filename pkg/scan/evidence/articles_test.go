//ff:func feature=scan type=parser control=iteration dimension=1 topic=evidence
//ff:what articles가 발행 글만 수집하고 draft를 제외하는지 검증 (픽스처 3편 중 2편)
package evidence

import "testing"

func TestArticles(t *testing.T) {
	root := writeRepoFixture(t, "https://cite.example/ref")
	arts := articles(root, testBlog(t))
	if len(arts) != 2 {
		t.Fatalf("want 2 published articles, got %d", len(arts))
	}
	for _, a := range arts {
		if a.Slug == "post-draft" {
			t.Error("draft must be excluded")
		}
	}
}
