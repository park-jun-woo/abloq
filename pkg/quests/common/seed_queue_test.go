//ff:func feature=quest type=parser control=sequence topic=queue
//ff:what SeedQueue가 kind 일치 큐 파일만 priority 내림차순으로 시드하고 payload(keys·queue)를 고정하며, 큐 디렉토리 부재·대상 글 부재는 에러인지 검증
package common

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/queueio"
)

func TestSeedQueue(t *testing.T) {
	root, _ := writeFixture(t, "content/en/posts/low.md", fixtureArticleMD)
	if err := os.WriteFile(filepath.Join(root, "content", "en", "posts", "high.md"), []byte(fixtureArticleMD), 0o644); err != nil {
		t.Fatal(err)
	}
	queueDir := filepath.Join(root, "quests", "queue")
	items := []queueio.Item{
		{Kind: "refresh", Slug: "low", Lang: "en", Section: "posts", Priority: 1,
			Keys: []string{"en/posts/low"}, Payload: map[string]string{"lastmod": "2026-01-01"}},
		{Kind: "refresh", Slug: "high", Lang: "en", Section: "posts", Priority: 9,
			Payload: map[string]string{"lastmod": "2026-02-01"}},
		{Kind: "evidence", Slug: "low", Lang: "en", Section: "posts", Priority: 99,
			Payload: map[string]string{}},
	}
	if err := queueio.WriteDir(queueDir, items); err != nil {
		t.Fatal(err)
	}
	got, err := SeedQueue("refresh", []string{root})
	if err != nil {
		t.Fatalf("SeedQueue: %v", err)
	}
	if len(got) != 2 || got[0].Key != "en/posts/high" || got[1].Key != "en/posts/low" {
		t.Fatalf("want [high low] by priority desc, got %+v", got)
	}
	var p QueuePayload
	if err := got[0].DecodePayload(&p); err != nil {
		t.Fatal(err)
	}
	if p.Root != root || p.Article != "content/en/posts/high.md" || p.Queue["lastmod"] != "2026-02-01" {
		t.Errorf("payload = %+v", p)
	}
	if len(p.Keys) != 1 || p.Keys[0] != "en/posts/high" {
		t.Errorf("missing keys: must fall back to declared languages: %v", p.Keys)
	}
	// A queue item whose article does not exist is a seed error.
	if err := queueio.WriteDir(queueDir, []queueio.Item{{Kind: "refresh", Slug: "ghost", Lang: "en", Section: "posts"}}); err != nil {
		t.Fatal(err)
	}
	if _, err := SeedQueue("refresh", []string{root}); err == nil {
		t.Error("queue item without an article: want error")
	}
	// No queue directory at all is an error (nothing exported yet).
	bare, _ := writeFixture(t, "", "")
	if _, err := SeedQueue("refresh", []string{bare}); err == nil {
		t.Error("missing quests/queue: want error")
	}
	if _, err := SeedQueue("refresh", []string{root, root}); err == nil {
		t.Error("two args: want usage error")
	}
}
