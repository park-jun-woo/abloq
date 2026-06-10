//ff:func feature=scan type=rule control=iteration dimension=1 topic=cluster
//ff:what violations가 4룰을 고정 순서로 수집하고 무위반 글에 빈 목록을 주는지 검증
package cluster

import "testing"

func TestViolations(t *testing.T) {
	b := testBlog()
	tags := map[string]int64{"geo": 2, "abloq": 2, "solo": 1, "rogue": 2}
	inlinks := map[string]int64{"tech/ok": 1}
	ok := post{Section: "tech", Slug: "ok", Tags: []string{"geo", "abloq"}, Outlinks: []string{"tech/a", "tech/b"}}
	if got := violations(ok, b, tags, inlinks); len(got) != 0 {
		t.Errorf("clean article flagged: %+v", got)
	}
	bad := post{Section: "tech", Slug: "bad", Tags: []string{"solo", "rogue"}, Outlinks: []string{}}
	got := violations(bad, b, tags, inlinks)
	wantRules := []string{"tag-taxonomy", "no-orphan-tag", "min-internal-links", "no-isolated-post"}
	if len(got) != len(wantRules) {
		t.Fatalf("violations = %+v, want %d rules", got, len(wantRules))
	}
	for i, rule := range wantRules {
		if got[i].Rule != rule {
			t.Errorf("violations[%d].Rule = %q, want %q", i, got[i].Rule, rule)
		}
	}
}
