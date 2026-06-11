//ff:func feature=quest type=parser control=sequence topic=queue
//ff:what Seed가 kind=refresh 큐 파일만 Item으로 시드하고(Key=큐 키, payload 고정) 큐 디렉토리 부재는 에러인지 검증
package refresh

import (
	"testing"

	"github.com/park-jun-woo/abloq/pkg/quests/common"
)

func TestSeed(t *testing.T) {
	root := writeInstance(t)
	items, err := Definition{}.Seed([]string{root})
	if err != nil {
		t.Fatalf("Seed: %v", err)
	}
	if len(items) != 1 || items[0].Key != "en/posts/a" {
		t.Fatalf("items = %+v, want one en/posts/a", items)
	}
	var p common.QueuePayload
	if err := items[0].DecodePayload(&p); err != nil {
		t.Fatal(err)
	}
	if p.Article != "content/en/posts/a.md" || p.Queue["freshness_days"] != "90" {
		t.Errorf("payload = %+v", p)
	}
	if len(p.Keys) != 1 || p.Keys[0] != "en/posts/a" {
		t.Errorf("keys = %v", p.Keys)
	}
	if _, err := (Definition{}).Seed([]string{t.TempDir()}); err == nil {
		t.Error("instance without blog.yaml: want error")
	}
}
