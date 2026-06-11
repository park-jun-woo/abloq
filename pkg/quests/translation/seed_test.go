//ff:func feature=quest type=parser control=iteration dimension=1
//ff:what Seed 검증 — 원문 1편이 (선언 언어 − 기본 언어) 매트릭스 Item으로 시드되고 Payload·Key가 채워지는지
package translation

import (
	"testing"

	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestSeed(t *testing.T) {
	root := writeInstance(t)
	origin, _ := passPair()
	path := writeFile(t, root, "content/en/posts/fixture.md", origin)
	items, err := Definition{}.Seed([]string{path})
	if err != nil {
		t.Fatalf("Seed: %v", err)
	}
	if len(items) != 2 {
		t.Fatalf("items = %d, want 2 (ko, ja)", len(items))
	}
	wantKeys := []string{"ko/posts/fixture", "ja/posts/fixture"}
	for i, it := range items {
		if it.Key != wantKeys[i] || it.State != quest.TODO {
			t.Errorf("items[%d] = %s (%s), want %s TODO", i, it.Key, it.State, wantKeys[i])
		}
	}
	var p Payload
	if err := items[0].DecodePayload(&p); err != nil {
		t.Fatalf("DecodePayload: %v", err)
	}
	if p.Root != root || p.Origin != "content/en/posts/fixture.md" ||
		p.Article != "content/ko/posts/fixture.md" || p.OriginLang != "en" ||
		p.Lang != "ko" || p.Section != "posts" || p.Slug != "fixture" {
		t.Errorf("Payload = %+v", p)
	}
}
