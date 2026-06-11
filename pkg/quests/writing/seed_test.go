//ff:func feature=quest type=parser control=sequence
//ff:what Seed 검증 — 사이드카 insight.yaml 1건이 Key=lang/section/slug Item 1개로 시드되고 Payload가 채워지는지
package writing

import (
	"testing"

	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestSeed(t *testing.T) {
	root := writeInstance(t)
	_, ins := passFixtures()
	path := writeFile(t, root, "content/en/posts/hello.insight.yaml", ins)
	items, err := Definition{}.Seed([]string{path})
	if err != nil {
		t.Fatalf("Seed: %v", err)
	}
	if len(items) != 1 {
		t.Fatalf("items = %d, want 1", len(items))
	}
	it := items[0]
	if it.Key != "en/posts/hello" {
		t.Errorf("Key = %q, want en/posts/hello", it.Key)
	}
	if it.State != quest.TODO {
		t.Errorf("State = %s, want TODO", it.State)
	}
	var p Payload
	if err := it.DecodePayload(&p); err != nil {
		t.Fatalf("DecodePayload: %v", err)
	}
	if p.Root != root {
		t.Errorf("Root = %q, want %q", p.Root, root)
	}
	if p.Article != "content/en/posts/hello.md" {
		t.Errorf("Article = %q", p.Article)
	}
	if p.Insight != "content/en/posts/hello.insight.yaml" {
		t.Errorf("Insight = %q", p.Insight)
	}
	if p.Lang != "en" || p.Section != "posts" || p.Slug != "hello" {
		t.Errorf("key parts = %s/%s/%s", p.Lang, p.Section, p.Slug)
	}
}
