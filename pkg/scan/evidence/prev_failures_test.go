//ff:func feature=scan type=parser control=sequence topic=evidence
//ff:what prevFailures가 이전 상태를 citation_checks 키로 색인하는지 검증
package evidence

import "testing"

func TestPrevFailures(t *testing.T) {
	m := prevFailures([]Check{
		{URL: "https://a.example/x", Lang: "ko", Section: "tech", Slug: "p", Status: "hard", ConsecutiveFailures: 2},
		{URL: "https://b.example/y", Lang: "ko", Section: "tech", Slug: "p", Status: "ok", ConsecutiveFailures: 0},
	})
	if m[checkKey("https://a.example/x", "ko", "tech", "p")] != 2 {
		t.Errorf("failures index = %v", m)
	}
	if m[checkKey("https://b.example/y", "ko", "tech", "p")] != 0 {
		t.Errorf("ok rows index at 0: %v", m)
	}
}
