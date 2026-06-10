//ff:func feature=scan type=rule control=sequence topic=evidence
//ff:what rotURLs가 해당 글의 연속 실패 3회 이상 URL만 — 2회는 미확정, 다른 글 좌표는 제외 — 거르는지 검증
package evidence

import "testing"

func TestRotURLs(t *testing.T) {
	checks := []Check{
		{URL: "https://gone.example/x", Lang: "ko", Section: "tech", Slug: "p", Status: "hard", ConsecutiveFailures: 3},
		{URL: "https://flaky.example/y", Lang: "ko", Section: "tech", Slug: "p", Status: "soft", ConsecutiveFailures: 2},
		{URL: "https://other.example/z", Lang: "ko", Section: "tech", Slug: "other", Status: "hard", ConsecutiveFailures: 5},
	}
	urls := rotURLs(checks, "ko", "tech", "p")
	if len(urls) != 1 || urls[0] != "https://gone.example/x" {
		t.Errorf("rotURLs = %v, want only the 3-strike URL of this article", urls)
	}
	if got := rotURLs(checks, "ko", "tech", "none"); got != nil {
		t.Errorf("article without rot must be nil: %v", got)
	}
}
