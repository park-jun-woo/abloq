//ff:func feature=gen type=generator control=iteration dimension=1
//ff:what sortPosts가 입력 순서와 무관하게 언어→섹션→최신 날짜 순의 복사본을 내고 원본을 보존하는지 검증
package llms

import "testing"

func TestSortPosts(t *testing.T) {
	posts := []Post{
		{Lang: "en", Section: "opinion", Slug: "e1", Date: "2026-01-01"},
		{Lang: "ko", Section: "tech", Slug: "k3", Date: "2026-02-01"},
		{Lang: "ko", Section: "opinion", Slug: "k1", Date: "2026-01-01"},
		{Lang: "ko", Section: "opinion", Slug: "k2", Date: "2026-03-01"},
	}
	sorted := sortPosts(posts, []string{"ko", "en"}, []string{"opinion", "tech"})
	wantSlugs := []string{"k2", "k1", "k3", "e1"}
	for i, want := range wantSlugs {
		if sorted[i].Slug != want {
			t.Errorf("sorted[%d].Slug = %q, want %q (full: %+v)", i, sorted[i].Slug, want, sorted)
		}
	}
	if posts[0].Slug != "e1" {
		t.Errorf("input slice mutated: %+v", posts)
	}
}
