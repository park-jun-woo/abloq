//ff:func feature=quest type=parser control=sequence
//ff:what transPath 검증 — 플랫 원문은 플랫 번역 경로로, 번들 원문은 번들 index.md로 언어 세그먼트 치환
package translation

import "testing"

func TestTransPath(t *testing.T) {
	flat := seedSrc{origin: "content/en/posts/fixture.md", section: "posts", slug: "fixture"}
	if got := transPath(flat, "ja"); got != "content/ja/posts/fixture.md" {
		t.Errorf("flat = %q", got)
	}
	bundle := seedSrc{origin: "content/en/posts/fixture/index.md", section: "posts", slug: "fixture"}
	if got := transPath(bundle, "ar"); got != "content/ar/posts/fixture/index.md" {
		t.Errorf("bundle = %q", got)
	}
}
