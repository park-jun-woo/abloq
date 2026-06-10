//ff:func feature=scan type=rule control=sequence topic=evidence
//ff:what nextChecks 케이스 — 실패는 이전 +1(분류 무관), 성공은 0 리셋, 이전 상태 없는 키는 1부터
package evidence

import "testing"

func TestNextChecks(t *testing.T) {
	cites := []cite{
		{Lang: "ko", Section: "tech", Slug: "p", URL: "https://dead.example/x"},
		{Lang: "ko", Section: "tech", Slug: "p", URL: "https://flaky.example/y"},
		{Lang: "ko", Section: "tech", Slug: "p", URL: "https://new.example/z"},
	}
	prev := []Check{
		{URL: "https://dead.example/x", Lang: "ko", Section: "tech", Slug: "p", Status: "hard", ConsecutiveFailures: 2},
		{URL: "https://flaky.example/y", Lang: "ko", Section: "tech", Slug: "p", Status: "soft", ConsecutiveFailures: 2},
	}
	statuses := map[string]string{
		"https://dead.example/x":  "soft", // classification changed — the count still rides
		"https://flaky.example/y": "ok",
		"https://new.example/z":   "hard",
	}
	checks := nextChecks(prev, cites, statuses)
	if len(checks) != 3 {
		t.Fatalf("want 3 checks, got %d", len(checks))
	}
	if checks[0].ConsecutiveFailures != 3 || checks[0].Status != "soft" {
		t.Errorf("persisting failure must count regardless of class: %+v", checks[0])
	}
	if checks[1].ConsecutiveFailures != 0 || checks[1].Status != "ok" {
		t.Errorf("recovery must reset to 0: %+v", checks[1])
	}
	if checks[2].ConsecutiveFailures != 1 {
		t.Errorf("first failure starts at 1: %+v", checks[2])
	}
}
