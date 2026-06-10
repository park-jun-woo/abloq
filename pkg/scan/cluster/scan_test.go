//ff:func feature=scan type=rule control=iteration dimension=1 topic=cluster
//ff:what Scan 통합 — 위반 4글만 큐 후보(hub 제외), 검출 순서·우선순위·payload JSON(violations/candidates)을 정확히 검증
package cluster

import "testing"

func TestScan(t *testing.T) {
	root := writeRepoFixture(t)
	items := Scan(root, testBlog())
	want := []struct {
		slug     string
		priority int64
	}{
		{"island", 6}, // tag-taxonomy + no-isolated-post, 4 candidates
		{"offtax", 4}, // tag-taxonomy, 3 candidates
		{"orphan", 5}, // no-orphan-tag, 4 candidates (hub never links orphan — "in" stays open)
		{"thin", 4},   // min-internal-links, 3 candidates
	}
	if len(items) != len(want) {
		t.Fatalf("items = %d, want %d (%+v)", len(items), len(want), items)
	}
	for i, w := range want {
		it := items[i]
		if it.Slug != w.slug || it.Priority != w.priority {
			t.Errorf("items[%d] = %s/%d, want %s/%d", i, it.Slug, it.Priority, w.slug, w.priority)
		}
		if it.Kind != "cluster" || it.Lang != "ko" || it.Section != "tech" {
			t.Errorf("items[%d] identity = %+v", i, it)
		}
		if it.Payload["violations"] == "" || it.Payload["candidates"] == "" {
			t.Errorf("items[%d] payload missing keys: %+v", i, it.Payload)
		}
	}
	wantViolations := `[{"rule":"tag-taxonomy","detail":"tags not in geo.taxonomy: rogue"},` +
		`{"rule":"no-isolated-post","detail":"no inbound internal links"}]`
	if items[0].Payload["violations"] != wantViolations {
		t.Errorf("island violations = %s", items[0].Payload["violations"])
	}
	wantCandidates := `[{"section":"tech","slug":"offtax","shared_tags":1,"directions":["in"]},` +
		`{"section":"tech","slug":"thin","shared_tags":1,"directions":["out","in"]},` +
		`{"section":"tech","slug":"hub","shared_tags":1,"directions":["in"]},` +
		`{"section":"tech","slug":"orphan","shared_tags":0,"directions":["out","in"]}]`
	if items[0].Payload["candidates"] != wantCandidates {
		t.Errorf("island candidates = %s", items[0].Payload["candidates"])
	}
	// A blog without declared languages has no default-language graph.
	none := testBlog()
	none.Languages = nil
	if got := Scan(root, none); len(got) != 0 {
		t.Errorf("no languages must scan empty, got %+v", got)
	}
}
