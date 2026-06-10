//ff:func feature=queueio type=generator control=sequence
//ff:what Serialize가 key 필드(게이트 계약 원문)·고정 필드 순서·payload 키 정렬·빈 payload {}를 바이트 결정적으로 내는지 검증
package queueio

import (
	"strings"
	"testing"
)

func TestSerialize(t *testing.T) {
	it := Item{
		Kind: "refresh", Slug: "post-a", Lang: "ko", Section: "tech",
		Priority: 20605,
		Payload:  map[string]string{"lastmod": "2026-06-05", "freshness_days": "1"},
	}
	got := string(Serialize(it))
	want := "key: \"ko/tech/post-a\"\n" +
		"kind: \"refresh\"\n" +
		"lang: \"ko\"\n" +
		"section: \"tech\"\n" +
		"slug: \"post-a\"\n" +
		"priority: 20605\n" +
		"payload:\n" +
		"  freshness_days: \"1\"\n" +
		"  lastmod: \"2026-06-05\"\n"
	if got != want {
		t.Errorf("Serialize mismatch:\n got: %q\nwant: %q", got, want)
	}
	if !strings.Contains(got, "ko/tech/post-a") {
		t.Error("gate join key must appear verbatim")
	}
	empty := string(Serialize(Item{Kind: "refresh", Slug: "s", Lang: "ko", Section: "tech"}))
	if !strings.HasSuffix(empty, "payload: {}\n") {
		t.Errorf("empty payload must render {}: %q", empty)
	}
}
