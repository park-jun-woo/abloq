//ff:func feature=scan type=parser control=iteration dimension=1 topic=cluster
//ff:what collectPosts가 기본 언어 발행 글만 디렉토리 순서로 수집하는지(draft·번역 제외) 검증
package cluster

import "testing"

func TestCollectPosts(t *testing.T) {
	root := writeRepoFixture(t)
	posts := collectPosts(root, testBlog(), "ko")
	want := []string{"hub", "island", "offtax", "orphan", "thin"} // ReadDir order, draft skipped
	if len(posts) != len(want) {
		t.Fatalf("posts = %d, want %d", len(posts), len(want))
	}
	for i, slug := range want {
		if posts[i].Slug != slug {
			t.Errorf("posts[%d].Slug = %q, want %q", i, posts[i].Slug, slug)
		}
	}
}
