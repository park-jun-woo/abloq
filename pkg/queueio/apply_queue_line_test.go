//ff:func feature=queueio type=parser control=iteration dimension=1
//ff:what applyQueueLine이 keys 항목·payload 라인·최상위 필드를 분류 반영하고 모르는 라인·불량 인용은 에러인지 검증
package queueio

import "testing"

func TestApplyQueueLine(t *testing.T) {
	it := Item{Payload: map[string]string{}}
	for _, ln := range []string{
		"key: \"ko/tech/a\"", "keys:", "  - \"ko/tech/a\"", "kind: \"refresh\"",
		"lang: \"ko\"", "section: \"tech\"", "slug: \"a\"", "priority: 7",
		"payload:", "  lastmod: \"2026-06-01\"", "",
	} {
		if err := applyQueueLine(&it, ln); err != nil {
			t.Fatalf("applyQueueLine(%q): %v", ln, err)
		}
	}
	if it.Kind != "refresh" || it.Lang != "ko" || it.Section != "tech" || it.Slug != "a" || it.Priority != 7 {
		t.Errorf("item = %+v", it)
	}
	if len(it.Keys) != 1 || it.Keys[0] != "ko/tech/a" || it.Payload["lastmod"] != "2026-06-01" {
		t.Errorf("keys/payload = %v / %v", it.Keys, it.Payload)
	}
	if err := applyQueueLine(&it, "payload: {}"); err != nil {
		t.Errorf("empty payload marker: %v", err)
	}
	if err := applyQueueLine(&it, "mystery: 1"); err == nil {
		t.Error("unknown top-level line must error")
	}
	if err := applyQueueLine(&it, "  - unquoted"); err == nil {
		t.Error("unquoted keys entry must error")
	}
	if err := applyQueueLine(&it, "kind: unquoted"); err == nil {
		t.Error("unquoted scalar must error")
	}
}
