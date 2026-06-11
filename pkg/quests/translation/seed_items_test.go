//ff:func feature=quest type=parser control=sequence
//ff:what seedItems 검증 — 최신 번역(lastmod ≥ 원문) 언어는 미생성, 구판·미존재 언어만 아이템화
package translation

import "testing"

func TestSeedItems(t *testing.T) {
	root := writeInstance(t)
	origin, ko := passPair()
	originPath := writeFile(t, root, "content/en/posts/fixture.md", origin)
	writeFile(t, root, "content/ko/posts/fixture.md", ko) // lastmod == origin → fresh
	src, err := seedOrigin(originPath)
	if err != nil {
		t.Fatalf("seedOrigin: %v", err)
	}
	items, err := seedItems(src)
	if err != nil {
		t.Fatalf("seedItems: %v", err)
	}
	if len(items) != 1 || items[0].Key != "ja/posts/fixture" {
		t.Fatalf("items = %v, want only ja/posts/fixture", items)
	}
}
