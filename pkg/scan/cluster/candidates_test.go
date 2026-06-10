//ff:func feature=scan type=generator control=iteration dimension=1 topic=cluster
//ff:what candidates가 자기 참조·완전 연결을 제외하고 정렬 상위 5건만 남기는지 검증
package cluster

import "testing"

func TestCandidates(t *testing.T) {
	self := post{Section: "tech", Slug: "self", Date: "2026-01-10", Tags: []string{"geo"}}
	posts := []post{
		self,
		{Section: "tech", Slug: "full", Date: "2026-01-10", Tags: []string{"geo"}},   // fully connected → dropped
		{Section: "tech", Slug: "close", Date: "2026-01-09", Tags: []string{"geo"}},  // shared 1, dist 1
		{Section: "tech", Slug: "far", Date: "2026-01-01", Tags: []string{"geo"}},    // shared 1, dist 9
		{Section: "tech", Slug: "alien", Date: "2026-01-10", Tags: []string{"etc"}},  // shared 0
		{Section: "tech", Slug: "twin-a", Date: "2026-01-08", Tags: []string{"geo"}}, // shared 1, dist 2
		{Section: "tech", Slug: "twin-b", Date: "2026-01-08", Tags: []string{"geo"}}, // tie with twin-a → key order
		{Section: "tech", Slug: "spare", Date: "2026-01-07", Tags: []string{"geo"}},  // shared 1, dist 3
	}
	edges := map[string]bool{"tech/self->tech/full": true, "tech/full->tech/self": true}
	got := candidates(self, posts, edges)
	want := []string{"close", "twin-a", "twin-b", "spare", "far"} // alien (shared 0) falls off the cap
	if len(got) != maxCandidates {
		t.Fatalf("candidates = %d, want %d (%+v)", len(got), maxCandidates, got)
	}
	for i, slug := range want {
		if got[i].Slug != slug {
			t.Errorf("candidates[%d] = %q, want %q", i, got[i].Slug, slug)
		}
	}
}
