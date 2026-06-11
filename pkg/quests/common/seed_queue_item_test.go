//ff:func feature=quest type=parser control=sequence topic=queue
//ff:what seedQueueItem이 Key·payload(keys 보유/legacy 보충)를 채우고 대상 글 부재는 에러인지 검증
package common

import (
	"reflect"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/queueio"
)

func TestSeedQueueItem(t *testing.T) {
	root, _ := writeFixture(t, "content/en/posts/a.md", fixtureArticleMD)
	qit := queueio.Item{Kind: "refresh", Slug: "a", Lang: "en", Section: "posts",
		Keys: []string{"en/posts/a"}, Payload: map[string]string{"lastmod": "2026-06-01"}}
	it, err := seedQueueItem(root, qit, []string{"en"})
	if err != nil || it.Key != "en/posts/a" {
		t.Fatalf("item = %+v (%v)", it, err)
	}
	var p QueuePayload
	if err := it.DecodePayload(&p); err != nil {
		t.Fatal(err)
	}
	if p.Article != "content/en/posts/a.md" || !reflect.DeepEqual(p.Keys, []string{"en/posts/a"}) {
		t.Errorf("payload = %+v", p)
	}
	// A legacy queue file without keys: falls back to the declared languages.
	qit.Keys = nil
	it, err = seedQueueItem(root, qit, []string{"en", "ko"})
	if err != nil {
		t.Fatal(err)
	}
	if err := it.DecodePayload(&p); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(p.Keys, []string{"en/posts/a", "ko/posts/a"}) {
		t.Errorf("legacy fallback keys = %v", p.Keys)
	}
	qit.Slug = "ghost"
	if _, err := seedQueueItem(root, qit, []string{"en"}); err == nil {
		t.Error("missing article: want error")
	}
}
