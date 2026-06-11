//ff:func feature=scan type=generator control=sequence topic=cluster
//ff:what clusterItem이 kind=cluster 항목을 결정적 payload(JSON 문자열 값)·전 언어 keys·위반+후보 합 priority로 조립하는지 검증
package cluster

import "testing"

func TestClusterItem(t *testing.T) {
	p := post{Section: "tech", Slug: "thin", Date: "2026-01-04"}
	viols := []Violation{{Rule: "min-internal-links", Detail: "outbound internal links 1 below min 2"}}
	cands := []Candidate{{Section: "tech", Slug: "hub", SharedTags: 2, Directions: []string{"out"}}}
	it := clusterItem(p, "ko", viols, cands, []string{"ko", "en"})
	if it.Kind != "cluster" || it.Slug != "thin" || it.Lang != "ko" || it.Section != "tech" {
		t.Errorf("item identity = %+v", it)
	}
	if it.Priority != 2 {
		t.Errorf("priority = %d, want 2 (1 violation + 1 candidate)", it.Priority)
	}
	if it.Payload["violations"] != `[{"rule":"min-internal-links","detail":"outbound internal links 1 below min 2"}]` {
		t.Errorf("violations payload = %s", it.Payload["violations"])
	}
	if it.Payload["candidates"] != `[{"section":"tech","slug":"hub","shared_tags":2,"directions":["out"]}]` {
		t.Errorf("candidates payload = %s", it.Payload["candidates"])
	}
	if len(it.Keys) != 2 || it.Keys[0] != "ko/tech/thin" || it.Keys[1] != "en/tech/thin" {
		t.Errorf("keys must cover every declared language: %v", it.Keys)
	}
	empty := clusterItem(p, "ko", viols, nil, []string{"ko"})
	if _, ok := empty.Payload["candidates"]; ok {
		t.Error("empty candidate list must omit the key")
	}
}
