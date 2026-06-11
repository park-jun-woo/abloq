//ff:func feature=quest type=parser control=sequence
//ff:what seedItem 검증 — 번들(디렉토리/insight.yaml) 경로가 index.md 대상과 slug 디렉토리명 Key로 시드되는지
package writing

import "testing"

func TestSeedItem(t *testing.T) {
	root := writeInstance(t)
	_, ins := passFixtures()
	path := writeFile(t, root, "content/en/posts/bundle/insight.yaml", ins)
	it, err := seedItem(path)
	if err != nil {
		t.Fatalf("seedItem: %v", err)
	}
	if it.Key != "en/posts/bundle" {
		t.Errorf("Key = %q, want en/posts/bundle", it.Key)
	}
	var p Payload
	if err := it.DecodePayload(&p); err != nil {
		t.Fatalf("DecodePayload: %v", err)
	}
	if p.Article != "content/en/posts/bundle/index.md" {
		t.Errorf("Article = %q, want content/en/posts/bundle/index.md", p.Article)
	}
}
